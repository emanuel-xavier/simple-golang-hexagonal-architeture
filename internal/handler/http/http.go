package httpHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/emanuel-xavier/hexagonal-architerure/configs"
	"github.com/emanuel-xavier/hexagonal-architerure/internal/core/port"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpHandler struct {
	uuc port.UserUseCase
}

func NewHandler(uUsecase port.UserUseCase) *HttpHandler {
	handler := HttpHandler{
		uuc: uUsecase,
	}

	return &handler
}

func (h *HttpHandler) Handle() {
	r := chi.NewRouter()

	r.Use(middleware.DefaultLogger)

	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", h.getUserById)
		r.Get("/list", h.getAllUsers)
		r.Post("/", h.createUser)
	})

	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r))
}

func (h *HttpHandler) getUserById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	user, err := h.uuc.GetUserById(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *HttpHandler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.uuc.GetAll()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	log.Println(users)
	json.NewEncoder(w).Encode(users)
}

func (h *HttpHandler) createUser(w http.ResponseWriter, r *http.Request) {
	type userRequest struct {
		Username string `json:"username"`
		Id       string `json:"id"`
	}
	var u userRequest

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Println(err)
		http.Error(w, "Can't decode user", http.StatusUnprocessableEntity)
		return
	}

	user, err := h.uuc.CreateUser(u.Username, u.Id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Can't inset this user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
