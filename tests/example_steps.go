package steps

import (
	"github.com/cucumber/godog"
)

func iHaveAConfiguredEnvironment() error {
	return nil
}

func theTestShouldPass() error {
	return nil
}

func RegisterExampleSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have a configured environment$`, iHaveAConfiguredEnvironment)
	ctx.Step(`^the test should pass$`, theTestShouldPass)
}

// Renomeie a função InitializeScenario para algo diferente
func InitializeExampleScenario(ctx *godog.ScenarioContext) {
	RegisterExampleSteps(ctx)
}
