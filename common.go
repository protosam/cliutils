package cliutils

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

// ReqArgCount is a helper function for urfave/cli/v2
// It is used as an alias to check if a required number of arguments is passed
// to the command via the cli context.
//
// The reason to use this is to standardize the exit code as well as the error
// output. Having the common string in one place is better than manually writing
// this small condition and string many times across a code-base.
func ReqArgCount(c *cli.Context, required_arg_count int) error {
	if c.NArg() < required_arg_count {
		exitstr := fmt.Sprintf("command requires at least %d arguments (see --help for more information)", required_arg_count)
		return cli.Exit(exitstr, 1)
	}
	return nil
}

// HelpFmt handles formatting help output for things such as descriptions.
//
// 1. Shared indentation in the help string that may occur from tabbed code is
// dedented.
// 2. Leading and trailing whitespace is trimmed.
// 3. Text is wrapped at 80 characters for that "book-like" feel.
func HelpFmt(s string) string {
	s = Dedent(s)
	s = strings.TrimSpace(s)
	s = Wrap(s, 80)
	return s
}

// Create a cli.App from a cli.Command.
func AppFrom(cmd *cli.Command, version string) *cli.App {
	return &cli.App{
		Name:        cmd.Name,
		Usage:       cmd.Usage,
		Version:     version,
		HelpName:    cmd.HelpName,
		Description: cmd.Description,
		Commands:    cmd.Subcommands,
	}
}
