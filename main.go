package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
	"user-management/store"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var router = gin.Default()


func main() {
	router.Use(cors.Middleware(cors.Config{
	Origins:        "*",
	Methods:        "GET, PUT, POST, DELETE",
	RequestHeaders: "Origin, Authorization, Content-Type",
	ExposedHeaders: "",
	MaxAge: 50 * time.Second,
	Credentials: false,
	ValidateHeaders: false,
}))
	router.GET("/users/:id", getUser)
	router.POST("/users", postUser)
	router.PATCH("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	router.GET("/users", getAllUsers)
	router.POST("/users/contracts/:id", userContractUpload)
	router.Run(":5000")
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	person, err := store.GetPersonByID(id)
	if err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	response := fmt.Sprintf("firstName: %s, secondName: %s, dob: %s",
		person.FirstName, person.SecondName, person.DOB)
		c.IndentedJSON(http.StatusOK, response)
}

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, store.PersonStorage)
}

func postUser(c *gin.Context){
	var newPerson store.Person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	store.PersonStorage[newPerson.UserID] = newPerson
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	_, ok := store.PersonStorage[id]
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	var updatedPerson store.Person
	if err := c.BindJSON(&updatedPerson); err != nil {
		return
	}
	store.PersonStorage[updatedPerson.UserID] = updatedPerson
	c.IndentedJSON(http.StatusCreated, updatedPerson)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var err = store.DeletePerson(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{})
}

func userContractUpload(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		fileName := filepath.Base(fmt.Sprintf("user-contract-id-%s.pdf", id))
		if err := c.SaveUploadedFile(file, fileName); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
		c.String(http.StatusOK, "contract uploaded successfully")
	}

