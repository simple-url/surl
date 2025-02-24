package utils_test

import (
	"testing"

	"github.com/simple-url/surl/utils"
)

func TestCapture(t *testing.T) {
	cc := utils.NewCommandCapture()
	cc.Capture([]string{"surl", "run", "hello", "-v", "-p", "./surl.json"})
	if cc.Command != "run" {
		t.Fatal("wrong command")
	}

	if cc.Arg != "hello" {
		t.Fatal("wrong arg")
	}

	if len(cc.Flags) != 2 {
		t.Fatal("wrong flag")
	}
	if cc.Flags[0] != "-v" {
		t.Fatal("flag -v not found")
	}
	if cc.Flags[1] != "-p" {
		t.Fatal("flag -p not found")
	}

	value, isFound := cc.FlagsKeyVal["-p"]
	if !isFound {
		t.Fatal("key val -p not found")
	}
	if *value != "./surl.json" {
		t.Fatal("key -p doesn't have correct value")
	}
}

func TestHasFlag(t *testing.T) {
	cc := utils.NewCommandCapture()
	cc.Capture([]string{"surl", "run", "hello", "-v", "-p", "./surl.json"})
	if !cc.HasFlag("-v") {
		t.Fatal("flag -v should exists")
	}
	if !cc.HasFlag("-p") {
		t.Fatal("flag -p should exists")
	}
	if cc.HasFlag("-z") {
		t.Fatal("flag -z should not exists")
	}
}

func TestGetVal(t *testing.T) {
	cc := utils.NewCommandCapture()
	cc.Capture([]string{"surl", "run", "hello", "-v", "-p", "./surl.json"})
	val, is_found := cc.GetVal("-p")
	if *val != "./surl.json" || !is_found {
		t.Fatal("flag -p should exists and has value './surl.json'")
	}
	val, is_found = cc.GetVal("-v")
	if val != nil || !is_found {
		t.Fatal("flag -v should exists and has value nil")
	}
	val, is_found = cc.GetVal("-z")
	if val != nil || is_found {
		t.Fatal("flag -z should not exists and has value nil")
	}
}
