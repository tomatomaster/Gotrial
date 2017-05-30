package main

import (
	"os"
	"os/exec"
	"testing"
)

var not = map[string][]string{
	"a": {"d", "e"},
	"b": {"f", "g"},
	"c": {"h"},
}

var cir = map[string][]string{
	"a": {"b", "e"},
	"b": {"a", "g"},
	"c": {"h"},
}

func TestCirculationCheck(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		checkcirculation(cir)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCirculationCheck")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
