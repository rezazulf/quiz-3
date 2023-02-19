package process

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SegitigaSamaSisi(c *gin.Context) {
	var (
		result gin.H
	)

	tinggi, err := strconv.Atoi(c.Query("tinggi"))
	if err != nil || tinggi < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params tinggi seharusnya angka",
		})
		return
	}

	alas, err := strconv.Atoi(c.Query("alas"))
	if err != nil || alas < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params alas seharusnya angka",
		})
		return
	}

	hitung := c.Query("hitung")
	if hitung == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Terdapat kesalahan pada params. Params hitung dibutuhkan",
		})
		return

	} else if hitung != "keliling" && hitung != "luas" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params hitung seharusnya luas atau keliling",
		})
		return

	}
	if hitung == "luas" {
		result = gin.H{
			"luas": (alas * tinggi) / 2,
		}
	} else if hitung == "keliling" {
		result = gin.H{
			"keliling": 3 * alas,
		}
	}

	c.JSON(http.StatusOK, result)
}

func Persegi(c *gin.Context) {
	var (
		result gin.H
	)

	// sisi, err := strconv.Atoi(c.Query("sisi"))
	sisi, err := strconv.Atoi(c.Query("sisi"))
	if err != nil || sisi < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params sisi seharusnya angka",
		})
		return
	}
	if err != nil {
		panic(err)
	}

	hitung := c.Query("hitung")
	if hitung == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Terdapat kesalahan pada params. Params hitung dibutuhkan",
		})
		return

	} else if hitung != "keliling" && hitung != "luas" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params hitung seharusnya luas atau keliling",
		})
		return

	}

	if hitung == "luas" {
		result = gin.H{
			"luas": sisi * sisi,
		}
	} else if hitung == "keliling" {
		result = gin.H{
			"keliling": 4 * sisi,
		}
	}

	c.JSON(http.StatusOK, result)
}

func PersegiPanjang(c *gin.Context) {
	var (
		result gin.H
	)

	panjang, err := strconv.Atoi(c.Query("panjang"))
	if err != nil || panjang < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params panjang seharusnya angka",
		})
		return
	}

	lebar, err := strconv.Atoi(c.Query("lebar"))
	if err != nil || lebar < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params lebar seharusnya angka",
		})
		return
	}

	hitung := c.Query("hitung")

	if hitung == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Terdapat kesalahan pada params. Params hitung dibutuhkan",
		})
		return

	} else if hitung != "keliling" && hitung != "luas" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params hitung seharusnya luas atau keliling",
		})
		return

	}

	if hitung == "luas" {
		result = gin.H{
			"luas": panjang * lebar,
		}
	} else if hitung == "keliling" {
		result = gin.H{
			"keliling": 2 * (panjang + lebar),
		}
	}

	c.JSON(http.StatusOK, result)
}

func Lingkaran(c *gin.Context) {
	var (
		result gin.H
	)

	jariJari, err := strconv.Atoi(c.Query("jariJari"))
	if err != nil || jariJari < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params jariJari seharusnya angka",
		})
		return
	}

	hitung := c.Query("hitung")

	if hitung == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Terdapat kesalahan pada params. Params hitung dibutuhkan",
		})
		return

	} else if hitung != "keliling" && hitung != "luas" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Params hitung seharusnya luas atau keliling",
		})
		return

	}

	if hitung == "luas" {
		result = gin.H{
			"luas": math.Pi * float64(jariJari*jariJari),
		}
	} else if hitung == "keliling" {
		result = gin.H{
			"keliling": 2 * math.Pi * float64(jariJari),
		}
	}

	c.JSON(http.StatusOK, result)
}
