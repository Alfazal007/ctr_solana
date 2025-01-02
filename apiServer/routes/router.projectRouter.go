package router

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/controllers"
	"github.com/go-chi/chi"
)

func ProjectRouter(apiCfg *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create-project", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.CreateNewCTR)).ServeHTTP)
	return r
}
