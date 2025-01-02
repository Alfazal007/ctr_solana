package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func VerifyJWT(apiCfg *ApiConf, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accessToken")
		var jwtToken string
		if err != nil {
			helpers.RespondWithError(w, 400, "Error reading cookie, try logging in again")
			return
		} else {
			jwtToken = cookie.Value
		}
		// Verify the JWT token
		jwtSecret := utils.LoadEnvVariables().AccessTokenSecret
		if jwtToken == "" {
			helpers.RespondWithError(w, 400, "Provide cookie")
			return
		}
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		if err != nil {
			helpers.RespondWithError(w, 401, fmt.Sprintf("Invalid token here %v", err))
			return
		}
		if !token.Valid {
			helpers.RespondWithError(w, 401, fmt.Sprintf("Invalid token %v", err))
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helpers.RespondWithError(w, 400, "Invalid claims login again")
			return
		}

		username := claims["username"].(string)
		id := claims["user_id"].(string)

		user, err := apiCfg.DB.GetUserByUsername(r.Context(), username)
		if err != nil {
			helpers.RespondWithError(w, 400, "Some manpulation done with the token")
			return
		}
		idUUID, err := uuid.Parse(id)
		if err != nil {
			helpers.RespondWithError(w, 400, "Some manpulation done with the token")
			return
		}
		if idUUID != user.ID {
			helpers.RespondWithError(w, 400, "Some manipulations done with the token try again")
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
