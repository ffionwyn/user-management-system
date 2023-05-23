package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"user-management/store"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()


func main() {
	router.GET("/users/:id", getUser)
	router.POST("/users", postUser)
	router.PATCH("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	router.GET("/users", getAllUsers)
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


func fileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	uploadFile(w, r)
	default:
	w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fileName := fmt.Sprintf("user-contracts/user-contract-id-%s.pdf", id)

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("userContract")
	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("MIME header: %+v\n", handler.Header)

	dst, err := os.Create(fileName)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully uploaded file\n")
}
