package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
)

// Estruturas para deserializar o JSON
type Report struct {
	Feature  string    `json:"name"`
	Elements []Element `json:"elements"`
}

type Element struct {
	Name  string `json:"name"`
	Steps []Step `json:"steps"`
}

type Step struct {
	Name   string `json:"name"`
	Result Result `json:"result"`
}

type Result struct {
	Status string `json:"status"`
}

func main() {
	// Leia o arquivo JSON
	byteValue, err := os.ReadFile("cucumber_report.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %s", err)
	}

	var reports []Report
	json.Unmarshal(byteValue, &reports)

	// Template HTML
	const htmlTemplate = `
    <!DOCTYPE html>
    <html>
    <head>
        <title>BDD Test Report</title>
        <style>
            body { font-family: Arial, sans-serif; }
            .passed { color: green; }
            .failed { color: red; }
        </style>
    </head>
    <body>
        <h1>BDD Test Report</h1>
        {{range .}}
            <h2>{{.Feature}}</h2>
            {{range .Elements}}
                <h3>{{.Name}}</h3>
                <ul>
                    {{range .Steps}}
                        <li class="{{.Result.Status}}">{{.Name}} - {{.Result.Status}}</li>
                    {{end}}
                </ul>
            {{end}}
        {{end}}
    </body>
    </html>
    `

	t := template.New("report")
	t, err = t.Parse(htmlTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %s", err)
	}

	// Crie o arquivo HTML
	htmlFile, err := os.Create("report.html")
	if err != nil {
		log.Fatalf("Failed to create HTML file: %s", err)
	}
	defer htmlFile.Close()

	// Execute o template com os dados do JSON
	err = t.Execute(htmlFile, reports)
	if err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}

	log.Println("Report generated: report.html")
}
