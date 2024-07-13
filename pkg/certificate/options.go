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

package certificate

import "context"

type Options struct {
	Country       []string `json:"country,omitempty"`
	CertPath      string   `json:"certPath,omitempty"`
	Organization  []string `json:"organization,omitempty"`
	ValidateYears int      `json:"validateYears,omitempty"`
	Domains       []string `json:"domains,omitempty"`
	CommonName    string   `json:"commonName,omitempty"`
	ctx           context.Context
}
