// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license,
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Cmd struct {
	Command   string
	Arguments []string
	Stdin     io.Reader
	Stdout    io.Writer
	Stderr    io.Writer
	Environ   []string
}

type CmdResult struct {
	Exit int
	Err  error
}

func resolveAlias(arg string) string {

	switch arg {
	case
		"date",
		"time":
		arg = "datetime"

	case "log":
		arg = "journal"

	case "net":
		arg = "network"

	case "host":
		arg = "hostname"

	case "tmp":
		arg = "tmpfiles"

	case "password":
		arg = "ask-password"

	case "esc":
		arg = "escape"

	case "ana":
		arg = "analyze"
	}

	return arg
}

func resolveArgs(isDefault bool, args []string) []string {

	switch len(args) {
	case 2:
		if isDefault {
			return args[1:]
		}
		return []string{}

	default:

		if isDefault {
			return args[1:]
		}
		return args[2:]
	}
}

func run(c Cmd) CmdResult {

	cmd := exec.Command(c.Command, c.Arguments...)
	cmd.Stdin = c.Stdin
	cmd.Stdout = c.Stdout
	cmd.Stderr = c.Stderr
	cmd.Env = c.Environ

	err := cmd.Run()

	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			return CmdResult{
				Exit: e.ExitCode(),
				Err:  err,
			}
		}
	}

	return CmdResult{
		Exit: 0,
		Err:  err,
	}
}

func getCmd(args []string) Cmd {

	cmd := Cmd{
		Stdin:   os.Stdin,
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		Environ: os.Environ(),
	}

	if len(args) >= 2 {

		switch name := resolveAlias(args[1]); name {

		case "complete":

			ccmd := getCmd(args[2:])

			cmd.Command = "/usr/lib/ctl/complete.sh"
			cmd.Arguments = []string{
				strings.Join(append([]string{ccmd.Command}, ccmd.Arguments...), " "),
			}
			return cmd

		case
			"nd",
			"wd",
			"pa",
			"evm",
			"key",
			"bus",
			"boot",
			"login",
			"panel",
			"teamd",
			"kdump",
			"udisks",
			"locale",
			"system",
			"resolve",
			"network",
			"journal",
			"portable",
			"hostname",
			"timedate",
			"coredump":

			cmd.Command = name + "ctl"
			cmd.Arguments = resolveArgs(false, args)
			return cmd

		case
			"run",
			"delta",
			"detect-virt",
			"machine-id-setup",
			"cgls",
			"sysusers",
			"stdio-bridge",
			"socket-activate",
			"tmpfiles",
			"notify",
			"hwdb",
			"umount",
			"firstboot",
			"id128",
			"ask-password",
			"tty-ask-password-agent",
			"cgtop",
			"mount",
			"escape",
			"inhibit",
			"analyze",
			"path":

			cmd.Command = "systemd-" + name
			cmd.Arguments = resolveArgs(false, args)
			return cmd

		case
			"dcat",
			"scat":

			cmd.Command = "systemd-cat"
			cmd.Arguments = resolveArgs(false, args)
			return cmd

		case
			"dresolve",
			"sresolve":

			cmd.Command = "systemd-resolve"
			cmd.Arguments = resolveArgs(false, args)
			return cmd
		}

	}

	cmd.Command = "systemctl"
	cmd.Arguments = resolveArgs(true, args)

	return cmd
}

func main() {

	res := run(getCmd(os.Args))

	if res.Err != nil {
		log.Print(res.Err)
	}

	os.Exit(res.Exit)
}
