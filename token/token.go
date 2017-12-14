/* token.go
Simple examples how to use JSON Web Token (JWT) with Go using jwt-go library.
Struktur JWT terdiri dari Header, payload (claims), dan signature,
yang dipisahkan oleh . (period)

This example using rsa asymmetric key to sign the token.
And working example to create go package
*/
package token

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Custom claims contains just username string and standard claims
type CustomUserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	// ReadFile return ([]byte, error)
	privatekeyBytes, err := ioutil.ReadFile("keys/private.key")
	if err != nil {
		fmt.Println("error reading private key:", err)
		return "", err
	}
	// fmt.Println("private key bytes:", privatekeyBytes)
	privatekey, err := jwt.ParseRSAPrivateKeyFromPEM(privatekeyBytes)
	if err != nil {
		fmt.Println("error parse private key:", err)
	}
	// fmt.Println("Private key resilt parsing:", privatekey)
	// Create the Claims
	claims := &CustomUserClaims{
		username,
		jwt.StandardClaims{
			// Just expires in one minute :-)..
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "test",
		},
	}
	// Create new token, signed it, and then return
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedstring, err := token.SignedString(privatekey)
	if err != nil {
		fmt.Println("error signed sttring:", err)
		return "", err
	}
	return signedstring, nil

}

func verifyToken(tokenString string) (bool, error) {
	// jwt.Parse return *Token object with Raw, Method, Header, Claims, Signature and Valid field.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// read the key
		publickeybytes, err := ioutil.ReadFile("keys/public.key")
		if err != nil {
			return nil, fmt.Errorf("could not open jwt key, %v", err)
		}
		// jwt.ParseRSAPublicKeyFromPEM return *rsa.PrivateKey struct, represents an RSA key
		pubkey, err := jwt.ParseRSAPublicKeyFromPEM(publickeybytes)
		if err != nil {
			fmt.Println("error parse public key:", err)
		}
		return pubkey, nil
	})
	if err == nil && token.Valid {
		// valid thing. token.Raw, token.Method, token.Header dll
		return true, nil
	} else {
		return false, err
	}
}

/*
func main() {
	var usr = "paijo"
	fmt.Println("Generating token .........")
	token, err := GenerateToken(usr)
	if err != nil {
		log.Fatal("Error in generate token:", err)
	}
	fmt.Println("Generated token:", token)
	fmt.Println("Verifying token......")
	result, err := verifyToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Result verify:", result)

}
*/
