package controllers

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/process"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	err, categories := process.GetAllCategory(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category database.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = process.InsertCategory(database.DbConnection, &category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Insert Category Sukses",
	})

}

func UpdateCategory(c *gin.Context) {
	var category database.Category

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = int64(id)

	err = process.UpdateCategory(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Update Kategori Sukses",
	})
}

func DeleteCategory(c *gin.Context) {
	var category database.Category

	id, err := strconv.Atoi(c.Param("id"))

	category.ID = int64(id)

	err = process.DeleteCategory(database.DbConnection, &category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Hapus Kategori Sukses",
	})
}

func GetCategoryBook(c *gin.Context) {
	var (
		result gin.H
		book   database.Book
	)

	id, err := strconv.Atoi(c.Param("id"))

	err = c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = int64(id)

	err, categories := process.GetCategoryBook(database.DbConnection, &book)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)

}
