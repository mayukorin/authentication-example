package handlers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !(username == "mayukorin" && password == "password") {
		c.AbortWithError(http.StatusUnauthorized, errors.New("not correct username and password"))
		return
	}

	secretKey, _, err := readJwtTokenAuthKeys("jwt_token_auth_secret.pem", "jwt_token_auth_public.pem")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	nowTime := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": username,
		"iat": nowTime,
		"eat": nowTime.AddDate(0, 0, 1),
	})
	fmt.Println(token.Method)
	SignedToken, err := token.SignedString(secretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt_token": SignedToken,
	})
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
