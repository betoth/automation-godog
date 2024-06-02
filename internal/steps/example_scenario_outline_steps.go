package steps

import (
	"fmt"

	"github.com/betoth/automation-godog/internal/support"

	"github.com/cucumber/godog"
)

func aPreconditionWithInput(input int) error {
	return nil
}

func anActionIsPerformedWithInput(input int) error {
	return nil
}

func theOutcomeShouldBe(output string) error {
	if output == "success" || output == "failure" {
		return nil
	}
	return fmt.Errorf("unexpected outcome: %s", output)
}

func RegisterExampleScenarioOutlineSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^a precondition with (\d+)$`, aPreconditionWithInput)
	ctx.Step(`^an action is performed with (\d+)$`, anActionIsPerformedWithInput)
	ctx.Step(`^the outcome should be (.+)$`, theOutcomeShouldBe)
}

func init() {
	support.RegisterSteps(RegisterExampleScenarioOutlineSteps)
}
