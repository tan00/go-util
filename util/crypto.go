package util

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

var ErrParaInValid = errors.New("invalid parameter")
var ErrPkNotBase64 = errors.New("pk is not base64 format")
var ErrPkInValid = errors.New("invalid public key")
var ErrSignNotBase64 = errors.New("signature is not base64 format")
var ErrVerifyFailed = errors.New("signature not match to data")

func RsaVerifySha1(b64Derpk string, data []byte, b64signature string) (err error) {
	var (
		sign  []byte
		derpk []byte
		pk    *rsa.PublicKey
	)
	if len(b64Derpk) <= 24 || data == nil {
		return ErrParaInValid
	}

	if derpk, err = base64.StdEncoding.DecodeString(b64Derpk); err != nil {
		return ErrPkNotBase64
	}

	hash := sha1.Sum(data)

	if sign, err = base64.StdEncoding.DecodeString(b64signature); err != nil {
		return ErrSignNotBase64
	}

	//去掉java生成公钥的oid标识符
	derpkNoOid := derpk[24:]
	if pk, err = x509.ParsePKCS1PublicKey(derpkNoOid); err != nil {
		return ErrPkInValid
	}

	err = rsa.VerifyPKCS1v15(pk, crypto.SHA1, hash[:], sign)
	if err != nil {
		return
	}
	return
}

//https://wenku.baidu.com/view/58ba863610661ed9ad51f390.html
// 302130 0906052B0E03021A05000414

func TrimJavaPkOid(b64pk string) (b64DerPk string) {
	return
}
