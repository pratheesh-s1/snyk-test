module snyk-test

go 1.25.1

require (
	github.com/google/uuid v1.6.0
	github.com/pratheesh-s1/snyk-test/sample-module v0.0.0-00010101000000-000000000000
)

replace github.com/pratheesh-s1/snyk-test/sample-module => ./sample-module
