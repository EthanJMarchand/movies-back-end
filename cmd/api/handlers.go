package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "0.1.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	// Fetch all movies
	var movies []models.Movie

	rd, err := time.Parse("2006-01-02", "1986-03-07")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: rd,
		MPAARating:  "R",
		RunTime:     116,
		Description: "This is a movie description for highlander",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	rd, err = time.Parse("2006-01-02", "1986-03-07")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
	rotla := models.Movie{
		ID:          2,
		Title:       "Raiders of the Lost Arc",
		ReleaseDate: rd,
		MPAARating:  "R",
		RunTime:     120,
		Description: "This is a movie description for raiders of the lost arc",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	movies = append(movies, rotla)

	// write the movies to the screen
	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
