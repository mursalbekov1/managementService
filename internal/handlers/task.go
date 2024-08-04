package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"task3/internal/models"
)

func (h Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var task models.Task
	json.Unmarshal(body, &task)

	// Append to the Tasks table
	if result := h.DB.Create(&task); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the project by Id

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that task
	h.DB.Delete(&task)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func (h Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	if result := h.DB.Find(&tasks); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateTask models.Task
	json.Unmarshal(body, &updateTask)

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	task.Name = updateTask.Name
	task.Description = updateTask.Description
	task.Project = updateTask.Project

	h.DB.Save(&task)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h Handler) GetTasksByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	var tasks []models.Task
	if result := h.DB.Where("name LIKE ?", "%"+title+"%").Find(&tasks); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// GetTasksByStatus обрабатывает запросы GET /tasks/search?status={status}
func (h Handler) GetTasksByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	var tasks []models.Task
	if result := h.DB.Where("state = ?", status).Find(&tasks); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h Handler) GetTasksByPriority(w http.ResponseWriter, r *http.Request) {
	priority := r.URL.Query().Get("priority")

	var tasks []models.Task
	if result := h.DB.Where("priority = ?", priority).Find(&tasks); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h Handler) GetTasksByAssignee(w http.ResponseWriter, r *http.Request) {
	assignee := r.URL.Query().Get("assignee")

	var tasks []models.Task
	if result := h.DB.Where("responsible = ?", assignee).Find(&tasks); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h Handler) GetTasksByProject(w http.ResponseWriter, r *http.Request) {
	project := r.URL.Query().Get("project")

	var tasks []models.Task
	if result := h.DB.Where("project = ?", project).Find(&tasks); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
