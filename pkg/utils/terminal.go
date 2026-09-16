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

package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/hurricane1988/cert-generator/pkg/certificate"
)

// Global color functions using fatih/color.
var (
	Yellow       = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	YellowItalic = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintFunc()
	Green        = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	Blue         = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	Cyan         = color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
	Red          = color.New(color.FgHiRed, color.Bold).SprintFunc()
	White        = color.New(color.FgWhite).SprintFunc()
	WhiteBold    = color.New(color.FgWhite, color.Bold).SprintFunc()
	RedBold      = color.New(color.FgHiRed, color.Bold).SprintFunc()
	forceDetail  = "yaml"
)

// StyleLightBlue is a table style with thin Unicode box characters and blue borders.
var StyleLightBlue = table.Style{
	Name: "StyleLightBlue",
	Box:  table.StyleBoxLight,
	Color: table.ColorOptions{
		Border:    text.Colors{text.FgBlue},
		Separator: text.Colors{text.FgBlue},
	},
	Format:  table.FormatOptionsDefault,
	Options: table.OptionsDefault,
	Size:    table.SizeOptionsDefault,
	Title:   table.TitleOptionsDefault,
}

// Print the terminal with code word.
func Print() {
	fmt.Println(YellowItalic(`
╭━━━╮╱╱╱╱╭╮╱╱╭━━━╮╱╱╱╱╱╱╱╱╱╱╱╱╱╭╮
┃╭━╮┃╱╱╱╭╯╰╮╱┃╭━╮┃╱╱╱╱╱╱╱╱╱╱╱╱╭╯╰╮
┃┃╱╰╋━━┳┻╮╭╯╱┃┃╱╰╋━━┳━╮╭━━┳━┳━┻╮╭╋━━┳━╮
┃┃╱╭┫┃━┫╭┫┣━━┫┃╭━┫┃━┫╭╮┫┃━┫╭┫╭╮┃┃┃╭╮┃╭╯
┃╰━╯┃┃━┫┃┃╰┳━┫╰┻━┃┃━┫┃┃┃┃━┫┃┃╭╮┃╰┫╰╯┃┃
╰━━━┻━━┻╯╰━╯╱╰━━━┻━━┻╯╰┻━━┻╯╰╯╰┻━┻━━┻╯
`))
}

func Info(country, organization, domains, certPath, commonName string, validateYears int) {
	fmt.Printf(`
--------------------------------------------------------------------------------------------
#   CA Country: %s
#   CA Organization: %s
#   CA Domains: %s
#   Cert Path: %s
#   Common Name: %s
#   Validate Years: %d years
#   CRT: %s
#   Key: %s
--------------------------------------------------------------------------------------------
`, country, organization, domains, certPath, commonName, validateYears, certPath+"/"+"tls.crt", certPath+"/"+"tls.key")
}

