README
This repository contains a simple user management system implemented in Go using the Gin framework. 
It provides a RESTful API for creating, retrieving, updating, and deleting user information.


Getting Started
The application will run on http://localhost:5000.

| Endpoint               | Method | Description                               |
|------------------------|--------|-------------------------------------------|
| `/users/:id`           | GET    | Retrieve a specific user by their ID.      |
| `/users`               | GET    | Retrieve all users.                        |
| `/users`               | POST   | Create a new user.                         |
| `/users/:id`           | PATCH  | Update an existing user.                   |
| `/users/:id`           | DELETE | Delete a user.                             |
| `/users/contracts/:id` | POST   | Upload a user contract file.                |

Authentication
The user management system uses basic authentication to secure the API endpoints. The default credentials for authentication are:

Username: testuser
Password: testpass

You can modify the authentication logic in the basicAuth function in main.go to use your own authentication mechanism.

Data Storage
The user data is stored in memory using a PersonStorage map in the store package. Each user is represented by a Person Struct, which contains the following fields:

FirstName: The first name of the user.
SecondName: The second name of the user.
DOB: The date of birth of the user.
Email: The email address of the user.

The store.go file in the store package provides functions to manipulate the user data, including creating, retrieving, updating, and deleting users.

File Upload
The user management system supports file upload for user contracts. The endpoint POST /users/contracts/:id allows you to upload a contract file for a specific user. The uploaded files are saved on the server using the format "user-contract-id-{id}.pdf".

CORS Configuration
The user management system allows cross-origin resource sharing (CORS) from all origins (*). The allowed HTTP methods are GET, PATCH, POST, and DELETE. The allowed request headers include Origin, Authorization, and Content-Type. The CORS configuration can be found in the main function in main.go.

Dependencies
The user management system relies on the following third-party dependencies:

github.com/gin-gonic/gin: A web framework for Go.
github.com/itsjamie/gin-cors: CORS middleware for Gin.
The Go modules system is used to manage the dependencies.

