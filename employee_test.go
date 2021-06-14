package main

import (
	"os/exec"
	"syscall"
	"testing"
)

func TestDataStore(t *testing.T) {
	s, err := launchApp()
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	defer s.Shutdown()

	e := testJSON(t)
	testGetEmployee(t, e)
	testSumSalary(t, e)
}

func testSumSalary(t *testing.T, e *employees) {
	expectedResult := 850000.00
	result := e.sumSalary()

	if result != expectedResult {
		t.Fatalf("expected 850000.00, but got %f", result)
	}
}

func testGetEmployee(t *testing.T, e *employees) {
	request := "Jeff"
	emptyRequest, err := e.getEmployee(request)
	if emptyRequest == nil {
		t.Fatalf("expected Employee Jeff, got nil")
	}
	if err == nil {
		t.Logf("getEmployee PASSED with %s", err)
	}
}

func testJSON(t *testing.T) *employees {
	emp, err := ReadFile("./employees.json")

	if err != nil {
		t.Fatalf("error reading JSON file: %s", err)
	}

	return emp
}

type testApp struct {
	cmd *exec.Cmd
}

func (s *testApp) Shutdown() {
	s.cmd.Process.Signal(syscall.SIGINT)
	s.cmd.Wait()
}

func launchApp() (*testApp, error) {
	cmd := exec.Command("./testapp")
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return &testApp{cmd: cmd}, nil
}
