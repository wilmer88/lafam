// Package mustcert generates certificates for TLS testing
package mustcert

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"sync"
	"time"
)

// Cert is an alias for tls.Certificate with extra helper methods.
type Cert tls.Certificate

// CA generates a new certificate that can sign leaf & intermediary certificates. The certificate is self-signed if parent is nil.
func CA(commonName string, parent *Cert) *Cert {
	cert := newCert(commonName)
	cert.Leaf.IsCA = true

	if parent == nil {
		return cert.Sign(cert)
	}
	return parent.Sign(cert)
}

// Leaf generates a new leaf certificate. The certificate is self-signed if
// parent is nil. If opts are provided, they are invoked on the certificate
// before it's signed
func Leaf(commonName string, parent *Cert, opts ...func(*x509.Certificate)) *Cert {
	cert := newCert(commonName, opts...)
	if parent == nil {
		return cert.Sign(cert)
	}
	return parent.Sign(cert)
}

func newCert(commonName string, opts ...func(*x509.Certificate)) *Cert {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	serial := nextSerial()
	x509Cert := &x509.Certificate{
		BasicConstraintsValid: true,
		SubjectKeyId:          big.NewInt(serial).Bytes(),
		SerialNumber:          big.NewInt(serial),
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore:   time.Now().Add(-5 * time.Minute),
		NotAfter:    time.Now().Add(5 * time.Minute),
		DNSNames:    []string{commonName},
		IPAddresses: []net.IP{net.ParseIP("0.0.0.0"), net.ParseIP("127.0.0.1"), net.ParseIP("::")},
	}

	for _, o := range opts {
		o(x509Cert)
	}

	return &Cert{
		PrivateKey: privateKey,
		Leaf:       x509Cert,
	}
}

// Sign returns a new Cert with an additional signature signed by c.
func (c *Cert) Sign(child *Cert) *Cert {
	pubKey := child.PrivateKey.(*ecdsa.PrivateKey).Public()
	cert, err := x509.CreateCertificate(rand.Reader, child.Leaf, c.Leaf, pubKey, c.PrivateKey)
	if err != nil {
		panic(err)
	}
	leaf, err := x509.ParseCertificate(cert)
	if err != nil {
		panic(err)
	}

	return &Cert{
		Certificate: [][]byte{cert},
		PrivateKey:  child.PrivateKey,
		Leaf:        leaf,
	}
}

// CertMap is a map holding the PEM encoded cert & key.
func (c *Cert) CertMap() map[string]string {
	return map[string]string{
		"cert": c.CertPEM(),
		"key":  c.KeyPEM(),
	}
}

// CertPEM is the PEM encoded x509 certificate data.
func (c *Cert) CertPEM() string {
	b := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: []byte{},
	}

	for _, derBytes := range c.Certificate {
		b.Bytes = append(b.Bytes, derBytes...)
	}

	return string(pem.EncodeToMemory(b))
}

// KeyPEM is the PEM encoded private key data.
func (c *Cert) KeyPEM() string {
	buf, err := x509.MarshalECPrivateKey(c.PrivateKey.(*ecdsa.PrivateKey))
	if err != nil {
		panic(err)
	}

	b := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: buf,
	}

	return string(pem.EncodeToMemory(b))
}

// TLS returns c as a *tls.Certificate.
func (c *Cert) TLS() *tls.Certificate {
	return (*tls.Certificate)(c)
}

var (
	serialMu  sync.Mutex
	serialNum int64
)

func nextSerial() int64 {
	serialMu.Lock()
	defer serialMu.Unlock()

	serialNum++
	return serialNum
}
