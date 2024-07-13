## Cert Generator

[]()

---

### Description

---
cert-generator is a tool to generate certs and to simplify the process and complexity of generating certificates.

### Getting Started

---

- [x] build the exec binary.
```shell
make build
```

```shell
bin/cert-generator
```
- [x] get help information
```shell
bin/cert-generator --help
```

```shell
Usage of bin/cert-generator:
  -ca-common-name string
        The common name of CA.
  -ca-country string
        The country of CA, multiple items separated by ',', Default: CN. (default "CN")
  -ca-domains string
        The domain of CA, multiple items separated by ','.
  -ca-organization string
        The organization of CA, multiple items separated by ','
  -ca-years int
        The validate years of CA, Default: 50. (default 50)
  -cert-path string
        The path to save certificate. (default "/tmp")
  -kubeconfig string
        Paths to a kubeconfig. Only required if out-of-cluster.
  -zap-devel
        Development Mode defaults(encoder=consoleEncoder,logLevel=Debug,stackTraceLevel=Warn). Production Mode defaults(encoder=jsonEncoder,logLevel=Info,stackTraceLevel=Error)
  -zap-encoder value
        Zap log encoding (one of 'json' or 'console')
  -zap-log-level value
        Zap Level to configure the verbosity of logging. Can be one of 'debug', 'info', 'error', or any integer value > 0 which corresponds to custom debug levels of increasing verbosity
  -zap-stacktrace-level value
        Zap Level at and above which stacktraces are captured (one of 'info', 'error', 'panic').
  -zap-time-encoding value
        Zap time encoding (one of 'epoch', 'millis', 'nano', 'iso8601', 'rfc3339' or 'rfc3339nano'). Defaults to 'epoch'.
```
- [x] generate a cert
```shell
bin/cert-generator -ca-country=china -ca-common-name=shenzhen -ca-domains=kube-system.metrics-server.svc,kube-system.coredns.svc -ca-organization=kubernetes -ca-years=50
```

```shell
--------------------------------------------------------------------------------------------
#   CA Country: china
#   CA Organization: kubernetes
#   CA Domains: kube-system.metrics-server.svc,kube-system.coredns.svc
#   Cert Path: /tmp
#   Common Name: shenzhen
#   Validate Years: 50 years
#   CRT: /tmp/tls.crt
#   Key: /tmp/tls.key
--------------------------------------------------------------------------------------------
```
