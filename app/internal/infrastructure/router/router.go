package router

import (
	"net/http"
	"projects/LDmitryLD/repository/app/internal/modules"

	"github.com/go-chi/chi/v5"
)

func NewRouter(controllers *modules.Controllers) *chi.Mux {
	r := chi.NewRouter()
	setDefaultRoutes(r)

	r.Post("/api/users/create", controllers.User.Craete)
	r.Get("/api/users/{id}", controllers.User.GetByID)
	r.Post("/api/users/update", controllers.User.Update)
	r.Post("/api/users/delete", controllers.User.Delete)
	r.Post("/api/users/list", controllers.User.List)

	return r
}

func setDefaultRoutes(r *chi.Mux) {
	r.Get("/swagger", swaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	})
}