// PrintCertTable renders parsed certificate details in a single unified table with blue thin borders.
// Output includes subject, issuer, certificate info, fingerprints, SANs, and public key.
func PrintCertTable(cert *certificate.ParsedCert) {
	t := table.NewWriter()
	t.SetStyle(StyleLightBlue)
	t.Style().Options.SeparateRows = false
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, WidthMin: 24},
		{Number: 2, WidthMax: 60},
	})

	mergeCfg := table.RowConfig{AutoMerge: true}

	// Title
	t.AppendRow(table.Row{YellowItalic("证书解析结果"), YellowItalic("证书解析结果")}, mergeCfg)

	// File path
	t.AppendRow(table.Row{Green("文件: "), White(cert.FilePath)})

	t.AppendSeparator()

	// Subject
	t.AppendRow(table.Row{Yellow("主题信息 (Subject)"), Yellow("主题信息 (Subject)")}, mergeCfg)
	addSubjectRows(t, cert.SubjectCommonName, cert.SubjectCountry, cert.SubjectProvince,
		cert.SubjectLocality, cert.SubjectOrg, cert.SubjectOrgUnit)

	t.AppendSeparator()

	// Issuer
	t.AppendRow(table.Row{Yellow("签发者信息 (Issuer)"), Yellow("签发者信息 (Issuer)")}, mergeCfg)
	addSubjectRows(t, cert.IssuerCommonName, cert.IssuerCountry, cert.IssuerProvince,
		cert.IssuerLocality, cert.IssuerOrg, cert.IssuerOrgUnit)

	t.AppendSeparator()

	// Certificate info
	t.AppendRow(table.Row{Yellow("证书信息"), Yellow("证书信息")}, mergeCfg)
	t.AppendRow(table.Row{Green("版本 (Version)"), White(fmt.Sprintf("%d", cert.Version))})
	t.AppendRow(table.Row{Green("序列号 (Serial Number)"), White(cert.SerialNumber)})
	t.AppendRow(table.Row{Green("签名算法"), White(cert.SignatureAlgorithm)})
	t.AppendRow(table.Row{Green("颁发时间 (Not Before)"), White(cert.NotBefore.Format("2006-01-02 15:04:05"))})

	if cert.IsExpired {
		t.AppendRow(table.Row{Green("过期时间 (Not After)"), RedBold(cert.NotAfter.Format("2006-01-02 15:04:05"))})
		t.AppendRow(table.Row{Green("有效期"), RedBold(fmt.Sprintf("已过期 %d 天", -cert.DaysRemaining))})
	} else {
		t.AppendRow(table.Row{Green("过期时间 (Not After)"), White(cert.NotAfter.Format("2006-01-02 15:04:05"))})
		t.AppendRow(table.Row{Green("有效期"), White(fmt.Sprintf("%d 天", cert.DaysRemaining))})
	}

	if cert.IsCA {
		t.AppendRow(table.Row{Green("CA 证书"), White("是")})
	}

	t.AppendSeparator()

	// Fingerprints
	t.AppendRow(table.Row{Yellow("指纹 (Fingerprints)"), Yellow("指纹 (Fingerprints)")}, mergeCfg)
	t.AppendRow(table.Row{Green("SHA1"), White(cert.SHA1Fingerprint)})
	t.AppendRow(table.Row{Green("SHA256"), White(cert.SHA256Fingerprint)})

	// SAN
	if len(cert.DNSNames) > 0 || len(cert.IPAddresses) > 0 || len(cert.EmailAddresses) > 0 ||
		len(cert.URIs) > 0 || len(cert.KeyUsage) > 0 || len(cert.ExtKeyUsage) > 0 {
		t.AppendSeparator()
		t.AppendRow(table.Row{Yellow("备用名 (SAN)"), Yellow("备用名 (SAN)")}, mergeCfg)
		if len(cert.DNSNames) > 0 {
			t.AppendRow(table.Row{Green("DNS Names"), White(strings.Join(cert.DNSNames, ", "))})
		}
		if len(cert.IPAddresses) > 0 {
			ips := make([]string, len(cert.IPAddresses))
			for i, ip := range cert.IPAddresses {
				ips[i] = ip.String()
			}
			t.AppendRow(table.Row{Green("IP Addresses"), White(strings.Join(ips, ", "))})
		}
		if len(cert.EmailAddresses) > 0 {
			t.AppendRow(table.Row{Green("Email Addresses"), White(strings.Join(cert.EmailAddresses, ", "))})
		}
		if len(cert.URIs) > 0 {
			uris := make([]string, len(cert.URIs))
			for i, u := range cert.URIs {
				uris[i] = u.String()
			}
			t.AppendRow(table.Row{Green("URIs"), White(strings.Join(uris, ", "))})
		}
		if len(cert.KeyUsage) > 0 {
			t.AppendRow(table.Row{Green("Key Usage"), White(strings.Join(cert.KeyUsage, ", "))})
		}
		if len(cert.ExtKeyUsage) > 0 {
			t.AppendRow(table.Row{Green("Ext Key Usage"), White(strings.Join(cert.ExtKeyUsage, ", "))})
		}
	}

	// Public key
	t.AppendSeparator()
	t.AppendRow(table.Row{Yellow("公钥 (Public Key)"), Yellow("公钥 (Public Key)")}, mergeCfg)
	t.AppendRow(table.Row{Green("算法 (Algorithm)"), White(cert.PublicKeyAlgorithm)})
	t.AppendRow(table.Row{Green("强度 (Key Size)"), White(fmt.Sprintf("%d bit", cert.PublicKeyBits))})

	fmt.Println()
	fmt.Println(t.Render())
	fmt.Println()
}

// addSubjectRows appends Subject/Issuer rows (CN, C, ST, L, O, OU) to the given table.
func addSubjectRows(t table.Writer, cn string, country, province, locality, org, orgUnit []string) {
	if cn != "" {
		t.AppendRow(table.Row{Green("CN  (通用名称)"), White(cn)})
	}
	if len(country) > 0 {
		t.AppendRow(table.Row{Green("C   (国家)"), White(strings.Join(country, ", "))})
	}
	if len(province) > 0 {
		t.AppendRow(table.Row{Green("ST  (省份)"), White(strings.Join(province, ", "))})
	}
	if len(locality) > 0 {
		t.AppendRow(table.Row{Green("L   (城市)"), White(strings.Join(locality, ", "))})
	}
	if len(org) > 0 {
		t.AppendRow(table.Row{Green("O   (组织)"), White(strings.Join(org, ", "))})
	}
	if len(orgUnit) > 0 {
		t.AppendRow(table.Row{Green("OU  (部门)"), White(strings.Join(orgUnit, ", "))})
	}
}
