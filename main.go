package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", indexHandler)
	r.POST("/", uploadHandler)
	r.GET("/download/:filename", downloadHandler)
	err := r.Run(":80")
	if err != nil {
		panic(err)
	}
}

func indexHandler(c *gin.Context) {
	files, err := filepath.Glob(filepath.Join("files", "*"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var links []string
	for _, file := range files {
		links = append(links,  strings.TrimPrefix(file, "files/"))
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Files": links,
	})
}

func downloadHandler(c *gin.Context) {
	path := c.Param("filename")
	file, err := os.Open("files/"+path)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename="+fi.Name())
	c.Header("Content-Type", "application/octet-stream")
	io.Copy(c.Writer, file)
}

func uploadHandler(c *gin.Context) {
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	// 创建一个名为 files 的目录
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		os.Mkdir("files", os.ModePerm)
	}
	f, err := os.OpenFile("files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	c.Redirect(http.StatusFound, "/")
}
