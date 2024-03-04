package bdd

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func PrepareTestSuites(t *testing.T, fContext *FeatureContext) []godog.TestSuite {
	options := func(featuresPath ...string) *godog.Options {
		return &godog.Options{
			Format:        "pretty",
			Paths:         featuresPath,
			Output:        colors.Colored(os.Stdout),
			StopOnFailure: true,
			Strict:        true,
			TestingT:      t,
		}
	}

	return []godog.TestSuite{
		{
			Name: "Calculate pack items",
			ScenarioInitializer: func(ctx *godog.ScenarioContext) {
				Initializer := newPackCalculatorFeatureEnvironment(t, true, fContext)
				Initializer.InitializeTest(ctx)
			},
			Options: options(
				"cmd/bdd/features/pack_calculator.feature",
			),
		},
	}
}
