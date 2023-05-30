package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
	"user-management/store"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var router = gin.Default()

func basicAuth(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == "testuser" && password == "testpass" {
		fmt.Println("user authenticated")
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError , gin.H{"message": "user not authenticated"})
		return 
	}
}

func main() {
	router.Use(cors.Middleware(cors.Config{
	Origins:        "*",
	Methods:        "GET, PATCH, POST, DELETE",
	RequestHeaders: "Origin, Authorization, Content-Type",
	ExposedHeaders: "",
	MaxAge: 50 * time.Second,
	Credentials: false,
	ValidateHeaders: false,
}))
	router.GET("/users/:id", basicAuth, getUser)
	router.POST("/users", basicAuth, postUser)
	router.PATCH("/users/:id", basicAuth,updateUser)
	router.DELETE("/users/:id", basicAuth,deleteUser)
	router.GET("/users", basicAuth, getAllUsers)
	router.POST("/users/contracts/:id", basicAuth,userContractUpload)
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
	c.IndentedJSON(http.StatusOK, person)
}

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, store.PersonStorage)
}

func postUser(c *gin.Context){
	var newUserID = strconv.Itoa(len(store.PersonStorage)+1)
	var newPerson store.Person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	store.PersonStorage[newUserID] = newPerson
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
	store.PersonStorage[id] = updatedPerson
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


