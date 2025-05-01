VERBOSE=-v

test:
	- go test ./internal/testdata/ $(VERBOSE)
	- go test ./internal/jobs/controller/ $(VERBOSE)

run: VERBOSE=
run: test
	- go run cmd/main.go

.PHONY: run test