package yandex

import (
	"fmt"
	"os"
)

type OAuthRequestBody struct {
	clientId     string
	clientSecret string
	grantType    string
	username     string
	password     string
}

func NewOAuthRequestBody() OAuthRequestBody {
	return OAuthRequestBody{
		clientId:     "23cabbbdc6cd418abb4b39c32c41195d", //yandex-music
		clientSecret: "53bc75238f0c4d08a118e51fe9203300",
		grantType:    "password",
		username:     os.Getenv("LOGIN"),
		password:     os.Getenv("PASSWORD"),
	}
}

func (b OAuthRequestBody) String() string {
	return fmt.Sprintf(
		"grant_type=%s&client_id=%s&client_secret=%s&username=%s&password=%s",
		b.grantType,
		b.clientId,
		b.clientSecret,
		b.username,
		b.password,
	)
}
