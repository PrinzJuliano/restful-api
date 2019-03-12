package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prinzjuliano/restful-api/routes"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

// loadTemplate loads templates embedded by go-assets-builder
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func main() {
	r := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.Static("/assets", "./public/assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{"a":"b"})
	})
	r.GET("/ping", routes.Ping)
	r.Run(":12345")
}
