package mysqlrouter

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
)

var (
	client *Client
)

func setupWithoutTLS() {
	var err error
	client, err = New("http://mysql-router-http:8080", "root", "mysql", nil)
	if err != nil {
		panic(err)
	}
}

// TestNewClientWithSkipTLSVerify check a connection with InsecureSkipVerify: true
func TestNewClientWithSkipTLSVerify(t *testing.T) {
	opts := &Options{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	_, err := New("https://mysql-router-https:8443", "root", "mysql", opts)
	assert.NoError(t, err)
}

// TestNewClientWithTLS check a connection with certificates
func TestNewClientWithTLS(t *testing.T) {
	certPath, err := filepath.Abs("test/mysql-router/certs/localhost.crt")
	if err != nil {
		assert.NoError(t, err)
	}
	keyPath, err := filepath.Abs("test/mysql-router/certs/localhost.key")
	if err != nil {
		assert.NoError(t, err)
	}
	caPath, err := filepath.Abs("test/mysql-router/certs/localCA.crt")
	if err != nil {
		assert.NoError(t, err)
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		assert.NoError(t, err)
	}

	caCert, err := ioutil.ReadFile(caPath)
	if err != nil {
		assert.NoError(t, err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	opts := &Options{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	_, err = New("https://mysql-router-https:8443", "root", "mysql", opts)
	assert.NoError(t, err)
}
