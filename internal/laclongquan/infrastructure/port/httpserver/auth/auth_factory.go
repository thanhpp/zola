package auth

import (
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
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
		return nil, ErrNilConfig
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
		return nil, ErrInvalidMethod
	}

	return fac.rsaPubKey, nil
}

func (fac authFactory) SignClaims(in *Claims) (string, error) {
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

func (fac authFactory) NewClaimsFromUser(user *entity.User) (*Claims, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        id.String(),
			Issuer:    fac.issuer,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(fac.expireDuration).Unix(),
		},
		User: JWTUserData{
			ID: user.ID().String(),
		},
	}, nil
}

func (fac authFactory) unmarshalMapClaims(mapClaims *jwt.MapClaims) (*Claims, error) {
	dataB, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, err
	}

	claims := new(Claims)
	err = json.Unmarshal(dataB, claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (fac authFactory) NewClaimsFromToken(token string) (*Claims, error) {
	jwtoken, err := jwt.Parse(token, fac.jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	if !jwtoken.Valid {
		return nil, ErrInvalidToken
	}

	mapClaims, ok := jwtoken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrClaimsIsNotMapClaims
	}

	claims, err := fac.unmarshalMapClaims(&mapClaims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
