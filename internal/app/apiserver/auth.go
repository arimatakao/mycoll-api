package apiserver

import (
	"github.com/golang-jwt/jwt/v4"
)

type tokenClaims struct {
	jwt.StandardClaims
	Id string `json:"_id" bson:"_id"`
}

func (srv APIServer) newToken(Id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{},
		Id,
	})
	tokenStr, err := token.SignedString([]byte(srv.cfg.SecretKey))
	if err != nil {
		srv.logger.Error("Fail to generate token")
		return ""
	}
	srv.logger.Info("Generated new token")
	return tokenStr
}

func (srv APIServer) isTokenValid(jwttoken string) bool {
	token, err := jwt.ParseWithClaims(
		jwttoken,
		jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(srv.cfg.SecretKey), nil
		},
	)
	if err != nil {
		return false
	}
	return token.Valid
}

func (srv APIServer) getClaimsFromJWT(jwttoken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(jwttoken,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(srv.cfg.SecretKey), nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
