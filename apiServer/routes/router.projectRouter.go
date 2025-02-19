package router

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/controllers"
	"github.com/go-chi/chi"
)

func ProjectRouter(apiCfg *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.GetProjectInfo)).ServeHTTP)
	r.Get("/labeller/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.GetProjectToVote)).ServeHTTP)
	r.Post("/create-project", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.CreateNewCTR)).ServeHTTP)
	r.Post("/add-image/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.GetUrlToUploadImage)).ServeHTTP)
	r.Put("/start/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.StartVote)).ServeHTTP)
	r.Put("/end/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.EndVote)).ServeHTTP)
	r.Post("/vote/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.CreateVote)).ServeHTTP)
	r.Get("/projects", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.GetCreatorProjects)).ServeHTTP)
	r.Get("/result/{projectId}", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.VotedProjects)).ServeHTTP)
	r.Get("/projectsToVote", controllers.VerifyJWT(apiCfg, http.HandlerFunc(apiCfg.ProjectsToVote)).ServeHTTP)
	return r
}
