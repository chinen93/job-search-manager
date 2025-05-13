VERBOSE=-v

test:
	- go test ./internal/testdata/ $(VERBOSE)
	- go test ./internal/jobs/controller/ $(VERBOSE)

run_helper: VERBOSE=
run_helper: test
	- go run cmd/helper/main.go

run: VERBOSE=
run: test
	- go run cmd/server/main.go

.PHONY: run run_helper test