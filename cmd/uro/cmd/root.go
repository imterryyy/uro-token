package cmd

import (
  "github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
  rootCmd := &cobra.Command{
    Use: "uro",
    Short: "Uro Token",
  }

  return rootCmd
}
