package bdd

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

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
	ctx.Step(`^calculate POST request is send to "([^"]*)" with (\d+) in the request body$`, t.calculatePOSTRequestIsSendToWithInTheRequestBody)
	ctx.Step(`^calculate POST request is send to "([^"]*)" with empty request body$`, t.calculatePOSTRequestIsSendToWithEmptyRequestBody)
	ctx.Step(`^the response code is (\d+)$`, t.theResponseCodeIs)
	ctx.Step(`^the response body contains (\d+),(\d+),(\d+),(\d+),(\d+)$`, t.theResponseBodyContains)
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

func (t *TestPackSizeCalculate) calculatePOSTRequestIsSendToWithInTheRequestBody(path string, itemsQty int) error {
	var boxItemsRequest []byte
	if len(t.fContext.packSizes) > 0 {
		boxItemsRequest = []byte(`{"quantity": ` + strconv.Itoa(itemsQty) + `, "packSizes": [` + fmt.Sprintf("%v", t.fContext.packSizes) + `]}`)
	} else {
		boxItemsRequest = []byte(`{"quantity": ` + strconv.Itoa(itemsQty) + `}`)

	}
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

func (t *TestPackSizeCalculate) theResponseBodyContains(arg1, arg2, arg3, arg4, arg5 int) error {
	expected := map[int]int{
		250:  arg1,
		500:  arg2,
		1000: arg3,
		2000: arg4,
		5000: arg5,
	}
	if err := checkBoxItemResponse(t.fContext.boxItemsResponse, expected); err != nil {
		t.testingT.Fatalf("Error checking response: %v", err)
	}

	return nil

}

func (t *TestPackSizeCalculate) customPackSizesAreSetTo(packSizes string) error {
	ps := strings.Split(packSizes, ",")
	for i, s := range ps {
		t.fContext.packSizes[i], _ = strconv.Atoi(s)
	}

	return nil
}

func (t *TestPackSizeCalculate) theResponseBodyHas(arg1 string) error {
	return godog.ErrPending
}

func checkBoxItemResponse(response boxItemsResponse, expected map[int]int) error {
	for k, v := range expected {
		if response.BoxItems[k] != v {
			return fmt.Errorf("expected %d pack to be %d, but got %d", k, v, response.BoxItems[k])
		}
	}
	return nil
}
