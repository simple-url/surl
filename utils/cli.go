package utils

import (
	"regexp"
	"strings"
)

type CommandCapture struct {
	Command     string
	Arg         string
	Flags       []string
	FlagsKeyVal map[string]*string
}

func IsFlag(text string) bool {
	isFlagShort, _ := regexp.MatchString("-.*", text)
	isFlagLong, _ := regexp.MatchString("--.*", text)
	return isFlagShort || isFlagLong
}

func NewCommandCapture() *CommandCapture {
	return &CommandCapture{
		Command:     "",
		Arg:         "",
		Flags:       []string{},
		FlagsKeyVal: map[string]*string{},
	}
}

func (f *CommandCapture) Capture(args []string) {
	if len(args) <= 1 {
		return
	}

	for idx, arg := range args {
		// Check is Command
		if f.Command == "" && idx == 1 {
			f.Command = strings.TrimSpace(arg)
		}

		// Check is Args
		if idx == 2 && !IsFlag(arg) {
			f.Arg = strings.TrimSpace(arg)
		}

		// Check is Flag
		isFlag := IsFlag(arg)
		if isFlag {
			f.Flags = append(f.Flags, strings.TrimSpace(arg))
		}

		// Check is Flag a key val
		if isFlag && (idx+1 < len(args)) {
			if IsFlag(args[idx+1]) {
				f.FlagsKeyVal[arg] = nil
			} else {
				val := strings.TrimSpace(args[idx+1])
				f.FlagsKeyVal[arg] = &val
			}
		}
	}
}

func (f *CommandCapture) HasFlag(flag string) bool {
	for _, item := range f.Flags {
		if item == flag {
			return true
		}
	}
	return false
}

func (f *CommandCapture) GetVal(flag string) (*string, bool) {
	val, is_found := f.FlagsKeyVal[flag]
	return val, is_found
}
