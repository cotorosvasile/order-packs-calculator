package bdd

import (
	"context"
	"strconv"
	"testing"

	"encoding/json"

	"github.com/cucumber/godog"
)

type TestPackSizeCalculate struct {
	fContext *FeatureContext
	testingT *testing.T
}

func newPackCalculatorFeatureEnvironment(t *testing.T, useClient bool, featureContext *FeatureContext) *TestPackSizeCalculate {
	return &TestPackSizeCalculate{
		fContext: featureContext,
		testingT: t,
	}
}

func (t *TestPackSizeCalculate) InitializeTest(ctx *godog.ScenarioContext) {
	ctx.Before(func(_ context.Context, sc *godog.Scenario) (context.Context, error) {
		t.fContext.reset()
		return nil, nil
	})
	ctx.Step(`^system API is up and running$`, t.systemAPIIsUpAndRunning)
	ctx.Step(`^calculate POST request is send to "([^"]*)" with empty request body$`, t.calculatePOSTRequestIsSendToWithEmptyRequestBody)
	ctx.Step(`^the response code is (\d+)$`, t.theResponseCodeIs)
	ctx.Step(`^calculate POST request is send to "([^"]*)" with "([^"]*)" and (\d+) in the request body$`, t.calculatePOSTRequestIsSendToWithAndInTheRequestBody)
	ctx.Step(`^the response body contains '(.*)'$`, t.theResponseBodyContains)
}

func (t *TestPackSizeCalculate) systemAPIIsUpAndRunning() error {
	err := t.fContext.sendRequest("GET", "/health", nil)
	if err != nil {
		t.testingT.Fatalf("Error sending request: %v", err)
	}
	if t.fContext.statusCode != 200 {
		t.testingT.Fatalf("Expected status code 200, but got %d", t.fContext.statusCode)
	}
	return nil
}

func (t *TestPackSizeCalculate) calculatePOSTRequestIsSendToWithEmptyRequestBody(path string) error {
	if err := t.fContext.sendRequest("POST", path, nil); err != nil {
		t.testingT.Fatalf("Error sending request: %v", err)
	}
	return nil
}
func (t *TestPackSizeCalculate) calculatePOSTRequestIsSendToWithAndInTheRequestBody(path, packSizes string, itemsQty int) error {
	boxItemsRequest := []byte(`{"quantity": ` + strconv.Itoa(itemsQty) + `, "packSizes": ` + packSizes + `}`)
	if err := t.fContext.sendRequest("POST", path, boxItemsRequest); err != nil {
		t.testingT.Fatalf("Error sending request: %v", err)
	}
	return nil
}

func (t *TestPackSizeCalculate) theResponseCodeIs(responseCode int) error {
	if t.fContext.statusCode != responseCode {
		t.testingT.Fatalf("Expected status code %d, but got %d", responseCode, t.fContext.statusCode)
	}

	return nil
}

func (t *TestPackSizeCalculate) theResponseBodyContains(packConfig string) error {

	var expectedBoxItemsResponse boxItemsResponse

	err := json.Unmarshal([]byte(packConfig), &expectedBoxItemsResponse)
	if err != nil {
		t.testingT.Fatalf("Error unmarshalling response: %v", err)
	}

	expectedJSON, err := json.Marshal(expectedBoxItemsResponse)
	if err != nil {
		t.testingT.Fatalf("Error marshalling expected response: %v", err)
	}

	actualJSON, err := json.Marshal(t.fContext.boxItemsResponse)
	if err != nil {
		t.testingT.Fatalf("Error marshalling actual response: %v", err)
	}

	if string(actualJSON) != string(expectedJSON) {
		t.testingT.Fatalf("Expected response %s, but got %s", string(expectedJSON), string(actualJSON))
	}

	return nil

}
