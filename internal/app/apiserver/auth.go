package apiserver

import (
	"github.com/golang-jwt/jwt/v4"
)

type tokenClaims struct {
	jwt.StandardClaims
	name string `json:"name"`
}

func (srv *APIServer) newToken(nameUser string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{},
		nameUser,
	})
	tokenStr, err := token.SignedString(srv.cfg.SecretKey)
	if err != nil {
		srv.logger.Error("Fail to generate token")
		return ""
	}
	srv.logger.Info("Generated new token")
	return tokenStr
}
func (srv *APIServer) validateToken(jwttoken string) bool {
	token, err := jwt.ParseWithClaims(
		jwttoken,
		jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return srv.cfg.SecretKey, nil
		},
	)
	if err != nil {
		return false
	}
	return token.Valid
}
