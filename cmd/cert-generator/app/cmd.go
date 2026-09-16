/*
Copyright 2024 Hurricane1988 Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/hurricane1988/cert-generator/cmd/cert-generator/app/options"
	"github.com/hurricane1988/cert-generator/pkg/certificate"
	"github.com/hurricane1988/cert-generator/pkg/utils"
)

// helpTemplate defines the custom help output format for the root command.
const helpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}{{end}}

Usage:
  {{.UseLine}}{{if .HasAvailableSubCommands}} <command>{{end}}{{if .HasExample}}

Examples:
  {{.Example}}{{end}}

{{if .HasAvailableSubCommands}}Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding}} {{.Short}}{{end}}{{end}}{{end}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

Use "{{.CommandPath}} [command] --help" for more information about a command.
`

// NewCommand builds the root cobra command.
func NewCommand() *cobra.Command {
	opts := options.NewOptions()

	cmd := &cobra.Command{
		Use:   "cert-generator",
		Short: "Generate TLS certificates and CA",
		Long: `cert-generator is a tool to generate self-signed CA certificates
and server certificates for TLS communication.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.Validate()
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		SilenceUsage: true,
	}
	cmd.AddCommand(newParseCommand())
	cmd.SetHelpTemplate(helpTemplate)
	opts.AddFlags(cmd)
	return cmd
}

func run(opts *options.Options) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg := certificate.Config{
		Country:       split(opts.Country),
		Organization:  split(opts.Organization),
		Domains:       split(opts.Domains),
		CommonName:    opts.CommonName,
		ValidateYears: opts.ValidateYears,
		CertPath:      opts.CertPath,
	}

	if err := certificate.Generate(cfg); err != nil {
		slog.Error("certificate generation failed", "error", err)
		return err
	}

	utils.Info(opts.Country, opts.Organization, opts.Domains, opts.CertPath, opts.CommonName, opts.ValidateYears)
	return nil
}

func split(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}
