package support

import (
	"github.com/cucumber/godog"
)

var steps []func(*godog.ScenarioContext)

func RegisterSteps(stepFunc func(*godog.ScenarioContext)) {
	steps = append(steps, stepFunc)
}

func InitializeScenarios(sc *godog.ScenarioContext) {

	for _, step := range steps {

		step(sc)
	}
}
