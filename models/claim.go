package models

import jwt "github.com/dgrijaval/jwt-go"

type Claim struct {
	User `json: "user"`
	jwt.StandardClaim
}
