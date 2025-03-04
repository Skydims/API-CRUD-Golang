// handlers/task_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"todo-app/database"
	"todo-app/models"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	if err := database.DB.First(&task, params["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	if err := database.DB.First(&task, params["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	if err := database.DB.First(&task, params["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}