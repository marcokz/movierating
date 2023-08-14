package rest

import "net/http"

// поставить пользовательскую оценку
func (h *Handler) SetUserRating(w http.ResponseWriter, r *http.Request) {
	//SQL: insert on conflict update(если оценка существует то апдейт делает(перезапись))
}

func (h *Handler) DeleteUserRating(w http.ResponseWriter, r *http.Request) {

}

// поиск оценок по movie id и rating ratio(совпадение оценок)
func (h *Handler) GetUserRatings(w http.ResponseWriter, r *http.Request) {

}
