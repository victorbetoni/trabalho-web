package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func GenerateToken(identifier string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user": identifier,
		"nbf":  time.Date(2024, 11, 18, 0, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func ExtractUserIdentifier(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%v", claims["user_email"]), nil
	}

	return "", errors.New("unauthorized")
}

func DefineKeyPair(priv []byte, pub []byte) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		log.Fatalf("Error while loading PRIVATE key: %s\n", err)
	}
	signKey = privateKey
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		log.Fatalf("Error while loading PUBLIC key: %s\n", err)
	}
	verifyKey = publicKey
}
