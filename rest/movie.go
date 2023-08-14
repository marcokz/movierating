// админ добавляет фильмы в базу
// create movie, get movie, update movie, delete movie
package rest

import (
	"encoding/json"
	"github.com/go-chi/render"
	"io"
	"log"
	"movieraiting/entity"
	"net/http"
)

type CreateMovieRequest struct {
	Title       string `json:"title"`
	Year        int64  `json:"year"`
	DirectorIDs int64  `json:"directorIDs"`
	ActorIDs    int64  `json:"actorIDs"`
	Description string `json:"description"`
}

func (h *Handler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var movie CreateMovieRequest
	err = json.Unmarshal(data, &movie)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	title := movie.Title
	if title == "" {
		errorMessage := "empty title"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	year := movie.Year
	if year < 1896 { // ?? первый фильм
		errorMessage := "empty year"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	directorID := movie.DirectorIDs
	if directorID < 1 {
		errorMessage := "empty director id"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	actorID := movie.ActorIDs
	if actorID < 1 {
		errorMessage := "empty actor id"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	description := movie.Description
	if description == "" {
		errorMessage := "empty description"
		log.Println(errorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorMessage))
		return
	}

	err = h.db.CreateMovieDB(r.Context(), movie)
	if err != nil {

	}

}

type MovieRequest struct {
	Title       string
	Year        int64
	Director    entity.Person
	Actor       entity.Person
	Description string
}

func (h *Handler) GetMovie(w http.ResponseWriter, r *http.Request) {
	results, err := h.db.GetMovie(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, results)
}

// по годам, по режиссерам, по актерам, по рейтингу
func (h *Handler) GetMovieList(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var list MovieRequest
	err = json.Unmarshal(data, &list)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	title := list.Title
	result, err := h.db.

	directorID := list.Director.ID
	results, err := h.db.GetMovieByDirecor(r.Context(), directorID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	render.JSON(results)
}
func (h *Handler) UpdateMovie(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {

}
