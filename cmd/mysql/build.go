package main

import (
	"github.com/squillace/porter-mysql/pkg/mysql"
	"github.com/spf13/cobra"
)

func buildBuildCommand(m *mysql.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Build()
		},
	}
	return cmd
}
