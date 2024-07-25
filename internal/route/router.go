package route

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthCheck", healthCheckHandler)
	})

	return router
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
