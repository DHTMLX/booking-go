package main

import (
	"crypto/ed25519"
	"net/http"

	"github.com/pascaldekloe/jwt"

	"fmt"
	"time"
)

var dID int
var JWTPrivateKey ed25519.PrivateKey
var JWTPublicKey ed25519.PublicKey

func init() {
	JWTPrivateKey = ed25519.NewKeyFromSeed([]byte("LAvp6NoR4M5e0vwO3n2ppHuirpsClNbV"))
	JWTPublicKey = []byte(JWTPrivateKey)[32:]
}

func createUserToken(id int, device int) ([]byte, error) {
	var claims jwt.Claims
	claims.Subject = "user"
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Set = map[string]interface{}{"id": id, "device": device}
	return claims.EdDSASign(JWTPrivateKey)
}

func verifyUserToken(token []byte) (int, int, error) {
	claims, err := jwt.EdDSACheck(token, JWTPublicKey)
	if err != nil {
		return 0, 0, err
	}
	if !claims.Valid(time.Now()) {
		return 0, 0, fmt.Errorf("credential time constraints exceeded")
	}
	if claims.Subject != "user" {
		return 0, 0, fmt.Errorf("wrong claims subject")
	}

	id, ok := claims.Set["id"].(float64)
	if !ok {
		return 0, 0, fmt.Errorf("wrong data in the token")
	}
	device, ok := claims.Set["device"].(float64)
	if !ok {
		return 0, 0, fmt.Errorf("wrong data in the token")
	}

	return int(id), int(device), nil
}

func newDeviceID() int {
	dID += 1
	return dID
}

func getDeviceID(r *http.Request) int {
	v := r.Context().Value("device_id")
	asInt, _ := v.(int)
	return asInt
}
