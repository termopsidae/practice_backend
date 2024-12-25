package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 公钥PEM格式字符串
const rsaPublicKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAy3gd3Damt3/F4q6F8wsB
u0Wrh71bd3aN/gVNlnTnAYn416efljNh3daEObgoKUKMcKXiscTBDLuCGJ+3J40B
XgKLfz4fc5VsHtuVjvfsxynhAe7eqlr5yJ3aTdoCELRmpvcHMlCHceGaI9mPKElw
fp2C/Ffsv0ZXudGfDBGiGLIoHusggmuYAcBvWJ4pXPL9hEow0X7dJukTw1iAJqfp
R6qhx4kenshOR8T2af8eEHNDLv/f8IAInv78vSGwNSp5Xl9r7vUmjD6YZY0ajBij
N/iXGVpnXOcv1K5P6WwcmWb3hPOqDbORFdGnLaiLj6ufFgr7zX9ubHLa1IdzGzX1
ewIDAQAB
-----END PUBLIC KEY-----
`

// 私钥PEM格式字符串
const rsaPrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAy3gd3Damt3/F4q6F8wsBu0Wrh71bd3aN/gVNlnTnAYn416ef
ljNh3daEObgoKUKMcKXiscTBDLuCGJ+3J40BXgKLfz4fc5VsHtuVjvfsxynhAe7e
qlr5yJ3aTdoCELRmpvcHMlCHceGaI9mPKElwfp2C/Ffsv0ZXudGfDBGiGLIoHusg
gmuYAcBvWJ4pXPL9hEow0X7dJukTw1iAJqfpR6qhx4kenshOR8T2af8eEHNDLv/f
8IAInv78vSGwNSp5Xl9r7vUmjD6YZY0ajBijN/iXGVpnXOcv1K5P6WwcmWb3hPOq
DbORFdGnLaiLj6ufFgr7zX9ubHLa1IdzGzX1ewIDAQABAoIBACnAVCBD9qHgzjJd
rY7Zy1kqSrBn4sT4xtpBLFKbWmuubCmUc+LWWFa1Fhzh9RvCVB8pawMfZJ3APEnh
PCwZexZXX1iU3s9Q8gbC+EWGCkg0B7/lzDsYv/iyq4EeIx5aZlYMiqWAcMqliIOi
uQWDTFou6Bnv0NzPCmFO6lwHokN+rWAyIWiTAP1/4g2hIBpHZnJRIcya4e4C4SR8
HFoTfqvIfzf9kleme/fc1Jovw+kpKHNtc48oTS7/omH1G0yx+3FgL6HxvxeHaj/4
8jbN8AsHGt3oJIvuvFuUWC0vOSNtCUSYvIx1qepuUITqMIVq0QknLX9+uRPu8epb
Zq0G4wECgYEA09o4190syYTMiX7obTZ7csyYJlPtDRblojO77Wj+pvijBphW5Jcq
9JT9WS7HU0O1YLINhnbT6Ls6tDcI5r7o7oiWksULX6Dr0DRkjKAPTmmtVIPXqEvs
XoFntmBuxZXNZuxp8YXCR74QRfFPNhFIhTYV8ySyoLW/2jv+GlL4dnECgYEA9d6r
22bf5mNZ8NzY9QrCdHNTPPERnLlEytCxHdEBRMJUTajI8x2HWgjO/c1VX2rf95Gy
tiHhMPWlMJV85NcDGcWSnR7sCqoE7YPKjLVcgPOfD1GvKx1nsyCzuGvC5HD6ov3s
1wrevO3oxFC0TNIVT2fLyfDnv4AGzY3GZUMIWKsCgYAv3VMQenk/ApEP06uB2Vhk
5JLPvhCF75FsZ1HjXuCCKxTyTYL8XKco4WyoNKh2SIr3UOM7aSeRopOt2e1Z7PJF
ynhyqrBaJ8p/nQbGuvcaWUf/G1ajQJwj8grTqs/8Nk9VHL8HAZiWivu0QcQKzDbd
Wg31hGoTY1z4WqubmPloEQKBgQDoLxR6B8dO6OmppNUZasTBdZrYhQrNxsOpB7UK
5DHffehgwhCqTWthYcofMYX6qpHAWA29I3dmZcNOgwzDiix6bPKMgAQF+hLXPUrM
4APwHqBJtijGfWlNZodxlAoi8nIt879yP+ih38Wdhl0N4qKPLwTquh4P4NYLzWPa
gpfiKwKBgAqOrT3gN7LWmSNAhP7oAKtO8OZ2vx3zrMIxBCZXUaaUgcFWx2kqjyFG
QxbaZZxaIPtKtlNcn9+3iI5TX9V0QYJuSrLOZGfUS5RdDNZ743ZvJI2jMpqx6lo9
q9EMSxNrQ1JqS7DQJ8hldlQuwLxD62O3VpDmVjmYBcnlELLlBcVG
-----END RSA PRIVATE KEY-----
`

// RSAEncrypt 加密数据
func RSAEncrypt(data string, publicKeyPEM string) string {
	if publicKeyPEM == "" {
		publicKeyPEM = rsaPublicKey
	}
	// 解析公钥
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		panic("failed to convert parsed key to RSA public key")
	}

	// 加密数据
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(data))
	if err != nil {
		panic(err)
	}
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded
}

// RSADecrypt 解密数据
func RSADecrypt(ciphertextStr string, privateKeyPEM string) (string, error) {
	if privateKeyPEM == "" {
		privateKeyPEM = rsaPrivateKey
	}
	// 解析私钥
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the private key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextStr)
	if err != nil {
		return "", err
	}
	// 解密数据
	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
