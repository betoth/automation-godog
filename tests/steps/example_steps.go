package steps

import (
	"fmt"

	"github.com/betoth/automation-godog/tests/support"

	"github.com/cucumber/godog"
)

func aPrecondition() error {
	fmt.Println("Executing: aPrecondition in example_steps.go")
	return nil
}

func anActionIsPerformed() error {
	fmt.Println("Executing: anActionIsPerformed in example_steps.go")
	return nil
}

func anOutcomeIsExpected() error {
	fmt.Println("Executing: anOutcomeIsExpected in example_steps.go")
	return nil
}

func RegisterExampleSteps(ctx *godog.ScenarioContext) {
	fmt.Println("Registering steps in example_steps.go")
	ctx.Step(`^a precondition$`, aPrecondition)
	ctx.Step(`^an action is performed$`, anActionIsPerformed)
	ctx.Step(`^an outcome is expected$`, anOutcomeIsExpected)
}

func init() {
	fmt.Println("Calling RegisterExampleSteps in example_steps.go")
	support.RegisterSteps(RegisterExampleSteps)
}
