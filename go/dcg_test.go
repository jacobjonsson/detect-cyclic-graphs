package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	Nodes  []Node `json:"nodes"`
	Edges  []Edge `json:"edges"`
	Cyclic bool   `json:"cyclic"`
}

func getTestData() []TestCase {
	testData, err := os.Open("../test-data.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var testCases []TestCase
	err = json.NewDecoder(testData).Decode(&testCases)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return testCases
}

func TestIsCyclic(t *testing.T) {
	testCases := getTestData()

	for _, testCase := range testCases {
		isCyclic := IsCyclic(testCase.Nodes, testCase.Edges)
		if isCyclic != testCase.Cyclic {
			t.Errorf("Expected %t, got %t", testCase.Cyclic, isCyclic)
		}
	}
}
