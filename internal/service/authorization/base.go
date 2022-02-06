package authorization

import "github.com/golang-jwt/jwt"

type AuthorizationClaims struct {
	ID        int    `json:"id"`
	Type      int    `json:"type"`
	ProfileID int    `json:"profileID"`
	Name      string `json:"profile.name"`
	Desc      string `json:"profile.desc"`
	jwt.StandardClaims
}
