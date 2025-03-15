package cli_test

import (
	"testing"

	"github.com/simple-url/surl/cli"
)

func TestCapture(t *testing.T) {
	cc := cli.NewCli()
	cc.Capture([]string{"surl", "run", "hello", "-v", "-p", "./surl.json"})
	// Test Exe
	if cc.Exe != "surl" {
		t.Fatal("wrong exe name")
	}

	// Test Command
	if cc.Command != "run" {
		t.Fatal("wrong command")
	}

	// Test Commands
	if len(cc.Commands) != 2 {
		t.Fatal("num commands should be 2")
	}
	commands := []string{"run", "hello"}
	for idx, item := range cc.Commands {
		if commands[idx] != item {
			t.Fatal("invalid commands")
		}
	}

	// Test Flags
	if len(cc.Flags) != 2 {
		t.Fatal("wrong number of flag")
	}
	v_val, v_is_found := cc.GetFlagVal([]string{"-v"})
	if !v_is_found || v_val != nil {
		t.Fatal("flag and value -v should found")
	}

	p_val, p_is_found := cc.GetFlagVal([]string{"-p"})
	if !p_is_found || *p_val != "./surl.json" {
		t.Fatal("flag and value -p should found and return ./surl.json")
	}
}

type Called struct {
	Run      bool
	RunHelp  bool
	List     bool
	ListHelp bool
	Default  bool
}

func (c *Called) Reset() {
	c.Run = false
	c.RunHelp = false
	c.List = false
	c.ListHelp = false
	c.Default = false
}

func TestRouteRun(t *testing.T) {
	// Given
	called := Called{
		Run:      false,
		RunHelp:  false,
		List:     false,
		ListHelp: false,
		Default:  false,
	}
	cc := cli.NewCli()
	cc.Route([]string{"list"}, func(c *cli.Cli) {
		called.List = true
	})
	cc.Route([]string{"list", "help"}, func(c *cli.Cli) {
		called.ListHelp = true
	})
	cc.Route([]string{"run", "help"}, func(c *cli.Cli) {
		called.RunHelp = true
	})
	cc.Route([]string{"run"}, func(c *cli.Cli) {
		if c.Commands[1] != "request" {
			t.Fatal("wrong command")
		}
		called.Run = true
	})
	cc.RouteDefault(func(c *cli.Cli) {
		called.Default = true
	})

	// When Expect
	cc.Capture([]string{"surl", "list"})
	cc.Run()
	if !called.List {
		t.Fatal("route list is not executed")
	}
	called.Reset()

	cc.Capture([]string{"surl", "list", "help"})
	cc.Run()
	if !called.ListHelp {
		t.Fatal("route list help is not executed")
	}
	called.Reset()

	cc.Capture([]string{"surl", "run", "help"})
	cc.Run()
	if !called.RunHelp {
		t.Fatal("route run help is not executed")
	}
	called.Reset()

	cc.Capture([]string{"surl", "run", "request"})
	cc.Run()
	if !called.Run {
		t.Fatal("route run is not executed")
	}
	called.Reset()

	cc.Capture([]string{"surl"})
	cc.Run()
	if !called.Default {
		t.Fatal("route default is not executed")
	}
}
