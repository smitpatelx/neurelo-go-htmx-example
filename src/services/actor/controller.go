package actor

import (
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
)

func RenderIndexPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func GetAllActors(c *fiber.Ctx) error {
	var req GetAllActorRequest
	err := c.QueryParser(&req)
	if err != nil || req.Page == 0 {
		return c.Redirect(fmt.Sprintf("/actors?page=1?search=%s", req.Search))
	}

	res := ReadAllActorSvc(req)
	actor_count := GetTotalActorsSvc(req)

	if actor_count == nil {
		return c.Render("index", fiber.Map{})
	}

	var pages []int
	total_page_count := int(math.Round(float64(*actor_count/12))) + 1

	if req.Page > total_page_count {
		// Redirect user back to last page.
		return c.Redirect(fmt.Sprintf("/actors?page=%d&search=%s", total_page_count, req.Search))
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
		return c.Render("actors", fiber.Map{
			"Actors":         res,
			"Page":           req.Page,
			"TotalPages":     pages,
			"TotalPageCount": total_page_count,
			"PerPage":        12,
			"PreviousPage":   prev_page,
			"Search":         req.Search,
		})
	}

	return c.Render("actors", fiber.Map{})
}

func GetAllFilms(c *fiber.Ctx) error {
	var req GetAllFilmRequest
	err := c.QueryParser(&req)
	if err != nil || req.Page == 0 {
		return c.Redirect("/actors?page=1")
	}

	res := ReadAllFilmSvc(req)
	film_count := GetTotalFilmSvc(req)

	if film_count == nil {
		return c.Render("index", fiber.Map{})
	}

	var pages []int
	total_page_count := int(math.Round(float64(*film_count/12))) + 1

	if req.Page > total_page_count {
		// Redirect user back to last page.
		return c.Redirect(fmt.Sprintf("/actors?page=%d", total_page_count))
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
			"Search":         req.Search,
		})
	}

	return c.Render("films", fiber.Map{})
}
