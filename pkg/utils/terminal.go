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
	"github.com/fatih/color"
)

// Define global color variables.
var (
	Yellow       = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	YellowItalic = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintFunc()
	Green        = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	Blue         = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	Cyan         = color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
	Red          = color.New(color.FgHiRed, color.Bold).SprintFunc()
	White        = color.New(color.FgWhite).SprintFunc()
	WhiteBold    = color.New(color.FgWhite, color.Bold).SprintFunc()
	forceDetail  = "yaml"
)

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
