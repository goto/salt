# salt

[![GoDoc reference](https://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/github.com/goto/salt)
![test workflow](https://github.com/goto/salt/actions/workflows/test.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/goto/salt)](https://goreportcard.com/report/github.com/goto/salt)

Shared libraries used in the gotocompany ecosystem. Use at your own risk. Breaking changes should be anticipated.

## Installation

To use, run the following command:

```
go get github.com/goto/salt
```

## Pacakages

### Audit

Package for adding audit events in your applications.

### Cmdx

Cobra based cli helper which allows adding command groups, provides custom help and usage functions.

```
var cmd = &cli.Command{
	Use:   "exec <command> <subcommand> [flags]",
	SilenceUsage:  true,
	SilenceErrors: true,
	Annotations: map[string]string{
		"group": "core",
		"help:learn": "Learn about the project",
	},
}

cmdx.SetHelp(cmd)
cmd.AddCommand(cmdx.SetCompletionCmd("exec"))
cmd.AddCommand(cmdx.SetHelpTopicCmd("environment", envHelp))
cmd.AddCommand(cmdx.SetHelpTopicCmd("auth", authHelp))
cmd.AddCommand(cmdx.SetRefCmd(cmd))
```

### Config

Viper abstractions which provides functions for loading config files for the application.

### DB

Postgres based database abstractions for creating a client and running migrations.

### Log

Logger for easy application loggging.

### Printer

Command line printer for CLI based applications.

### Server

GRPC based server abstraction.

### Term

Helper functions for working with terminal.

### Version

Helper functions for fetching github latest and outdated releases.
