package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/temurova-ui/FinanceBro/internal/model"
)

type UserService interface {
	Register(ctx context.Context, req model.Register) (int, error)
}

type UserHandler struct{
	ser UserService
}

func NewUserHandler(ser UserService) *UserHandler{
	return &UserHandler{
		ser: ser,
	}
}

func (s *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.Register
	json.NewDecoder(r.Body).Decode(&user)
	id, err := s.ser.Register(r.Context(),user)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.User{
		ID: id,
	})
	
}