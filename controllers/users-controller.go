package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/myapp/models"
)

// UserController represents the controller for the users resource.
type UserController struct {
    // UserService is the service that handles business logic for users.
    UserService *models.UserService
}

// NewUserController creates a new UserController with the given UserService.
func NewUserController(userService *models.UserService) *UserController {
    return &UserController{
        UserService: userService,
    }
}

// RegisterRoutes registers the HTTP routes for the users resource.
func (c *UserController) RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/users", c.createUser).Methods(http.MethodPost)
    r.HandleFunc("/users/{id}", c.getUser).Methods(http.MethodGet)
}

// createUser handles the creation of a new user.
func (c *UserController) createUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := c.UserService.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(&user)
}

// getUser handles the retrieval of a user by ID.
func (c *UserController) getUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    user, err := c.UserService.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if user == nil {
        http.NotFound(w, r)
        return
    }

    json.NewEncoder(w).Encode(user)
}
