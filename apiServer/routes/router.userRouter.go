package router

import (
	"github.com/Alfazal007/ctr_solana/controllers"
	"github.com/go-chi/chi"
)

func UserRouter(apiCfg *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create", apiCfg.CreateUser)
	return r
}
