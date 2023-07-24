package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// AutoMigrate the model to create the table
	db.AutoMigrate(&Task{})
}

func main() {
	initDB()
	r := gin.Default()

	// Define the CRUD endpoints
	r.POST("/tasks", createTask)
	r.GET("/tasks/:id", getTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)
	r.GET("/tasks", listTasks)

	// Run the server on port 8080
	r.Run(":8080")
}

func createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	db.Create(&task)
	c.JSON(201, task)
}

func getTask(c *gin.Context) {
	var task Task
	if err := db.First(&task, c.Param("id")).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}

func updateTask(c *gin.Context) {
	var task Task
	if err := db.First(&task, c.Param("id")).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Task not found"})
		return
	}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	db.Save(&task)
	c.JSON(200, task)
}

func deleteTask(c *gin.Context) {
	var task Task
	if err := db.First(&task, c.Param("id")).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Task not found"})
		return
	}
	db.Delete(&task)
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}

func listTasks(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(200, tasks)
}
