// Package classification 8th Light Coding Challenge
//
// the purpose of this code is to answer 3 questions from 8th Light. These solutions will
// be used to gauge my coding ability and style. Note: this header is the header that I use
// to autogenerate swagger documentation for my microservices. While that isn't needed here,
// I figured it would be good to keep my code formats similar.
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     BasePath: /api
//     Version: 0.0.1
//
//     Consumes:
//     - application/vnd.api+json
//     - application/json
//
//     Produces:
//     - application/vnd.api+json
//     - application/json
//
//     security:
//     - wbrush_apikey:
//	   - wbrush_oauth2:
//	     read_scope
//
// 	   securityDefinitions:
//       wbrush_apikey:
//         type: apiKey
//         name: KEY
//         in: header
//       wbrush_oauth2:
//         type: oauth2
//         description: example
//         flow: accessCode
//         authorizationUrl: 'https://localhost/oauth2/auth'
//         tokenUrl: 'https://localhost/oauth2/token'
//         scopes:
//           read_scope: description here
//           write_scope: description here
//
// swagger:meta
package main

//go:generate swagger generate spec -m -o ./docs/swagger.json

/*
for more information about generating swagger.json from comments, see:
	https://www.ribice.ba/swagger-golang/
*/

import (
	"github.com/sirupsen/logrus"
	"github.com/wbrush/8th_light_test/question1"
	"github.com/wbrush/8th_light_test/question2"
)

var (
	commit  string
	builtAt string
)

func main() {
	logrus.Info("------------------------------")
	logrus.Info("Starting 8th_light_test")
	logrus.Info("Version:", commit, "; Build Date:", builtAt)
	logrus.Info("------------------------------")

	//  do work here
	DataQ1 := "..2965a"
	ResultQ1, ErrQ1 := question1.RemoveNonNumbers(DataQ1)
	logrus.WithFields(logrus.Fields{
		"ResultQ1": ResultQ1,
		"ErrQ1":    ErrQ1,
	}).Debug("Question1 Answers")

	DataQ2 := []int{1, 2, 3, 2, 4}
	ResultQ2, ErrQ2 := question2.RemoveDuplicates(DataQ2)
	logrus.WithFields(logrus.Fields{
		"ResultQ2": ResultQ2,
		"ErrQ2":    ErrQ2,
	}).Debug("Question2 Answers")

}
