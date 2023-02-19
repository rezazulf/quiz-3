package process

import (
	"database/sql"
	"quiz-3/database"
)

func GetAllCategory(db *sql.DB) (err error, results []database.Category) {
	query := "SELECT id, name FROM category ORDER BY ID ASC"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = database.Category{}

		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category *database.Category) (err error) {
	query := "INSERT INTO category (name) VALUES ($1)"
	errs := db.QueryRow(query, category.Name)
	return errs.Err()

}

func UpdateCategory(db *sql.DB, category database.Category) (err error) {
	query := "UPDATE category SET name = $1 WHERE id = $2"
	errs := db.QueryRow(query, category.Name, category.ID)
	return errs.Err()

}

func DeleteCategory(db *sql.DB, category *database.Category) (err error) {
	query := "DELETE FROM category WHERE id = $1"
	errs := db.QueryRow(query, category.ID)
	return errs.Err()
}

func GetCategoryBook(db *sql.DB, id_book *database.Book) (err error, results []database.Book) {
	query := "SELECT id, title, description, image_url, release_year, price, total_page, thickness FROM book WHERE id = $1"
	rows, err := db.Query(query, id_book.ID)
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
			&book.Thickness)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}
