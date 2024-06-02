package steps

import (
	"github.com/betoth/automation-godog/internal/support"

	"github.com/cucumber/godog"
)

func anotherPrecondition() error {
	return nil
}

func anotherActionIsPerformed() error {

	return nil
}

func anotherOutcomeIsExpected() error {

	return nil
}

func RegisterExample2Steps(ctx *godog.ScenarioContext) {

	ctx.Step(`^another precondition$`, anotherPrecondition)
	ctx.Step(`^another action is performed$`, anotherActionIsPerformed)
	ctx.Step(`^another outcome is expected$`, anotherOutcomeIsExpected)
}

func init() {

	support.RegisterSteps(RegisterExample2Steps)
}
