package main

import (
	"io/ioutil"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/twitter"
)

func boot() {
	//load keys from files in .env
	signBytes, _ := ioutil.ReadFile(os.Getenv("RSASECRET"))
	signKey, _ = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	verifyBytes, _ := ioutil.ReadFile(os.Getenv("RSAPUBLIC"))
	verifyKey, _ = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	//init goth authentication providers
	//gothic.Store = sessions.NewCookieStore([]byte("secret"))
	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
	)

	//init db

}
