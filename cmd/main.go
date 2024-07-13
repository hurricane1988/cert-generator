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

package main

import (
	"flag"
	"github.com/hurricane1988/cert-generator/pkg/certificate"
	"github.com/hurricane1988/cert-generator/pkg/utils"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"strings"
)

var (
	setupLog = ctrl.Log.WithName("setup")
)

func main() {
	var (
		country       string
		organization  string
		certPath      string
		validateYears int
		domains       string
		commonName    string
	)
	flag.StringVar(&country, "ca-country", "CN", "The country of CA, multiple items separated by ',', Default: CN.")
	flag.StringVar(&organization, "ca-organization", "", "The organization of CA, multiple items separated by ','")
	flag.StringVar(&domains, "ca-domains", "", "The domain of CA, multiple items separated by ','.")
	flag.StringVar(&certPath, "cert-path", "/tmp", "The path to save certificate.")
	flag.StringVar(&commonName, "ca-common-name", "", "The common name of CA.")
	flag.IntVar(&validateYears, "ca-years", 50, "The validate years of CA, Default: 50.")
	opts := zap.Options{
		Development: false,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
	certificate.NewCertificate(certificate.Options{
		Country:       strings.Split(country, ","),
		CertPath:      certPath,
		Organization:  strings.Split(organization, ","),
		ValidateYears: validateYears,
		Domains:       strings.Split(domains, ","),
		CommonName:    commonName,
	}).CreateCertificate()

	// 打印终端
	// utils.Print()
	utils.Info(country, organization, domains, certPath, commonName, validateYears)
}
