package auth

import (
	"crypto/rsa"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open("/Users/tanmoysaha/Works/GoProjects/GoBase/pkg/templates/assets/keys/gobase_private_key.pem")
	if err != nil {
		return nil, err
	}

	defer privateKeyFile.Close()

	privateKeyData, err := io.ReadAll(privateKeyFile)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func getPublicKey() (*rsa.PublicKey, error) {
	publicKeyFile, err := os.Open("/Users/tanmoysaha/Works/GoProjects/GoBase/pkg/templates/assets/keys/gobase_public_key.pem")
	if err != nil {
		return nil, err
	}

	publicKeyData, err := io.ReadAll(publicKeyFile)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func GenerateJWTToken(username string, fullname string) (string, error) {
	privateKey, err := getPrivateKey()
	if err != nil {
		return "", err
	}


	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": username,
		"name": fullname,
		"admin": false,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	publicKey, err := getPublicKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}