// models/task.go
package models

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
}
