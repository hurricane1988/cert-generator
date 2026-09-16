## Cert Generator

[![Go](https://img.shields.io/badge/Go-1.27.1-blue.svg)](https://go.dev)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)

---

### Description

---

`cert-generator` is a command-line tool for generating self-signed CA certificates and server certificates. It also provides a `parse` subcommand to inspect and display SSL/TLS certificate details in a formatted table view.

### Getting Started

---

#### Build

```shell
make build
```

#### Usage

```shell
# Show help
bin/cert-generator --help

# Generate a certificate
bin/cert-generator \
  --country CN \
  --organization "My Org" \
  --domains "example.com,www.example.com" \
  --common-name "example.com" \
  --years 10 \
  --cert-path /tmp/certs

# Parse and display certificate details
bin/cert-generator parse /tmp/certs/tls.crt
```
```shell
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                                       证书解析结果                                      │
│ 文件:                    │ /tmp/certs/tls.crt                                           │
├──────────────────────────┴──────────────────────────────────────────────────────────────┤
│                                    主题信息 (Subject)                                   │
│ CN  (通用名称)           │ example.com                                                  │
│ O   (组织)               │ My Org                                                       │
├──────────────────────────┴──────────────────────────────────────────────────────────────┤
│                                   签发者信息 (Issuer)                                   │
│ CN  (通用名称)           │ example.com                                                  │
│ C   (国家)               │ CN                                                           │
│ O   (组织)               │ My Org                                                       │
├──────────────────────────┴──────────────────────────────────────────────────────────────┤
│                                         证书信息                                        │
│ 版本 (Version)           │ 3                                                            │
│ 序列号 (Serial Number)   │ 07:E5                                                        │
│ 签名算法                 │ SHA512-RSA                                                   │
│ 颁发时间 (Not Before)    │ 2026-09-16 02:35:43                                          │
│ 过期时间 (Not After)     │ 2036-09-16 02:35:43                                          │
│ 有效期                   │ 3652 天                                                      │
├──────────────────────────┴──────────────────────────────────────────────────────────────┤
│                                   指纹 (Fingerprints)                                   │
│ SHA1                     │ 1C:E9:89:6F:52:95:C3:10:3A:6A:7A:9B:E2:63:DA:7D:C7:7B:AA:DC  │
│ SHA256                   │ C8:CE:F4:FE:17:D9:8A:08:EF:88:02:0A:20:8A:2F:A7:E1:06:6A:FD: │
│                          │ 6D:BB:72:E3:3B:44:C4:2A:85:E5:C6:2D                          │
├──────────────────────────┴──────────────────────────────────────────────────────────────┤
│                                       备用名 (SAN)                                      │
│ DNS Names                │ example.com, www.example.com                                 │
│ Key Usage                │ Digital Signature                                            │
│ Ext Key Usage            │ Client Authentication, Server Authentication                 │
├──────────────────────────┴──────────────────────────────────────────────────────────────┤
│                                    公钥 (Public Key)                                    │
│ 算法 (Algorithm)         │ RSA                                                          │
│ 强度 (Key Size)          │ 4096 bit                                                     │
└──────────────────────────┴──────────────────────────────────────────────────────────────┘
```
