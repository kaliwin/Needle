package certificate

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/mail"
	"net/url"
	"time"
)

// MakePemCert 证书生成器
func MakePemCert(cACert CACert, hosts []string, CommonName string, subject pkix.Name) (cert []byte, privacy []byte, er error) {
	priv, err := GeneratePrivateKey(Rsa2048) // 生成私钥
	if err != nil {
		return nil, nil, err
	}

	subject.CommonName = CommonName
	pub := priv.(crypto.Signer).Public() // 生成公钥
	expiration := time.Now().AddDate(2, 3, 0)
	tpl := &x509.Certificate{
		SerialNumber: RandomSerialNumber(),
		Subject:      subject,
		NotBefore:    time.Now(), NotAfter: expiration,
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			tpl.IPAddresses = append(tpl.IPAddresses, ip)
		} else if email, err := mail.ParseAddress(h); err == nil && email.Address == h {
			tpl.EmailAddresses = append(tpl.EmailAddresses, h)
		} else if uriName, err := url.Parse(h); err == nil && uriName.Scheme != "" && uriName.Host != "" {
			tpl.URIs = append(tpl.URIs, uriName)
		} else {
			tpl.DNSNames = append(tpl.DNSNames, h)
		}
	}

	c, err := x509.CreateCertificate(rand.Reader, tpl, cACert.CaCert, pub, cACert.CaKey) // 签发证书
	if err != nil {
		return nil, nil, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: c})
	privDER, err := x509.MarshalPKCS8PrivateKey(priv)
	privPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privDER})
	return certPEM, privPEM, nil
}

type CACert struct {
	CaCert *x509.Certificate
	CaKey  crypto.PrivateKey
}

// LoadPemCA 加载CA证书 pem 格式
func LoadPemCA(Cert []byte, Key []byte) (ca CACert, er error) {

	certDERBlock, _ := pem.Decode(Cert) // 解码
	keyDERBlock, _ := pem.Decode(Key)   // 解码

	caCert, err := x509.ParseCertificate(certDERBlock.Bytes)   // 解析
	caKey, err := x509.ParsePKCS8PrivateKey(keyDERBlock.Bytes) // 解析

	return CACert{CaCert: caCert, CaKey: caKey}, err
}

// LoadCA 加载CA证书 cer
func LoadCA(Cert []byte, Key []byte) (ca CACert, er error) {
	caCert, err := x509.ParseCertificate(Cert)   // 解析
	caKey, err := x509.ParsePKCS8PrivateKey(Key) // 解析
	return CACert{CaCert: caCert, CaKey: caKey}, err
}

// MakeCA 生成CA证书
func MakeCA(k CryptoType, subject pkix.Name) ([]byte, []byte, error) {
	priv, err := GeneratePrivateKey(k) // 生成私钥
	if err != nil {
		return nil, nil, err
	}
	pub := priv.(crypto.Signer).Public() // 生成公钥

	pkixPublicKey, err := x509.MarshalPKIXPublicKey(pub) // 生成公钥PKIX格式
	if err != nil {
		return nil, nil, err
	}
	var spki struct {
		Algorithm        pkix.AlgorithmIdentifier
		SubjectPublicKey asn1.BitString
	}
	_, err = asn1.Unmarshal(pkixPublicKey, &spki) // 解析公钥
	skid := sha1.Sum(spki.SubjectPublicKey.Bytes) // 生成公钥指纹

	tpl := &x509.Certificate{
		SerialNumber: RandomSerialNumber(),
		Subject:      subject,
		SubjectKeyId: skid[:],

		NotAfter:  time.Now().AddDate(10, 0, 0),
		NotBefore: time.Now(),

		KeyUsage: x509.KeyUsageCertSign,

		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLenZero:        true,
	}

	// 创造证书
	cert, err := x509.CreateCertificate(rand.Reader, tpl, tpl, pub, priv)
	// 生成私钥DER格式
	privDER, err := x509.MarshalPKCS8PrivateKey(priv)
	return pem.EncodeToMemory(
			&pem.Block{Type: "CERTIFICATE", Bytes: cert}),
		pem.EncodeToMemory(
			&pem.Block{Type: "PRIVATE KEY", Bytes: privDER}),
		err
}

// GeneratePrivateKey 生成私钥
func GeneratePrivateKey(k CryptoType) (crypto.PrivateKey, error) {
	switch k {
	case ECDSA:
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case Rsa3072:
		return rsa.GenerateKey(rand.Reader, 3072)
	case Rsa2048:
		return rsa.GenerateKey(rand.Reader, 2048)
	}
	return nil, fmt.Errorf("unknown key type: %d", k)
}

type CryptoType int

const (
	ECDSA   CryptoType = 11
	Rsa3072 CryptoType = 12
	Rsa2048 CryptoType = 13
)

// GetTlsCASubject 获取CA证书的主题
func GetTlsCASubject() pkix.Name {
	return pkix.Name{
		Organization:       []string{"MorePossibility"},
		OrganizationalUnit: []string{"MorePossibility CA"},
		CommonName:         "MorePossibility",
	}
}
func fatalIfErr(err error, msg string) {
	if err != nil {
		log.Fatalf("ERROR: %s: %s", msg, err)
	}
}

// CACerToPemToByte pem 编码
func CACerToPemToByte(ca []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: ca})
}

// CAPrivCerToByte pem 编码
func CAPrivCerToByte(priv []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: priv})
}

// RandomSerialNumber 生产随机序列号
func RandomSerialNumber() *big.Int {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	fatalIfErr(err, "failed to generate serial number")
	return serialNumber
}
