package main

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/betoth/automation-godog/tests/steps"
	"github.com/betoth/automation-godog/tests/support"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	fmt.Println("Initializing scenario in main_test.go")
	support.InitializeScenarios(ctx)
}

func TestMain(m *testing.M) {
	fmt.Println("Running TestMain")

	// Configuração para a saída "pretty" no terminal
	prettyOpts := godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
		Paths:  []string{"features"},
	}

	// Executar testes com saída "pretty"
	prettyStatus := godog.TestSuite{
		Name:                "godogs-pretty",
		ScenarioInitializer: InitializeScenario,
		Options:             &prettyOpts,
	}.Run()

	// Configuração para gerar o arquivo JSON
	jsonFile, err := os.Create("cucumber_report.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonOpts := godog.Options{
		Output: jsonFile,
		Format: "cucumber",
		Paths:  []string{"features"},
	}

	// Executar testes para gerar o arquivo JSON
	jsonStatus := godog.TestSuite{
		Name:                "godogs-json",
		ScenarioInitializer: InitializeScenario,
		Options:             &jsonOpts,
	}.Run()

	if prettyStatus > 0 || jsonStatus > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}
