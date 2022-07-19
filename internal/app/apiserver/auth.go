package apiserver

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func (srv *APIServer) newToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{ID: id})
	tokenStr, err := token.SignedString([]byte(srv.cfg.SecretKey))
	if err != nil {
		srv.logger.Error("Fail to generate token")
		return ""
	}
	srv.logger.Info("Generated new token")
	return tokenStr
}

func (srv *APIServer) parseTokenFromHeader(r *http.Request) (string, bool) {
	header, ok := r.Header["Authorization"]
	if !ok {
		return "", false
	}
	tokenStr := strings.Split(header[0], " ")
	if len(tokenStr) != 2 || tokenStr[0] != "Bearer" {
		return "", false
	}

	t, err := jwt.ParseWithClaims(tokenStr[1], &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(srv.cfg.SecretKey), nil
	})
	if !t.Valid || err != nil {
		return "", false
	}
	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", false
	}
	return claims.ID, true
}
