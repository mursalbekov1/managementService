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
			r.Get("/users", h.GetAllUsers)             // GET /users
			r.Post("/users", h.AddUser)                // POST /users
			r.Get("/users/{id}", h.GetUser)            // GET /users/{id}
			r.Put("/users/{id}", h.UpdateUser)         // PUT /users/{id}
			r.Delete("/users/{id}", h.DeleteUser)      // DELETE /users/{id}
			r.Get("/users/{id}/tasks", h.GetUserTasks) // GET /users/{id}/tasks
			r.Get("/users/search", func(w http.ResponseWriter, r *http.Request) {
				name := r.URL.Query().Get("name")
				email := r.URL.Query().Get("email")

				if name != "" {
					h.GetUsersByName(w, r)
					return
				}
				if email != "" {
					h.GetUsersByEmail(w, r)
					return
				}

				http.Error(w, "Query parameter 'name' or 'email' is required", http.StatusBadRequest)
			})
		})

		// Tasks routes
		r.Route("/task", func(r chi.Router) {
			r.Get("/tasks", h.GetAllTasks)        // GET /tasks
			r.Post("/tasks", h.AddTask)           // POST /tasks
			r.Get("/tasks/{id}", h.GetTask)       // GET /tasks/{id}
			r.Put("/tasks/{id}", h.UpdateTask)    // PUT /tasks/{id}
			r.Delete("/tasks/{id}", h.DeleteTask) // DELETE /tasks/{id}

			r.Get("/tasks/search", func(w http.ResponseWriter, r *http.Request) {
				title := r.URL.Query().Get("title")
				status := r.URL.Query().Get("status")
				priority := r.URL.Query().Get("priority")
				assignee := r.URL.Query().Get("assignee")
				project := r.URL.Query().Get("project")

				if title != "" {
					h.GetTasksByTitle(w, r)
					return
				}
				if status != "" {
					h.GetTasksByStatus(w, r)
					return
				}
				if priority != "" {
					h.GetTasksByPriority(w, r)
					return
				}
				if assignee != "" {
					h.GetTasksByAssignee(w, r)
					return
				}
				if project != "" {
					h.GetTasksByProject(w, r)
					return
				}

				http.Error(w, "Query parameter 'title', 'status', 'priority', 'assignee', or 'project' is required", http.StatusBadRequest)
			})
		})

		// Projects routes
		r.Route("/project", func(r chi.Router) {
			r.Get("/projects", h.GetAllProjects)        // GET /projects
			r.Post("/projects", h.AddProject)           // POST /projects
			r.Get("/projects/{id}", h.GetProject)       // GET /projects/{id}
			r.Put("/projects/{id}", h.UpdateProject)    // PUT /projects/{id}
			r.Delete("/projects/{id}", h.DeleteProject) // DELETE /projects/{id}
			r.Get("/projects/{id}/tasks", h.GetProject) // GET /projects/{id}/tasks
			r.Get("/projects/search", func(w http.ResponseWriter, r *http.Request) {
				title := r.URL.Query().Get("title")
				manager := r.URL.Query().Get("manager")

				if title != "" {
					h.GetProjectsByTitle(w, r)
					return
				}
				if manager != "" {
					h.GetProjectsByManager(w, r)
					return
				}

				http.Error(w, "Query parameter 'title' or 'manager' is required", http.StatusBadRequest)
			})
		})
	})

	return router
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
