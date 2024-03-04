package main

import (
	"log"
	"pack-items/cmd/bdd"
	"testing"
)

func RunService() {
	if err := bdd.Setup(); err != nil {
		log.Fatalf("Error setting up docker-compose: %v", err)
	}
}

func Test_GherkinFeatureScenarios(t *testing.T) {
	RunService()

	featureContext := bdd.NewFeatureContext("http://localhost:8080")

	for _, s := range bdd.PrepareTestSuites(t, featureContext) {
		if s.Run() != 0 {
			t.FailNow()
		}
	}
}
