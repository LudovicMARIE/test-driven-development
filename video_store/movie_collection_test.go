package video_store

import (
	"testing"
)

func TestFindByTitle(t *testing.T) {
	testCase := []struct {
		name       string
		title      string
		collection []Movie
		result     []Movie
	}{
		{
			name:  "Cas nominal - The Matrix",
			title: "The Matrix",
			collection: []Movie{
				{Title: "The Matrix", Year: 1999},
				{Title: "Avatar", Year: 2009},
				{Title: "Avatar2", Year: 2014},
				{Title: "Forest Gump", Year: 2008},
				{Title: "AAAAAAAAAAAAAAAA, le film", Year: 2036},
			},
			result: []Movie{
				{Title: "The Matrix", Year: 1999},
			},
		},
		{
			name:  "Recherche avec la lettre 't'",
			title: "t",
			collection: []Movie{
				{Title: "The Matrix", Year: 1999},
				{Title: "Intouchable", Year: 2011},
				{Title: "Forest Gump", Year: 1994},
				{Title: "A beautiful mind", Year: 2001},
			},
			result: []Movie{
				{Title: "The Matrix", Year: 1999},
				{Title: "Intouchable", Year: 2011},
				{Title: "Forest Gump", Year: 1994},
				{Title: "A beautiful mind", Year: 2001},
			},
		},
		{
			name:  "Recherche vide",
			title: "",
			collection: []Movie{
				{Title: "The Matrix", Year: 1999},
				{Title: "Intouchable", Year: 2011},
				{Title: "Forest Gump", Year: 1994},
				{Title: "A beautiful mind", Year: 2001},
			},
			result: []Movie{
				{Title: "The Matrix", Year: 1999},
				{Title: "Intouchable", Year: 2011},
				{Title: "Forest Gump", Year: 1994},
				{Title: "A beautiful mind", Year: 2001},
			},
		},
		{
			name:       "Recherche sur collection vide",
			title:      "",
			collection: []Movie{},
			result:     []Movie{},
		},
		{
			name:  "Recherche qui ne donne pas de résultat",
			title: "lkdsqljdnsqkbcezicndjncolkjd",
			collection: []Movie{
				{Title: "The Matrix", Year: 1999},
				{Title: "Intouchable", Year: 2011},
				{Title: "Forest Gump", Year: 1994},
				{Title: "A beautiful mind", Year: 2001},
			},
			result: []Movie{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := FindByTitle(tc.title, tc.collection)
			if len(result) != len(tc.result) {
				t.Errorf("Result was incorrect, got: %v, want: %v.", result, tc.result)
			}
			for i := range result {
				if result[i] != tc.result[i] {
					t.Errorf("Result was incorrect, got: %v, want: %v.", result, tc.result)
				}
			}
		})
	}
}
