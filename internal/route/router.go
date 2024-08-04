package route

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"task3/internal/handlers"
)

func Router(h handlers.Handler) http.Handler {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthCheck", healthCheckHandler)

		// Users routes
		r.Route("/user", func(r chi.Router) {
			r.Get("/users", h.GetAllUsers)                // GET /users
			r.Post("/users", h.AddUser)                   // POST /users
			r.Get("/users/{id}", h.GetUser)               // GET /users/{id}
			r.Put("/users/{id}", h.UpdateUser)            // PUT /users/{id}
			r.Delete("/users/{id}", h.DeleteUser)         // DELETE /users/{id}
			r.Get("/users/{id}/tasks", h.GetUser)         // GET /users/{id}/tasks
			r.Get("/users/search?name={name}", h.GetUser) // GET /users/search?name={name} or GET /users/search?email={email}
		})

		// Tasks routes
		r.Route("/task", func(r chi.Router) {
			r.Get("/tasks", h.GetAllTasks)                  // GET /tasks
			r.Post("/tasks", h.AddTask)                     // POST /tasks
			r.Get("/tasks/{id}", h.GetTask)                 // GET /tasks/{id}
			r.Put("/tasks/{id}", h.UpdateTask)              // PUT /tasks/{id}
			r.Delete("/tasks/{id}", h.DeleteTask)           // DELETE /tasks/{id}
			r.Get("/tasks/search?title={title}", h.GetTask) // GET /tasks/search?title={title}, status={status}, priority={priority}, assignee={userId}, project={projectId}
		})

		// Projects routes
		r.Route("/project", func(r chi.Router) {
			r.Get("/projects", h.GetAllProjects)                  // GET /projects
			r.Post("/projects", h.AddProject)                     // POST /projects
			r.Get("/projects/{id}", h.GetProject)                 // GET /projects/{id}
			r.Put("/projects/{id}", h.UpdateProject)              // PUT /projects/{id}
			r.Delete("/projects/{id}", h.DeleteProject)           // DELETE /projects/{id}
			r.Get("/projects/{id}/tasks", h.GetProject)           // GET /projects/{id}/tasks
			r.Get("/projects/search?title={title}", h.GetProject) // GET /projects/search?title={title}, manager={userId}
		})
	})

	return router
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
