package service

import (
	"context"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/markbates/goth"
)

type hasuraClaims struct {
	AccessToken    string      `json:"accesstoken"`
	CustomUserInfo interface{} `json:"https://hasura.io/jwt/claims"`
	jwt.StandardClaims
}

func (s *Service) createJWT(uid string) (string, error) {
	hasura := make(map[string]interface{})
	hasura["x-hasura-default-role"] = "user"
	hasura["x-hasura-allowed-roles"] = []string{"user"}
	hasura["x-hasura-user-id"] = "1"
	claims := hasuraClaims{
		uid,
		hasura,
		jwt.StandardClaims{
			//ExpiresAt: 150000000,
			Issuer: "test",
		},
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, err := t.SignedString(s.signKey)
	if err != nil {
		log.Printf("Token Signing error: %v\n", err)
		return "", err
	}
	return tokenString, nil

}

//CompleteAuth func
func (s *Service) CompleteAuth(ctx context.Context, guser *goth.User) (string, error) {
	uid, err := s.checkUser(ctx, guser)
	if err != nil {
		log.Printf(`cannot check user in db : %v`, err)
	}
	log.Printf(`the new user is %v`, uid)
	return s.createJWT(uid)

}
