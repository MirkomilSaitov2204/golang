package models

var DB []Book

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	Lastname string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		Id:            1,
		Title:         "Title-1",
		YearPublished: 1111,
		Author: Author{
			Name:     "Author-1",
			Lastname: "Author-1",
			BornYear: 1111,
		},
	}

	book2 := Book{
		Id:            2,
		Title:         "Title-2",
		YearPublished: 2222,
		Author: Author{
			Name:     "Author-2",
			Lastname: "Author-2",
			BornYear: 2222,
		},
	}

	book3 := Book{
		Id:            3,
		Title:         "Title-3",
		YearPublished: 3333,
		Author: Author{
			Name:     "Author-3",
			Lastname: "Author-3",
			BornYear: 3333,
		},
	}

	DB = append(DB, book1, book2, book3)
}

func FindBookById(id int) (Book, bool) {
	var book Book
	var found bool

	for _, b := range DB {
		if b.Id == id {
			book = b
			found = true
			break
		}
	}
	return book, found
}
