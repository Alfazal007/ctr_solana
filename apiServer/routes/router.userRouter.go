package router

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/controllers"
	"github.com/go-chi/chi"
)

func UserRouter(apiCfg *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create", apiCfg.CreateUser)
	r.Post("/login", apiCfg.LoginUser)

	r.Get("/current-user", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.CurrentUser)).ServeHTTP)
	r.Post("/update-balance", apiCfg.IncreaseBalance)
	r.Post("/verify", apiCfg.AddCreatorPK)
	r.Get("/balance", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.FetchBalance)).ServeHTTP)
	r.Post("/logout", apiCfg.Logout)
	return r
}
