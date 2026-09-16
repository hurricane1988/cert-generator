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
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/hurricane1988/cert-generator/pkg/certificate"
	"github.com/hurricane1988/cert-generator/pkg/utils"
)

// newParseCommand creates the "parse" subcommand for displaying SSL certificate details.
func newParseCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "parse <cert-file>",
		Short: "Parse and display SSL certificate details",
		Long: `Read a PEM or DER encoded certificate file, parse it, and display
subject, issuer, certificate info, fingerprints, SANs, and public key details.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runParse(args[0])
		},
		SilenceUsage: true,
	}
}

// runParse parses the certificate file at certPath and prints the result table.
func runParse(certPath string) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cert, err := certificate.Parse(certPath)
	if err != nil {
		slog.Error("certificate parse failed", "error", err)
		return fmt.Errorf("parse %s: %w", certPath, err)
	}

	utils.PrintCertTable(cert)
	return nil
}
