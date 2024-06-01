package main

import (
	"testing"

	"github.com/cucumber/godog"

	steps "github.com/betoth/automation-godog/tests"
)

func InitializeAllScenarios(ctx *godog.ScenarioContext) {
	steps.RegisterExemploAPISteps(ctx)
	steps.RegisterExampleSteps(ctx)
	// Adicione mais chamadas para as funções de inicialização de cenário de outros arquivos de passos, se houver
}

func TestFeature(t *testing.T) {
	opts := godog.Options{
		Format:   "pretty",
		Paths:    []string{"tests/features"},
		TestingT: t,
	}

	status := godog.TestSuite{
		Name:                "bdd",
		ScenarioInitializer: InitializeAllScenarios,
		Options:             &opts,
	}.Run()

	if status != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
