package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/authgear/authgear-server/pkg/lib/authflowclient"
	"github.com/authgear/authgear-server/pkg/util/httputil"
	"gopkg.in/yaml.v2"
)

func TestAuthflow(t *testing.T) {
	err := filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".yaml" {
			return nil
		}

		testCases, err := loadTestCasesFromPath(path)
		if err != nil {
			return err
		}

		for _, testCase := range testCases {
			runTestCase(t, testCase)
		}

		return nil
	})

	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func loadTestCasesFromPath(path string) ([]TestCase, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var testCases []TestCase
	decoder := yaml.NewDecoder(strings.NewReader(string(data)))
	for {
		var testCase TestCase
		err := decoder.Decode(&testCase)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		testCases = append(testCases, testCase)
	}

	return testCases, nil
}

func runTestCase(t *testing.T, testCase TestCase) {
	t.Logf("running test case: %s\n", testCase.Name)
	client := authflowclient.NewClient(context.Background(), "localhost:4000", httputil.HTTPHost(fmt.Sprintf("%s.portal.localhost:4000", testCase.Project)))

	var stateToken string

	for i, step := range testCase.Steps {
		var stepName = step.Name
		if stepName == "" {
			stepName = fmt.Sprintf("step %d", i+1)
		}

		var flowResponse *authflowclient.FlowResponse
		var flowErr error

		switch step.Action {
		case StepActionCreate:
			var flowReference authflowclient.FlowReference
			err := json.Unmarshal([]byte(step.Input), &flowReference)
			if err != nil {
				t.Errorf("failed to parse input in '%s': %v\n", stepName, err)
				return
			}

			flowResponse, flowErr = client.Create(flowReference, "")

		case StepActionInput:
			fallthrough

		default:
			var input map[string]interface{}
			err := json.Unmarshal([]byte(step.Input), &input)
			if err != nil {
				t.Errorf("failed to parse JSON input in '%s': %v\n", stepName, err)
				return
			}

			flowResponse, flowErr = client.Input(nil, nil, stateToken, input)
		}

		if flowResponse != nil {
			stateToken = flowResponse.StateToken
		}

		for _, assertion := range step.Assert {
			value, ok := TranslateAssertValue(flowResponse, flowErr, assertion.Field)
			if !ok {
				t.Errorf("field '%s' not found in '%s'\n%v\n%v", assertion.Field, stepName, flowResponse, flowErr)
				return
			}

			assertErr := PerformAssertion(assertion, value)
			if assertErr != nil {
				t.Errorf("assertion failed in '%s': %v\n%v\n%v", stepName, assertErr, flowResponse, flowErr)
				return
			}
		}
	}
}
