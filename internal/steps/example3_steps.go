package steps

import (
	"github.com/betoth/automation-godog/internal/support"
	"github.com/cucumber/godog"
)

func yetAnotherPrecondition() error {
	return nil
}

func yetAnotherActionIsPerformed() error {

	return nil
}

func yetAnotherOutcomeIsExpected() error {

	return nil
}

func RegisterExample3Steps(ctx *godog.ScenarioContext) {

	ctx.Step(`^yet another precondition$`, yetAnotherPrecondition)
	ctx.Step(`^yet another action is performed$`, yetAnotherActionIsPerformed)
	ctx.Step(`^yet another outcome is expected$`, yetAnotherOutcomeIsExpected)
}

func init() {

	support.RegisterSteps(RegisterExample3Steps)
}
