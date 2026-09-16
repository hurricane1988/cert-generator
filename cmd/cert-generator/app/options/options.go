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

package options

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Options holds all configurable parameters for cert-generator.
type Options struct {
	Country       string
	Organization  string
	Domains       string
	CertPath      string
	CommonName    string
	ValidateYears int
}

// NewOptions returns Options with default values, overridden by environment variables.
func NewOptions() *Options {
	return &Options{
		Country:       envOrDefault("CERT_COUNTRY", "CN"),
		Organization:  os.Getenv("CERT_ORGANIZATION"),
		Domains:       os.Getenv("CERT_DOMAINS"),
		CertPath:      envOrDefault("CERT_PATH", "/tmp"),
		CommonName:    os.Getenv("CERT_COMMON_NAME"),
		ValidateYears: envOrDefaultInt("CERT_YEARS", 50),
	}
}

// AddFlags binds Options fields to cobra flags.
func (o *Options) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.Country, "country", o.Country, "The country of CA, multiple items separated by ','. Default: CN.")
	cmd.Flags().StringVar(&o.Organization, "organization", o.Organization, "The organization of CA, multiple items separated by ','.")
	cmd.Flags().StringVar(&o.Domains, "domains", o.Domains, "The domains of the certificate, multiple items separated by ','.")
	cmd.Flags().StringVar(&o.CertPath, "cert-path", o.CertPath, "The path to save certificate.")
	cmd.Flags().StringVar(&o.CommonName, "common-name", o.CommonName, "The common name of CA.")
	cmd.Flags().IntVar(&o.ValidateYears, "years", o.ValidateYears, "The validate years of the certificate. Default: 50.")
}

// Validate checks that required fields are set.
func (o *Options) Validate() error {
	if o.CommonName == "" {
		return fmt.Errorf("--common-name is required")
	}
	if o.ValidateYears < 1 {
		return fmt.Errorf("--years must be at least 1")
	}
	return nil
}

func envOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envOrDefaultInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		var n int
		if _, err := fmt.Sscanf(v, "%d", &n); err == nil {
			return n
		}
	}
	return def
}
