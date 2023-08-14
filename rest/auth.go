package rest

import (
	"crypto/rand"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"movieraiting/entity"
	"net/http"
	"net/mail"
	"strings"
	"time"
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

	err = h.db.SignUpDB(r.Context(), user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

////////////////////////// логин вводить. почитать про куки http.Cookie  как посадить и как достать

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req SignInRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	login := req.Login
	password := req.Password

	user, err := h.db.GetUserByLogin(r.Context(), login)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		errorMessage := "invalid password"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(errorMessage))
		return
	}
	// ?? лучше для генерации
	//sessionID, err := generateSessionID(64)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    randomValue(64),               // sessionID   // сохранить session_id value куда-нибудь
		Expires:  time.Now().Add(3 * time.Hour), //прибавить 3 часа к текущему времени
		HttpOnly: true,
	})
}

// ?? лучше для генерации
//func generateSessionID(size int) (string, error) {
//	randomBytes := make([]byte, size)
//	_, err := rand.Read(randomBytes)
//	if err != nil {
//		log.Println(err)
//		return "", err
//	}
//	sessionID := base64.URLEncoding.EncodeToString(randomBytes)
//	return sessionID, nil
//}

func randomValue(size int) string { // SetCookie
	alphabet := "abcdefghigklmnopqrstuvwxyz"
	digits := "0123456789"
	symbols := strings.ToUpper(alphabet) + alphabet + digits
	val := ""
	for i := 0; i < size; i++ {

		r := rand.Intn(len(symbols) - 1)
		val += string(symbols[r])
	}
	return val
}

////////////////////////////

//change password

//
