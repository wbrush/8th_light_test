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
	"github.com/wbrush/8th_light_test/models"
	"github.com/wbrush/8th_light_test/question1"
	"github.com/wbrush/8th_light_test/question2"
	"github.com/wbrush/8th_light_test/question3"
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

	DataQ3_ids := []string{"7-623", "8235", "8-235"}
	DataQ3_specialties := []models.Specialty{
		{Id: 1381, Name: "front-end web development"},
		{Id: 8235, Name: "data engineering"},
		{Id: 3434, Name: "API design"},
		{Id: 7623, Name: "security"},
		{Id: 9153, Name: "UX"},
	}
	ResultQ3, ErrQ3 := question3.GetSpecialty(DataQ3_ids, DataQ3_specialties)
	logrus.WithFields(logrus.Fields{
		"ResultQ3": ResultQ3,
		"ErrQ3":    ErrQ3,
	}).Debug("Question3 Answers")

}
