package controllers

import (
	"net/http"
	"net/url"
	"quiz-3/database"
	"quiz-3/process"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	books []database.Book
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	var err, books = process.GetAllBook(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book database.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = int64(len(*&books) + 1)

	// cek thickness
	if book.Total_page <= 100 {
		book.Thickness = "tipis"
	} else if book.Total_page > 100 && book.Total_page <= 200 {
		book.Thickness = "sedang"
	} else {
		book.Thickness = "tebal"
	}

	// cek url
	_, err2 := url.ParseRequestURI(book.Image_url)
	if err2 != nil {
		url, _ := regexp.MatchString("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)", book.Image_url)
		if !url {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "url tidak sesuai",
			})
		}
	}

	// cek tahun
	if book.Release_year <= 1980 && book.Release_year >= 2021 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tahun harus antara 1980-2021",
		})
	}

	err = process.InsertBook(database.DbConnection, &book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Insert Book Sukses",
	})

}

func UpdateBook(c *gin.Context) {
	var book database.Book

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = int64(id)

	// cek thickness
	if book.Total_page <= 100 {
		book.Thickness = "Tipis"
	} else if book.Total_page > 100 && book.Total_page <= 200 {
		book.Thickness = "Sedang"
	} else {
		book.Thickness = "Tebal"
	}

	// cek url
	_, err2 := url.ParseRequestURI(book.Image_url)
	if err2 != nil {
		url, _ := regexp.MatchString("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)", book.Image_url)
		if !url {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "URL Salah!",
			})
		}
	}

	// cek tahun
	if book.Release_year <= 1980 && book.Release_year >= 2021 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tahun harus antara 1980-2021",
		})
	}

	err = process.UpdateBook(database.DbConnection, &book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Update Buku Sukses",
	})
}

func DeleteBook(c *gin.Context) {
	var book database.Book

	id, err := strconv.Atoi(c.Param("id"))

	book.ID = int64(id)

	err = process.DeleteBook(database.DbConnection, &book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Hapus Buku Sukses",
	})
}
