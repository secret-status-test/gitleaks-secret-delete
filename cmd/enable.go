package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/zricethezav/gitleaks/v8/ucmp"
	"strconv"
)

func init() {
	enableCmd.Flags().String("url", "", "Backend URL")
	enableCmd.Flags().Bool("debug", false, "Enable debug output")
	rootCmd.AddCommand(enableCmd)
}

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable gitleaks in pre-commit script",
	Run:   runEnable,
}

func runEnable(cmd *cobra.Command, args []string) {
	// Setting .git/config : Gitleaks.url
	urlFlag, _ := cmd.Flags().GetString("url")
	ucmp.SetGitleaksConfig(ucmp.ConfigUrl, urlFlag)

	debugFlag, _ := cmd.Flags().GetBool("debug")
	if debugFlag {
		// If enable command with --debug flag, set Gitleaks.debug to true
		// Using this flag, print the all commands logs
		ucmp.SetGitleaksConfig(ucmp.ConfigDebug, strconv.FormatBool(debugFlag))
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Setting .git/config : Gitleaks.enable
	ucmp.SetGitleaksConfig(ucmp.ConfigEnable, "true")

	// Setting .git/hooks/pre-commit
	ucmp.EnableGitHooks(ucmp.PreCommitScriptPath, ucmp.PreCommitScript)

	// Setting .git/hooks/post-commit
	ucmp.EnableGitHooks(ucmp.PostCommitScriptPath, ucmp.PostCommitScript)

	log.Debug().Msg("Gitleaks Enabled")
}
