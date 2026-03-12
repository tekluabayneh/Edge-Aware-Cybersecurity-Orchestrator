package utils

import (
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"time"
)

func Generate2FASetup(email string) (string, string, []byte, error) {
	// Generate a new TOTP key (secret)
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "edge-aware",
		AccountName: email,
	})
	if err != nil {
		return "", "", []byte{}, err
	}
	SecretKey := key.Secret()
	OtpUrl := key.URL()

	qrPng, err := qrcode.Encode(OtpUrl, qrcode.Medium, 256)
	if err != nil {
		return "", "", []byte{}, err
	}

	return SecretKey, OtpUrl, qrPng, nil
}

func Verify2FACode(secret string, userCode string) (bool, error) {
	valid, err := totp.ValidateCustom(userCode, secret, time.Now(),
		totp.ValidateOpts{
			Period:    30,
			Skew:      1,
			Digits:    6,
			Algorithm: 0,
		})

	return valid, err
}
