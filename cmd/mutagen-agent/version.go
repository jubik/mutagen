package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/pkg/agent"
	"github.com/mutagen-io/mutagen/pkg/mutagen"
)

func versionMain(command *cobra.Command, arguments []string) error {
	// Print version information.
	fmt.Println(mutagen.Version)

	// Success.
	return nil
}

var versionCommand = &cobra.Command{
	Use:          agent.ModeVersion,
	Short:        "Show version information",
	RunE:         versionMain,
	SilenceUsage: true,
}

var versionConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := versionCommand.Flags()

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&versionConfiguration.help, "help", "h", false, "Show help information")
}
