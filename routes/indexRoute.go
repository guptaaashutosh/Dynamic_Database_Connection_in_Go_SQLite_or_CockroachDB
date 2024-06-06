package routes

import (
	"database/sql"
	"net/http"
	"tutorials/dynamic_db_conn_project/controllers"
	"tutorials/dynamic_db_conn_project/middlewares"

	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(db *sql.DB) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middlewares.SetDBConnection(db))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dynamic DB Connection Project"))
	})

	r.Post("/register", controllers.RegisterUser)
	r.Get("/users", controllers.GetUsers)
	r.Put("/update/{id}", controllers.UpdateUser)
	r.Delete("/{id}", controllers.DeleteUser)

	return r
}
