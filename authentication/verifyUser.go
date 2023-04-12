package authentication

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	secret = []byte("signature_hmac_secret_shared_key")
)

type fooClaims struct {
	Email string `json:"email"`
}

func GenerateToken(ctx iris.Context) {
	signer := jwt.NewSigner(jwt.HS256, secret, 100*time.Minute)
	var claims fooClaims
	ctx.ReadJSON(&claims)

	token, err := signer.Sign(claims)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	ctx.SetCookieKV("token", string(token))
	ctx.Write(token)
}

func VerifyMiddleware(ctx iris.Context) {
	token := []byte(ctx.GetCookie("token"))
	_, err := jwt.Verify(jwt.HS256, secret, token)
	if err != nil {
		ctx.JSON("authentication failed")
		return
	}
	ctx.Next()
}
func Logout(ctx iris.Context) {
	// err := ctx.Logout()
	// if err != nil {
	// 	ctx.WriteString(err.Error())
	// } else {
	// 	ctx.Writef("token invalidated, a new token is required to access the protected API")
	// }
	cookie := iris.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	ctx.SetCookie(&cookie)
	ctx.JSON("successfullt logout")
}
