package steps

import (
	"fmt"
	"net/http"

	"github.com/cucumber/godog"
)

var (
	response *http.Response
	err      error
)

func iMakeAGETRequestToThePublicAPI() error {
	response, err = http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		return fmt.Errorf("failed to make GET request: %v", err)
	}
	return nil
}

func theResponseStatusCodeShouldBe(expectedStatusCode int) error {
	if response.StatusCode != expectedStatusCode {
		return fmt.Errorf("expected status code %d but got %d", expectedStatusCode, response.StatusCode)
	}
	return nil
}

func RegisterExemploAPISteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I make a GET request to the public API$`, iMakeAGETRequestToThePublicAPI)
	ctx.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
}

// Renomeie a função InitializeScenario para algo diferente
func InitializeExemploAPIScenario(ctx *godog.ScenarioContext) {
	RegisterExemploAPISteps(ctx)
}
