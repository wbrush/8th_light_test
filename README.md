# 8th_light_test
This service was written as a take home coding test for a position at 8th Light. 

## Test Status

### Coding Status
Completed all coding challenges. However if this code were to be used in a production environment, I would curious about the performance and memory requirements of some of the code. Therefore ideally,
 I would add performance benchmarks to the code which I didn't have time to include here.

### Question Status
Question answer is included here.

Question A) How might your design differ if the list of specialties was expected to never change? What if changes frequently? 

Answer: 
If the specialties don't change, it would be easier to create a map the first time and keep it in memory for more efficient processing. I would probably create a function to load the map and update the 
existing function(s) to check that it has been loaded (returning an error if not). In production, I might also be tempted to store the specialties in local storage (file, SQL DB, hardcoded, etc) that the
code could read to initialize the running code (probably using an init() function to preload the in-memory map). Which local storage mechanism would probably depend on the amount of data and how it is 
provided (i.e. initial data set is in Excel).

If it changes frequently, I might be tempted to have a function to load the specialties into in-memory storage. This is especially true in a production environment where this would be accessed through 
an API. There would need to be a way to perform the basic CRUD operations on the specialties list. Ideally to keep the service "stateless", the data would need to be persisted into a memory cache (such
as Redis) or DB (probably SQL such as Postgres or MySQL) which can be accessed by all instances of the running code.

## Design tradeoffs
Searching on Google gave a quick, simple answer to the first challenge. This seems fairly efficient and simple so I would probably stay with this solution. Challenges 2 & 3 were solved more manually
without an internet search. There may be more efficient ways to solve or test them. My first priority when coding is find a solution that works and then optimize it if I have time or optimization is 
needed. As noted in the question2 unit tests, the testing for the RemoveDuplicates() function could probably be simplified.

# Development Notes

## Unit Testing
To run the unit tests locally, you would need to have the go compiler installed. I can add the linux shell executable to the repo if needed. To run locally:
1. run the command "go get github.com/wbrush/8th_light_test" from the command line
2. run " go test ./... -covermode=count"
4. you should see unit test messages similar to:

?       github.com/wbrush/8th_light_test        [no test files]
?       github.com/wbrush/8th_light_test/models [no test files]
ok      github.com/wbrush/8th_light_test/question1      (cached)        coverage: 80.0% of statements
ok      github.com/wbrush/8th_light_test/question2      (cached)        coverage: 100.0% of statements
ok      github.com/wbrush/8th_light_test/question3      (cached)        coverage: 89.3% of statements

## Developer Notes
These are notes that I recommend providing to the next dev who might be working on this.

### Status
This code is complete and fully working as I understand the requirements. I may have missed some unit test cases that should be added if bugs are found.

## Required Packages
This service uses go-modules so the modules are documented in the go.mod file in the mainline. However, key modules are listed here for completeness.

### logrus
Framework for logging.

go get -u github.com/sirupsen/logrus

## Process
This would spell out the deployment steps once developed and tested. I would expect this to be a stand alone microservice. I considered serverless but I don't think it would be good to implement it using serverless, especially considering questions 2 & 3.

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
