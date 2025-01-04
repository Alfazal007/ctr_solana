package router

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/controllers"
	"github.com/go-chi/chi"
)

func ProjectRouter(apiCfg *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create-project", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.CreateNewCTR)).ServeHTTP)
	r.Post("/add-image/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.GetUrlToUploadImage)).ServeHTTP)
	r.Put("/start/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.StartVote)).ServeHTTP)
	r.Put("/end/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.EndVote)).ServeHTTP)
	r.Post("/vote/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.CreateVote)).ServeHTTP)
	return r
}
