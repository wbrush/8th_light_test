# 8th_light_test
This service was written as a take home coding test for a position at 8th Light. 

## Test Status

## Test Task

## Design tradeoffs


# Development Notes

## Unit Testing

## Developer Notes
These are notes that I recommend providing to the next dev who might be working on this.

### Status

## Required Packages
This service uses go-modules so the modules are documented in the go.mod file in the mainline. However, key modules are listed here for completeness.

### logrus
Framework for logging.

go get -u github.com/sirupsen/logrus

## Process
This would spell out the deployment steps once developed and tested.

## Running locally
To run this locally, you would need to have the go compiler installed. I can add the windows executable to the repo if needed. To run locally:
1. run the command "go get github.com/wbrush/8th_light_test" from the command line
2. run " go build -ldflags "-X main.commit=`git rev-parse --short HEAD` -X main.builtAt=`date +%FT%T%z`" "
3. if on unix, run "./healthjoy_test" 
4. you should see startup messages similar to:

time="2021-07-11 19:32:30.812310" level=info msg=------------------------------
time="2021-07-11 19:32:30.812310" level=info msg="Starting 8th_light_test"
time="2021-07-11 19:32:30.812310" level=info msg="Version:; Build Date:"
time="2021-07-11 19:32:30.812310" level=info msg=------------------------------

