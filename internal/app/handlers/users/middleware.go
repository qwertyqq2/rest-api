package users

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"test_go/pkg/auth"

	"github.com/julienschmidt/httprouter"
)

func authorizeMiddleware(next http.HandlerFunc) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				tokenClaims, err := auth.ParseToken(bearerToken[1])
				if err != nil {
					log.Fatal(err)
					return
				}
				log.Println(tokenClaims)
				cookieId := &http.Cookie{Name: "id", Value: fmt.Sprintf("%d", tokenClaims.UserId)}
				cookieStatus := &http.Cookie{Name: "status", Value: tokenClaims.Status}
				req.AddCookie(cookieId)
				req.AddCookie(cookieStatus)

				next(w, req)
				return
			}
		}
		log.Fatal(ErrTokenValid)
	})
}
