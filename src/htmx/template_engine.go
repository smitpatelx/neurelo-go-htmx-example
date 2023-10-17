package htmx

import (
	"github.com/gofiber/template/html/v2"
)

func GetTemplateEngine() *html.Engine {
	// Create a new engine
	engine := html.New("./src/views", ".html")

	if engine == nil {
		panic("Could not load template engine")
	}

	return engine
}
