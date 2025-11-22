

run-job:
	bash scripts/job.sh

build:
	go build -o ./bin/out ./cmd/.

run: build
	./bin/out