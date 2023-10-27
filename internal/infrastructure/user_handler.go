package handler

import (
	"InterviewPracticeBot-BE/internal/usecase/userusecase"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userUsecase *userusecase.UserUsecase
}

func NewauserHandler(userUsecase *userusecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call uh.userUsecase.Register with parsed data
	err = uh.userUsecase.Register(req.Email, req.Password)
	if err != nil {
		// Here you should handle different types of errors differently.
		// For example, if the error indicates that the email is already in use,
		// you should return a 409 Conflict status code.
		// But for simplicity, let's just return a 500 Internal Server Error for all errors.
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registration successful",
	})
}