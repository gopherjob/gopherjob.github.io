package templates

import (
	"embed"
	"html/template"
	"log"
)

//go:embed assets/*
var templatesFS embed.FS

var (
	JobListTemplate        = loadTemplate("job list.html")
	JobDescriptionTemplate = loadTemplate("job description.html")
)

func loadTemplate(name string) *template.Template {
	tmpl, err := template.New(name).Funcs(template.FuncMap{
		"name":  func() string { return name },
		"minus": func(a, b int) int { return a - b },
		"plus":  func(a, b int) int { return a + b },
		"seq": func(start, end int) []int {
			r := make([]int, end-start+1)
			for i := range r {
				r[i] = start + i
			}
			return r
		},
	}).ParseFS(templatesFS, "assets/"+name)
	if err != nil {
		log.Fatalf("Error parsing template file %s: %v", name, err)
	}
	return tmpl
}
