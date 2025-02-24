package main

import (
	"fmt"
	"os"

	"github.com/simple-url/surl/command"
	"github.com/simple-url/surl/utils"
)

func RunCommand(surl *command.Surl, cc *utils.CommandCapture) {
	if cc.Command == "" || cc.Command == "help" {
		utils.HelpMessage()
		return
	}

	json_path := "./surl.json"
	overide_path, is_found := cc.GetVal("-p")
	if is_found && overide_path != nil {
		json_path = *overide_path
	}
	if cc.Command == "list" {
		if cc.HasFlag("-h") || cc.HasFlag("--help") {
			surl.ListHelp()
		} else {
			err := surl.ReadJson(json_path)
			if err != nil {
				fmt.Println(fmt.Sprint("json file ", json_path, " not found"))
				os.Exit(1)
				return
			}
			surl.List()
		}
	} else if cc.Command == "run" {
		if cc.HasFlag("-h") || cc.HasFlag("--help") {
			surl.RunHelp()
		} else {
			err := surl.ReadJson(json_path)
			if err != nil {
				fmt.Println(fmt.Sprint("json file ", json_path, " not found"))
				os.Exit(1)
				return
			}
			err = surl.Run(cc.Arg, cc.HasFlag("-v"))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
				return
			}
		}
	} else {
		utils.HelpMessage()
		fmt.Println("invalid command")
		os.Exit(1)
		return
	}
}

func main() {
	// Read args
	argsWithProg := os.Args
	cc := utils.NewCommandCapture()
	cc.Capture(argsWithProg)
	surl_requests := command.NewSurl()
	RunCommand(surl_requests, cc)
}
