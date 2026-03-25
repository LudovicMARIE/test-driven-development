package video_store

import "strings"

type Movie struct {
	Title string
	Year  int
}

func FindByTitle(title string, collection []Movie) []Movie {
	var foundMovies []Movie
	for _, movie := range collection {
		if strings.Contains(strings.ToLower(movie.Title), strings.ToLower(title)) {
			foundMovies = append(foundMovies, movie)
		}
	}

	return foundMovies
}
