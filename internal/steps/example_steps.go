package steps

import (
	"github.com/betoth/automation-godog/internal/support"

	"github.com/cucumber/godog"
)

func aPrecondition() error {

	return nil
}

func anActionIsPerformed() error {

	return nil
}

func anOutcomeIsExpected() error {

	return nil
}

func RegisterExampleSteps(ctx *godog.ScenarioContext) {

	ctx.Step(`^a cachorro$`, aPrecondition)
	ctx.Step(`^an action is performed$`, anActionIsPerformed)
	ctx.Step(`^an outcome is expected$`, anOutcomeIsExpected)
}

func init() {

	support.RegisterSteps(RegisterExampleSteps)
}
