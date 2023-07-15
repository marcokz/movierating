package rest

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"movieraiting/entity"
	"net/http"
	"net/mail"
	"unicode/utf8"
)

type SignUpRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req SignUpRequest
	err = json.Unmarshal(data, &req) //десерелизация
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//принцип нулевого доверия к входным данным с фронта, потому что неизвестно что придет с фронта

	login := req.Login
	if login == "" {
		errorMessage := "empty login"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}
	addr, err := mail.ParseAddress(req.Email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	email := addr.Address

	pass := req.Password
	if utf8.RuneCountInString(pass) < 8 {
		errorMessage := "password requires at least 8 characters"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	user := entity.User{
		Login:        login,
		Email:        email,
		PasswordHash: passHash,
	}

	err = h.db.CreateUser(r.Context(), user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
