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
	"math/big"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

const (
	DefaultTlsKey = "tls.key"
	DefaultTlsCrt = "tls.crt"
)

type Certificate interface {
	CreateCertificate()
	WriteFile(filepath string, cert *bytes.Buffer) error
}

func NewCertificate(tlsOptions Options) Certificate {
	return &tlsOptions
}

// CreateCertificate 生成TLS证书
func (o *Options) CreateCertificate() {
	Log := log.FromContext(o.ctx)
	var caPEM, serverCertPEM, serverPrivyKeyPEM *bytes.Buffer

	// CA 配置
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2024),
		Subject: pkix.Name{
			Country:      o.Country,
			CommonName:   o.CommonName,
			Organization: o.Organization,
		},
		SignatureAlgorithm:    x509.SHA512WithRSA,
		PublicKeyAlgorithm:    x509.RSA,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(o.ValidateYears, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	// CA私钥
	caPrivKey, err := rsa.GenerateKey(cryptorand.Reader, 4096)
	if err != nil {
		Log.Error(err, "generate key failed",
			"validate-years", o.ValidateYears,
		)
	}

	// 自签名CA证书
	caBytes, err := x509.CreateCertificate(cryptorand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		Log.Error(err, "create certificate failed",
			"validate-years", o.ValidateYears,
		)
	}

	// 将CA证书进行PEM编码
	caPEM = new(bytes.Buffer)
	_ = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	// 服务器证书配置
	cert := &x509.Certificate{
		DNSNames:     o.Domains,
		SerialNumber: big.NewInt(2021),
		Subject: pkix.Name{
			CommonName:   o.CommonName,
			Organization: o.Organization,
		},
		SignatureAlgorithm: x509.SHA512WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		NotBefore:          time.Now(),
		NotAfter:           time.Now().AddDate(o.ValidateYears, 0, 0),
		SubjectKeyId:       []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:        []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:           x509.KeyUsageDigitalSignature,
	}

	// 服务器私钥
	serverPrivKey, err := rsa.GenerateKey(cryptorand.Reader, 4096)
	if err != nil {
		Log.Error(err, "generate server key failed",
			"validate-years", o.ValidateYears,
		)
	}

	// 签署服务器证书
	serverCertBytes, err := x509.CreateCertificate(cryptorand.Reader, cert, ca, &serverPrivKey.PublicKey, caPrivKey)
	if err != nil {
		Log.Error(err, "generate server certificate failed",
			"validate-years", o.ValidateYears,
		)
	}

	// 将服务器证书进行PEM编码
	serverCertPEM = new(bytes.Buffer)
	_ = pem.Encode(serverCertPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: serverCertBytes,
	})

	// 将服务器私钥进行PEM编码
	serverPrivyKeyPEM = new(bytes.Buffer)
	_ = pem.Encode(serverPrivyKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(serverPrivKey),
	})
	err = os.MkdirAll(o.CertPath, os.ModeDir)
	if err != nil {
		Log.Error(err, "failed mkdir cert save path",
			"directory", o.CertPath,
		)
	}

	// 写入证书文件
	err = o.WriteFile(o.CertPath+"/"+DefaultTlsCrt, serverCertPEM)
	if err != nil {
		Log.Error(err, "failed save"+DefaultTlsCrt,
			"directory", o.CertPath,
		)
	}
	// 写入证书秘钥文件
	err = o.WriteFile(o.CertPath+"/"+DefaultTlsKey, serverPrivyKeyPEM)
	if err != nil {
		Log.Error(err, "failed save"+DefaultTlsKey,
			"directory", o.CertPath,
		)
	}
}

// WriteFile 将证书写入文件
func (o *Options) WriteFile(filepath string, cert *bytes.Buffer) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(cert.Bytes())
	if err != nil {
		return err
	}
	return nil
}
