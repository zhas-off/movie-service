package data

import (
	"time"

	"github.com/zhas-off/movie-service/internal/validator"
)

// Movie type whose fields describe the movie.
// Note that the Runtime type uses a custom Runtime type instead of int32. Furthermore, the omitempy
// directive on the Runtime type will still work on this: if the Runtime field has the underlying
// value 0, then it will be considered empty and omitted -- and the MarshalJSON() method won't
// be called.
type Movie struct {
	ID        int64     `json:"id"` // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`  // Use the - directive to never export in JSON output
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // Movie release year0
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"` // The version number starts at 1 and is incremented each
	// time the movie information is updated.
}

// ValidateMovie runs validation checks on Movie type.
func ValidateMovie(v *validator.Validator, movie *Movie) {
	// Check movie.Title
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	// Check input.Year
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	// Check input.Runtime
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	// Check input.Genres
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")

}
