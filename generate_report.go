package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Estruturas para deserializar o JSON do Godog
type Step struct {
	Keyword string `json:"keyword"`
	Name    string `json:"name"`
	Result  struct {
		Status string `json:"status"`
	} `json:"result"`
}

type Element struct {
	Name     string    `json:"name"`
	Steps    []Step    `json:"steps"`
	Keyword  string    `json:"keyword"`
	Examples []Example `json:"examples"`
}

type Example struct {
	Name        string        `json:"name"`
	TableHeader []TableHeader `json:"tableHeader"`
	TableBody   []TableBody   `json:"tableBody"`
}

type TableHeader struct {
	Cells []string `json:"cells"`
}

type TableBody struct {
	Cells []string `json:"cells"`
}

type Feature struct {
	Name     string    `json:"name"`
	Elements []Element `json:"elements"`
}

type Report struct {
	Features []Feature `json:"features"`
}

func main() {
	// Lê o arquivo JSON gerado pelo Godog
	jsonFile, err := os.Open("cucumber_report.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var report Report
	err = json.Unmarshal(byteValue, &report)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	// Template HTML baseado no fornecido
	const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Living Documentation</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .header {
            background-color: #333;
            color: #fff;
            padding: 10px 0;
            text-align: center;
            margin-bottom: 20px;
        }
        .container {
            display: flex;
        }
        .sidebar {
            width: 20%;
            background-color: #2f2f2f;
            color: #fff;
            padding: 20px;
            box-shadow: 2px 0 5px rgba(0,0,0,0.1);
            height: 100vh;
            overflow-y: auto;
        }
        .content {
            width: 75%;
            padding: 20px;
            overflow-y: auto;
            background-color: #fff;
            box-shadow: 2px 0 5px rgba(0,0,0,0.1);
        }
        .feature-list {
            list-style-type: none;
            padding: 0;
            margin: 0;
        }
        .feature-list li {
            padding: 10px;
            cursor: pointer;
            border-bottom: 1px solid #444;
        }
        .feature-list li:hover {
            background-color: #444;
        }
        .scenario {
            margin-bottom: 20px;
            background-color: #f9f9f9;
            padding: 10px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
        }
        .step {
            margin-left: 20px;
        }
        .passed {
            color: green;
        }
        .failed {
            color: red;
        }
        table {
            border-collapse: collapse;
            width: 100%;
            margin-bottom: 20px;
        }
        th, td {
            border: 1px solid #ddd;
            padding: 8px;
        }
        th {
            background-color: #f2f2f2;
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>Living Documentation</h1>
        <p>Generated on {{.GeneratedAt}}</p>
    </div>
    <div class="container">
        <div class="sidebar">
            <h2>Features</h2>
            <ul class="feature-list">
                {{range .Features}}
                    <li onclick="showFeature('{{.Name}}')">{{.Name}}</li>
                {{end}}
            </ul>
        </div>
        <div class="content">
            <div id="feature-details">
                <h2>Select a feature to see details</h2>
            </div>
        </div>
    </div>
    <script>
        const features = {{.FeaturesJSON}};

        function showFeature(featureName) {
            const feature = features.find(f => f.name === featureName);
            const featureDetailsDiv = document.getElementById('feature-details');
            let html = '<h2>Feature: ' + feature.name + '</h2>';
            feature.elements.forEach(function(element) {
                html += '<div class="scenario">' +
                    '<h3>' + element.keyword + ': ' + element.name + '</h3>' +
                    '<table>' +
                        '<thead>' +
                            '<tr>' +
                                '<th>Keyword</th>' +
                                '<th>Step</th>' +
                                '<th>Status</th>' +
                            '</tr>' +
                        '</thead>' +
                        '<tbody>';
                element.steps.forEach(function(step) {
                    html += '<tr class="' + step.result.status + '">' +
                        '<td>' + step.keyword + '</td>' +
                        '<td>' + step.name + '</td>' +
                        '<td>' + step.result.status + '</td>' +
                    '</tr>';
                });
                html += '</tbody>' +
                    '</table>' +
                '</div>';
                // Adicionar exemplos se houver
                if (element.examples && element.examples.length > 0) {
                    html += '<div class="examples">' +
                        '<h4>Examples:</h4>' +
                        '<table>' +
                            '<thead>' +
                                '<tr>';
                    element.examples[0].tableHeader[0].cells.forEach(function(cell) {
                        html += '<th>' + cell + '</th>';
                    });
                    html += '</tr>' +
                            '</thead>' +
                            '<tbody>';
                    element.examples[0].tableBody.forEach(function(row) {
                        html += '<tr>';
                        row.cells.forEach(function(cell) {
                            html += '<td>' + cell + '</td>';
                        });
                        html += '</tr>';
                    });
                    html += '</tbody>' +
                        '</table>' +
                    '</div>';
                }
            });
            featureDetailsDiv.innerHTML = html;
        }
    </script>
</body>
</html>
`

	featuresJSON, err := json.Marshal(report.Features)
	if err != nil {
		log.Fatalf("Failed to marshal features: %s", err)
	}

	data := struct {
		GeneratedAt  string
		Features     []Feature
		FeaturesJSON template.JS
	}{
		GeneratedAt:  time.Now().Format("January 2, 2006 at 3:04pm"),
		Features:     report.Features,
		FeaturesJSON: template.JS(featuresJSON),
	}

	tmpl, err := template.New("report").Parse(tpl)
	if err != nil {
		log.Fatalf("Failed to parse template: %s", err)
	}

	reportFile, err := os.Create("final_report.html")
	if err != nil {
		log.Fatalf("Failed to create HTML file: %s", err)
	}
	defer reportFile.Close()

	err = tmpl.Execute(reportFile, data)
	if err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}

	fmt.Println("HTML report generated successfully.")
}
