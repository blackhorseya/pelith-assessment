package web

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/**/*
var f embed.FS

var templates = template.Must(template.ParseFS(f, "templates/**/*"))

// SetHTMLTemplate set html template.
func SetHTMLTemplate(r *gin.Engine) {
	r.SetHTMLTemplate(templates)

	r.StaticFS("/static", http.FS(f))
}
