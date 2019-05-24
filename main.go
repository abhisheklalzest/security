package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

func init() {
	// fmt.Println("INIT Security")
}

// CustomPayload struct
type CustomPayload struct {
	jwt.Payload

	IsLoggedIn bool   `json:"isLoggedIn"`
	Scope      string `json:"scope"`
	ClientID   string `json:"client_id"`
}

// ErrInvalidToken invalid token
var ErrInvalidToken = errors.New("jwt: token is invalid")

// ScopeValidator validate Scope
func ScopeValidator(p CustomPayload, scope string) error {
	if p.Scope == scope {
		return nil
	}
	return errors.New("jwt: scope is invalid")
}

// Validate jwt token
func Validate(tok string) error {

	now := time.Now()
	// hs256 := jwt.NewHMAC(jwt.SHA256, []byte("secret"))
	token := []byte(tok)

	raw, err := jwt.Parse(token)
	if err != nil {
		// fmt.Println("ERROR Parse", err)
		return ErrInvalidToken
	}

	// if err = raw.Verify(hs256); err != nil {
	// 	// Handle error.
	fmt.Println("ERROR Verify", err)
	// 	// return ErrInvalidToken
	// }
	var (
		h jwt.Header
		p CustomPayload
	)
	_ = h
	if h, err = raw.Decode(&p); err != nil {
		// Handle error.
		return ErrInvalidToken
	}
	var iss = "http://staging.zestmoney.in"
	var scope = "internal_services"
	var audience = "http://staging.zestmoney.in/resources"
	issValidator := jwt.IssuerValidator(iss)
	audValidator := jwt.AudienceValidator(jwt.Audience{audience})
	expValidator := jwt.ExpirationTimeValidator(now, true)
	nbfValidator := jwt.NotBeforeValidator(now)
	iatValidator := jwt.IssuedAtValidator(now)
	isScopeVaid := ScopeValidator(p, scope)
	if err := p.Validate(issValidator, iatValidator, expValidator, nbfValidator, audValidator); err != nil {
		// fmt.Println("ERROR Validate", err)
		return ErrInvalidToken
	} else if isScopeVaid != nil {
		// fmt.Println("in valid scope")
		return ErrInvalidToken
	} else {
		// fmt.Println("Valid")
		return nil
	}
}

func main() {
	fmt.Println("Hello")
}