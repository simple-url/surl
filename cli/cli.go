package cli

import (
	"regexp"
	"sort"
)

type Route struct {
	Prefix  []string
	Execute func(*Cli)
}

type RouteFlag struct {
	Command []string
	Flag    []string
	Execute func(*Cli)
}

type Cli struct {
	Exe            string
	Command        string
	Commands       []string
	Flags          map[string]*string
	routes         []Route
	is_routes_sort bool
	route_default  func(*Cli)
	isCapture      bool
}

func IsFlag(text string) bool {
	isFlagShort, _ := regexp.MatchString("-.*", text)
	isFlagLong, _ := regexp.MatchString("--.*", text)
	return isFlagShort || isFlagLong
}

func NewCli() *Cli {
	return &Cli{
		Command:        "",
		Commands:       []string{},
		Flags:          map[string]*string{},
		route_default:  func(c *Cli) {},
		is_routes_sort: false,
		isCapture:      false,
	}
}

func (c *Cli) ResetCapture() {
	c.Command = ""
	c.Commands = []string{}
	c.Flags = map[string]*string{}
	c.isCapture = false
}

func (c *Cli) Capture(args []string) {
	if c.isCapture {
		c.ResetCapture()
	}
	if len(args) < 1 {
		return
	}
	c.isCapture = true
	// set exe
	c.Exe = args[0]
	args = args[1:]
	num_args := len(args)
	for i := 0; i < num_args; i++ {
		// IsFlag or Command
		if IsFlag(args[i]) {
			key := args[i]
			var value *string = nil
			// IsFlag has value
			if i+1 < num_args {
				if !IsFlag(args[i+1]) {
					value = &args[i+1]
					i += 1
				}
			}
			c.Flags[key] = value
		} else {
			if c.Command == "" {
				c.Command = args[i]
			}
			c.Commands = append(c.Commands, args[i])
		}
	}
}

func (c *Cli) GetFlagVal(flags []string) (*string, bool) {
	var val *string = nil
	is_found := false
	for _, item := range flags {
		val, is_found = c.Flags[item]
		if is_found {
			break
		}
	}
	return val, is_found
}

func (c *Cli) Route(prefix []string, exe func(*Cli)) {
	c.routes = append(c.routes, Route{
		Prefix:  prefix,
		Execute: exe,
	})
}

func (c *Cli) RouteDefault(exe func(*Cli)) {
	c.route_default = exe
}

func IsSliceIn(a []string, b []string) bool {
	if len(a) > len(b) {
		return false
	}
	for idx, item := range a {
		if item != b[idx] {
			return false
		}
	}
	return true
}

func (c *Cli) SortRoute() {
	if c.is_routes_sort {
		return
	}
	sort.Slice(c.routes, func(i, j int) bool {
		return len(c.routes[i].Prefix) > len(c.routes[j].Prefix)
	})
	c.is_routes_sort = true
}

func (c *Cli) Run() {
	// check route
	c.SortRoute()
	for _, item := range c.routes {
		if IsSliceIn(item.Prefix, c.Commands) {
			item.Execute(c)
			return
		}
	}
	// if route not found
	c.route_default(c)
}
