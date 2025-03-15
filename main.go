package main

import (
	"fmt"
	"os"

	"github.com/simple-url/surl/cli"
	"github.com/simple-url/surl/command"
	"github.com/simple-url/surl/utils"
)

func RunCommand(surl_command *command.Surl, resp utils.IResponse, args []string) {
	json_path := "./surl.json"
	app := cli.NewCli()
	app.Capture(args)
	app.Route([]string{"list"}, func(c *cli.Cli) {
		_, isHelp := c.GetFlagVal([]string{"--help", "-h"})
		if isHelp {
			surl_command.ListHelp()
			return
		}
		override_path, is_override := c.GetFlagVal([]string{"--path", "-p"})
		if is_override {
			json_path = *override_path
		}
		err := surl_command.ReadJson(json_path)
		if err != nil {
			resp.Error(fmt.Sprint("json file ", json_path, " not found"))
			return
		}
		surl_command.List()
	})
	app.Route([]string{"run"}, func(c *cli.Cli) {
		_, isHelp := c.GetFlagVal([]string{"-h", "--help"})
		if isHelp {
			surl_command.RunHelp()
		}
		err := surl_command.ReadJson(json_path)
		if err != nil {
			resp.Error(fmt.Sprint("json file ", json_path, " not found"))
			return
		}
		if len(c.Commands) <= 1 {
			resp.Println("name not defined")
			surl_command.RunHelp()
			resp.Exit()
		}
		_, IsVerbose := c.GetFlagVal([]string{"-v"})
		err = surl_command.Run(c.Commands[1], IsVerbose)
		if err != nil {
			resp.Error(err.Error())
			return
		}

	})
	app.Route([]string{"help"}, func(c *cli.Cli) {
		surl_command.HelpMessage()
	})
	app.RouteDefault(func(c *cli.Cli) {
		surl_command.HelpMessage()
	})
	app.Run()
}

func main() {
	surl_requests := command.NewSurl()
	resp := utils.Response{}
	RunCommand(surl_requests, &resp, os.Args)
}
