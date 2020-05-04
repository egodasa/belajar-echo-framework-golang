package api

import (
	md "belajar-ech0-framework/models"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

type CoreApi struct {
	HttpStatus int         `json:"code"`              // field untuk kode http
	Messages   string      `json:"message,omitempty"` // field untuk pesan
	Data       interface{} `json:"data,omitempty"`    // field yang menamping data yang akan diberikan ke klien
	ApiModels  md.IModels  `json:"-"`
}

func (c CoreApi) Get(ctx *gin.Context) {
	var data, err = c.ApiModels.All()
	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam mengambil data"
	} else {
		c.Data = data
	}
	ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Find(ctx *gin.Context) {
	var id string = ctx.Param("id")
	var data, err = c.ApiModels.Find(id)

	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam mencari data"
	} else {
		c.Data = data
	}

	ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Insert(ctx *gin.Context) {
	data := make(map[string]string)

	// data yang dimasukkan hanyalah data yang sudah ditentukan di ColumnList
	// serta data yang ada nilainya
	columnList := c.ApiModels.GetColumnList()
	for _, value := range columnList {
		if ctx.PostForm(value.Name) != "" {
			data[value.Name] = string(ctx.PostForm(value.Name))
		}
	}

	// proses insert
	err := c.ApiModels.Insert(data).Run()

	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan saat menambahkan data"
	}

	ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Update(ctx *gin.Context) {
	var id string = ctx.Param("id")

	data := make(map[string]string)

	// data yang dimasukkan hanyalah data yang sudah ditentukan di ColumnList
	// serta data yang ada nilainya
	columnList := c.ApiModels.GetColumnList()
	for _, value := range columnList {
		if ctx.PostForm(value.Name) != "" {
			data[value.Name] = ctx.PostForm(value.Name)
		}
	}

	// proses insert
	err := c.ApiModels.Update(id, data).Run()

	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan saat mengupdate data"
	}

	ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Delete(ctx *gin.Context) {
	var id string = ctx.Param("id")
	var data, err = c.ApiModels.Find(id)
	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam menghapus data"
	} else {
		if data == nil {
			c.HttpStatus = http.StatusNotFound
			c.Messages = "Data tidak ditemukan"
		} else {
			// proses insert
			err = c.ApiModels.Delete(id).Run()

			if err == true {
				c.HttpStatus = http.StatusInternalServerError
				c.Messages = "Terjadi kesalahan. Tidak dapat menyimpan data!"
			}
		}
	}

	ctx.JSON(c.HttpStatus, c)
}
