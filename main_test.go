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
	support.InitializeScenarios(ctx)
}

func TestMain(m *testing.M) {
	fmt.Println("Running TestMain")
	opts := godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
		Paths:  []string{"features"},
	}

	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	os.Exit(status)
}
