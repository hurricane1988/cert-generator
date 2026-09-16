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

import (
	"bytes"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

const (
	DefaultTlsKey = "tls.key"
	DefaultTlsCrt = "tls.crt"
)

// Config holds all parameters needed to generate a certificate.
type Config struct {
	Country       []string
	Organization  []string
	Domains       []string
	CommonName    string
	ValidateYears int
	CertPath      string
}

// Generate creates a self-signed CA and a server certificate signed by it,
// writing tls.crt and tls.key to cfg.CertPath.
func Generate(cfg Config) error {
	// CA 配置
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2024),
		Subject: pkix.Name{
			Country:      cfg.Country,
			CommonName:   cfg.CommonName,
			Organization: cfg.Organization,
		},
		SignatureAlgorithm:    x509.SHA512WithRSA,
		PublicKeyAlgorithm:    x509.RSA,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(cfg.ValidateYears, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	// CA私钥
	caPrivKey, err := rsa.GenerateKey(cryptorand.Reader, 4096)
	if err != nil {
		slog.Error("generate CA key failed", "error", err)
		return fmt.Errorf("generate CA key: %w", err)
	}

	// 自签名CA证书
	caBytes, err := x509.CreateCertificate(cryptorand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		slog.Error("create CA certificate failed", "error", err)
		return fmt.Errorf("create CA certificate: %w", err)
	}

	// 将CA证书进行PEM编码
	caPEM := new(bytes.Buffer)
	_ = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	// 服务器证书配置
	cert := &x509.Certificate{
		DNSNames:     cfg.Domains,
		SerialNumber: big.NewInt(2021),
		Subject: pkix.Name{
			CommonName:   cfg.CommonName,
			Organization: cfg.Organization,
		},
		SignatureAlgorithm: x509.SHA512WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		NotBefore:          time.Now(),
		NotAfter:           time.Now().AddDate(cfg.ValidateYears, 0, 0),
		SubjectKeyId:       []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:        []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:           x509.KeyUsageDigitalSignature,
	}

	// 服务器私钥
	serverPrivKey, err := rsa.GenerateKey(cryptorand.Reader, 4096)
	if err != nil {
		slog.Error("generate server key failed", "error", err)
		return fmt.Errorf("generate server key: %w", err)
	}

	// 签署服务器证书
	serverCertBytes, err := x509.CreateCertificate(cryptorand.Reader, cert, ca, &serverPrivKey.PublicKey, caPrivKey)
	if err != nil {
		slog.Error("sign server certificate failed", "error", err)
		return fmt.Errorf("sign server certificate: %w", err)
	}

	// 将服务器证书进行PEM编码
	serverCertPEM := new(bytes.Buffer)
	_ = pem.Encode(serverCertPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: serverCertBytes,
	})

	// 将服务器私钥进行PEM编码
	serverPrivyKeyPEM := new(bytes.Buffer)
	_ = pem.Encode(serverPrivyKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(serverPrivKey),
	})

	if err := os.MkdirAll(cfg.CertPath, 0755); err != nil {
		slog.Error("create cert directory failed", "error", err, "path", cfg.CertPath)
		return fmt.Errorf("mkdir %s: %w", cfg.CertPath, err)
	}

	// 写入证书文件
	crtPath := filepath.Join(cfg.CertPath, DefaultTlsCrt)
	if err := writeFile(crtPath, serverCertPEM); err != nil {
		slog.Error("save certificate failed", "error", err, "path", crtPath)
		return fmt.Errorf("write %s: %w", crtPath, err)
	}
	slog.Info("certificate saved", "path", crtPath)

	// 写入证书秘钥文件
	keyPath := filepath.Join(cfg.CertPath, DefaultTlsKey)
	if err := writeFile(keyPath, serverPrivyKeyPEM); err != nil {
		slog.Error("save private key failed", "error", err, "path", keyPath)
		return fmt.Errorf("write %s: %w", keyPath, err)
	}
	slog.Info("private key saved", "path", keyPath)

	return nil
}

// writeFile 将证书写入文件
func writeFile(path string, cert *bytes.Buffer) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(cert.Bytes())
	return err
}
