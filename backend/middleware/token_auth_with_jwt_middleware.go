package middleware

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TokenAuthWithJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromAuthorizationHeader(c.GetHeader("Authorization"))
		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, err
			}
			_, publicKey, err := readJwtTokenAuthKeys("jwt_token_auth_secret.pem", "jwt_token_auth_public.pem")
			if err != nil {
				return nil, err
			}
			return publicKey, nil
		})
		if err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("cannot assert jwt.MapClaims")
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}
		if claims["sub"] != "mayukorin" {
			fmt.Println("invalid username")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		c.Next()
	}
}

func readJwtTokenAuthKeys(secretKeyFilePath string, publicKeyFilePath string) (*rsa.PrivateKey, *rsa.PublicKey, error) {

	secretBytes, err := os.ReadFile(secretKeyFilePath)
	if err != nil {
		return nil, nil, err
	}
	secretBlock, _ := pem.Decode(secretBytes)
	if secretBlock == nil {
		return nil, nil, errors.New("cannot decode secret bytes.")
	}
	if secretBlock.Type != "PRIVATE KEY" {
		return nil, nil, errors.New("this is not a private key.")
	}
	secretKeyInterface, err := x509.ParsePKCS8PrivateKey(secretBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}
	secretKey, ok := secretKeyInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, errors.New("cannot assert")
	}

	publicBytes, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		return nil, nil, err
	}
	publicBlock, _ := pem.Decode(publicBytes)
	if publicBlock == nil {
		return nil, nil, errors.New("cannot decode public bytes.")
	}
	if publicBlock.Type != "PUBLIC KEY" {
		return nil, nil, errors.New("this is not a public key.")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}
	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, nil, errors.New("cannot assert")
	}

	return secretKey, publicKey, nil
}

func extractTokenFromAuthorizationHeader(authorizationHeader string) (string, error) {
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return "", errors.New("invalid authorizationHeader")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	return token, nil
}
