package authorization

import (
	"github.com/golang-jwt/jwt"
	"github.com/meloalright/guora/conf"
	"time"

	"github.com/meloalright/guora/internal/model"
	"net/url"
	"strings"
)

func Gen(user model.User) (tokenString string, err error) {

	SecretString := conf.Config().Secretkey

	Secret := []byte(SecretString)

	claims := AuthorizationClaims{
		user.ID,
		user.Type,
		user.ProfileID,
		compatibleJSEncodeURIComponent(url.QueryEscape(user.Profile.Name)),
		compatibleJSEncodeURIComponent(url.QueryEscape(user.Profile.Desc)),
		jwt.StandardClaims{
			ExpiresAt: time.Time.AddDate(time.Now(), 0, 0, 7).Unix(),
			Issuer:    "localhost",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(Secret)

	return
}

func compatibleJSEncodeURIComponent(str string) (resultStr string) {
	resultStr = str
	resultStr = strings.Replace(resultStr, "+", "%20", -1)
	resultStr = strings.Replace(resultStr, "%21", "!", -1)
	resultStr = strings.Replace(resultStr, "%27", "'", -1)
	resultStr = strings.Replace(resultStr, "%28", "(", -1)
	resultStr = strings.Replace(resultStr, "%29", ")", -1)
	resultStr = strings.Replace(resultStr, "%2A", "*", -1)
	return
}
