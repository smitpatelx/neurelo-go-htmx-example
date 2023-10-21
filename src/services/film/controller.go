package film

import (
	"fmt"
	"math"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAllFilms(c *fiber.Ctx) error {
	var req GetAllFilmRequest

	trimmed_search := ""

	err := c.QueryParser(&req)
	if err != nil || req.Page == 0 {
		return c.Redirect(fmt.Sprintf("/films?page=1&search=%s", trimmed_search))
	}

	trimmed_search = strings.TrimSpace(req.Search)

	res := ReadAllFilmSvc(req)
	film_count := GetTotalFilmSvc(req)

	if film_count == nil {
		return c.Render("index", fiber.Map{})
	}

	var pages []int
	total_page_count := int(math.Round(float64(*film_count/12))) + 1

	if req.Page > total_page_count {
		// Redirect user back to last page.
		return c.Redirect(fmt.Sprintf("/films?page=%d&search=%s", total_page_count, trimmed_search))
	}

	for i := 1; i <= total_page_count; i++ {
		if (i > req.Page-1 && i <= req.Page+5) || (i > (total_page_count-5) && i <= total_page_count) {
			pages = append(pages, i)
		}
	}

	prev_page := req.Page - 1
	if prev_page < 1 {
		prev_page = 1
	}

	if res != nil {
		return c.Render("films", fiber.Map{
			"Films":          res,
			"Page":           req.Page,
			"TotalPages":     pages,
			"TotalPageCount": total_page_count,
			"PerPage":        12,
			"PreviousPage":   prev_page,
			"Search":         trimmed_search,
		})
	}

	return c.Render("films", fiber.Map{})
}

func GetAllFilmsAPI(c *fiber.Ctx) error {
	var req GetAllFilmRequest
	trimmed_search := ""

	err := c.QueryParser(&req)
	if err != nil || req.Page == 0 {
		return c.Redirect(fmt.Sprintf("/api/v1/films?page=1&search=%s", trimmed_search))
	}

	trimmed_search = strings.TrimSpace(req.Search)

	res := ReadAllFilmSvc(req)
	film_count := GetTotalFilmSvc(req)

	if film_count == nil {
		return c.Render("partials/film_data", fiber.Map{})
	}

	var pages []int
	total_page_count := int(math.Round(float64(*film_count/12))) + 1

	if req.Page > (total_page_count) {
		// Redirect user back to last page.
		return c.Redirect(
			fmt.Sprintf("/api/v1/films?page=%d&search=%s",
				total_page_count,
				trimmed_search,
			))
	}

	for i := 1; i <= total_page_count; i++ {
		if (i > req.Page-1 && i <= req.Page+5) || (i > (total_page_count-5) && i <= total_page_count) {
			pages = append(pages, i)
		}
	}

	prev_page := req.Page - 1
	if prev_page < 1 {
		prev_page = 1
	}

	if res != nil {
		return c.Render("partials/film_data", fiber.Map{
			"Films":          res,
			"Page":           req.Page,
			"TotalPages":     pages,
			"TotalPageCount": total_page_count,
			"PerPage":        12,
			"PreviousPage":   prev_page,
			"Search":         trimmed_search,
		})
	}

	return c.Render("partials/film_data", fiber.Map{})
}
