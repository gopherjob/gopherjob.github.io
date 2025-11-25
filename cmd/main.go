package main

import (
	"fmt"
	htmlCreator "github.com/gopherjobs/gopherjobs.github.io/internal/html_creator"
	jobReader2 "github.com/gopherjobs/gopherjobs.github.io/internal/job_reader"
	"log"
	"log/slog"
)

func main() {
	println("running")
	logger := slog.Default()
	jr := jobReader2.NewJobReader(logger)
	jobs, err := jr.ReadAllFiles("job_results")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(jobs))
	for idx, job := range jobs {
		fmt.Printf("%d. %v, %v\n", idx+1, job.Remote, job.Relocation)
	}

	jg := htmlCreator.NewHTMLCreator(logger)
	err = jg.Generate(jobs, "static")
	if err != nil {
		log.Fatal(err)
	}

}
