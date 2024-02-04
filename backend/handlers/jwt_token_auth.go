package handlers

import (
	"bytes"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": nowTime,
		"eat": nowTime.AddDate(0, 0, 1),
	})
	SignedToken, err := token.SignedString(secretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("cannot sign token"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jwt token": SignedToken,
	})
}

func readJwtTokenAuthKeys(secretKeyFilePath string, publicKeyFilePath string) (*rsa.PrivateKey, *rsa.PublicKey, error) {

	// secretBytes, err := ioutil.ReadFile(secretKeyFilePath)
	// secretBytes := []byte("-----BEGIN PRIVATE KEY-----")
	secretBytes := []byte(`-----BEGIN PRIVATE KEY-----
	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
	hYOk9GXnJgXYF+Il2vkv4ujjFhjZPP8=
	-----END PRIVATE KEY-----`)
	// if err != nil {
	// 	return nil, nil, err
	// }
	var pemStart = []byte("\n-----BEGIN ")
	var pemEnd = []byte("\n-----END ")

	fmt.Println(pemStart[1:])
	fmt.Println(secretBytes)
	fmt.Println(bytes.HasPrefix(secretBytes, pemStart[1:]))
	fmt.Println(pemEnd)
	fmt.Println(bytes.Index(secretBytes, pemEnd))

	secretBlock, _ := pem.Decode(secretBytes)
	fmt.Println(secretBlock)
	if secretBlock == nil {
		return nil, nil, errors.New("cannot decode secret bytes.")
	}
	if secretBlock.Type != "RSA PRIVATE KEY" {
		return nil, nil, errors.New("this is not a private key.")
	}
	secretKey, err := x509.ParsePKCS1PrivateKey(secretBlock.Bytes)
	if err != nil {
		return nil, nil, err
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
