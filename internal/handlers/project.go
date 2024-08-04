package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"task3/internal/models"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func (h Handler) AddProject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var project models.Project
	json.Unmarshal(body, &project)

	// Append to the Books table
	if result := h.DB.Create(&project); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the project by Id

	var project models.Project

	if result := h.DB.First(&project, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	h.DB.Delete(&project)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func (h Handler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project

	if result := h.DB.Find(&projects); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}

func (h Handler) GetProject(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var project models.Project

	if result := h.DB.First(&project, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}

func (h Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedProject models.Project
	json.Unmarshal(body, &updatedProject)

	var project models.Project

	if result := h.DB.First(&project, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	project.Name = updatedProject.Name
	project.Description = updatedProject.Description
	project.Manager = updatedProject.Manager

	h.DB.Save(&project)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h Handler) GetProjectsByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	var projects []models.Project

	if result := h.DB.Where("name ILIKE ?", "%"+title+"%").Find(&projects); result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, "Failed to retrieve projects", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}

func (h Handler) GetProjectsByManager(w http.ResponseWriter, r *http.Request) {
	manager := r.URL.Query().Get("manager")
	managerID, err := strconv.Atoi(manager)
	if err != nil {
		http.Error(w, "Invalid manager ID", http.StatusBadRequest)
		return
	}

	var projects []models.Project
	if result := h.DB.Where("manager = ?", managerID).Find(&projects); result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, "Failed to retrieve projects", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
