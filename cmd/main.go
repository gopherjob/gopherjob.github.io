package main

import (
	"fmt"
	jobReader2 "github.com/milad-rasouli/bluejob/internal/job_reader"
	"log"
	"log/slog"
)

func main() {
	println("running")
	jr := jobReader2.NewJobReader(slog.Default())
	jobs, err := jr.ReadAllFiles("job_results")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(jobs))
	for idx, job := range jobs {
		fmt.Printf("%d. %v, %v\n", idx+1, job.Remote, job.Relocation)
	}
}
