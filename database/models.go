package database

type Category struct {
	ID   int64  `json:"category_id"`
	Name string `json:"name"`
}

type Book struct {
	ID           int64  `json:"book_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image_url    string `json:"image_url"`
	Release_year int64  `json:"release_year"`
	Price        string `json:"price"`
	Total_page   int64  `json:"total_page"`
	Thickness    string `json:"thickness"`
	Category_id  int64  `json:"category_id"`
}
