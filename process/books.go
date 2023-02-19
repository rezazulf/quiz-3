package process

import (
	"database/sql"
	"quiz-3/database"
)

// router menampilkan semua kategori
func GetAllBook(db *sql.DB) (err error, results []database.Book) {
	query := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM book"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = database.Book{}

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Category_id)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}

// mengedit kategori
func InsertBook(db *sql.DB, book *database.Book) (err error) {
	query := "INSERT INTO book (id, title, description, image_url, release_year, price, total_page, thickness, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	errs := db.QueryRow(
		query,
		book.ID,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Thickness,
		book.Category_id)

	return errs.Err()

}

// mengupdate book
func UpdateBook(db *sql.DB, book *database.Book) (err error) {
	query := "UPDATE book SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8 WHERE id = $9"

	errs := db.QueryRow(
		query,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Thickness,
		book.Category_id,
		book.ID)

	return errs.Err()

}

// mengahpus kategory
func DeleteBook(db *sql.DB, book *database.Book) (err error) {
	query := "DELETE FROM book WHERE id = $1"

	errs := db.QueryRow(query, book.ID)

	return errs.Err()
}
