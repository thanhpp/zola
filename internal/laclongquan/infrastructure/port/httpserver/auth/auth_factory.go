package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type authFactory struct {
	issuer         string
	expireDuration time.Duration
	rsaKeyName     string
	rsaPubKey      *rsa.PublicKey
	rsaPriKey      *rsa.PrivateKey
}

func newFactoryFromConfig(cfg *Config) (*authFactory, error) {
	if cfg == nil {
		return nil, errors.New("nil config")
	}

	// read public
	pubF, err := ioutil.ReadFile(cfg.RSAPubKeyPath)
	if err != nil {
		return nil, err
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubF)
	if err != nil {
		return nil, err
	}

	// read private
	priF, err := ioutil.ReadFile(cfg.RSAPriKeyPath)
	if err != nil {
		return nil, err
	}

	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(priF)
	if err != nil {
		return nil, err
	}

	return &authFactory{
		issuer:         cfg.Issuer,
		expireDuration: time.Hour * time.Duration(cfg.ExpireDuration),
		rsaKeyName:     cfg.RSAKeyName,
		rsaPubKey:      pubKey,
		rsaPriKey:      priKey,
	}, nil
}

func (fac authFactory) jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	// validate alg
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, errors.New("invalid signing method")
	}

	return fac.rsaPubKey, nil
}

func (fac authFactory) signJWT(in *JWT) (string, error) {
	mapClaims, err := in.exportMapClaims()
	if err != nil {
		return "", err
	}

	newJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, mapClaims)
	// add key id
	newJWT.Header["kid"] = fac.rsaKeyName

	token, err := newJWT.SignedString(fac.rsaPriKey)
	if err != nil {
		return "", errors.WithMessage(err, "signed string")
	}

	return token, nil
}
