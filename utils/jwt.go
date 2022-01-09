package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func SignToken(tokenData interface{}) (string, error) {
	//load Token Data
	data, dataErr := json.Marshal(tokenData)
	if dataErr != nil {
		return "", fmt.Errorf("Couldn't parse token data: %v", tokenData)
	}

	// parse the JSON of the claims
	var claims jwt.MapClaims
	if err := json.Unmarshal(data, &claims); err != nil {
		return "", fmt.Errorf("Couldn't parse claims JSON: %v", err)
	}

	// get the signing alg
	alg := jwt.GetSigningMethod("HS256")
	if alg == nil {
		return "", fmt.Errorf("Couldn't find signing method: HS256")
	}

	// create a new token
	token := jwt.NewWithClaims(alg, claims)

	if out, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))); err == nil {
		return out, nil
	} else {
		return "", fmt.Errorf("Error signing token: %v", err)
	}

}

func VerifyToken(tokenData string) (interface{}, error) {
	// Parse the token.  Load the key from command line option
	token, err := jwt.Parse(string(tokenData), func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	// Print an error if we can't parse for some reason
	if err != nil {
		return "", fmt.Errorf("Couldn't parse token: %v", err)
	}
	// Print some debug data
	if token != nil {
		return token.Claims, nil
	}
	return "", fmt.Errorf("Couldn't parse token: %v", err)
}
