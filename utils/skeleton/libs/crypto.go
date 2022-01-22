package libs

import (
	"crypto/x509"
	"io/ioutil"
)

// LoadCACertPool reads cert files from paths, and creates cert pool
func LoadCACertPool(certPaths []string) (caCertPool *x509.CertPool, err error) {
	caCertPool = x509.NewCertPool()
	for _, path := range certPaths {
		var caCert []byte
		caCert, err = ioutil.ReadFile(path)
		if err != nil {
			return
		}
		caCertPool.AppendCertsFromPEM(caCert)
	}
	return caCertPool, nil
}
