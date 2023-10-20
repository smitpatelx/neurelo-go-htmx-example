package actor

import (
	"fmt"
	"math"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetAllActorsAPI(c *fiber.Ctx) error {
	var req GetAllActorRequest
	err := c.QueryParser(&req)
	if err != nil || req.Page == 0 {
		return c.Redirect("/api/v1/actors?page=1")
	}

	res := ReadAllActorSvc(req)
	actor_count := GetTotalActorsSvc(req)

	if actor_count == nil {
		return c.Render("partials/actor_data", fiber.Map{})
	}

	var pages []int
	total_page_count := int(math.Round(float64(*actor_count/12))) + 1

	if req.Page > (total_page_count) {
		// Redirect user back to last page.
		return c.Redirect(
			fmt.Sprintf("/api/v1/actors?page=%d&search=%s",
				total_page_count,
				strings.TrimSpace(req.Search),
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
		return c.Render("partials/actor_data", fiber.Map{
			"Actors":         res,
			"Page":           req.Page,
			"TotalPages":     pages,
			"TotalPageCount": total_page_count,
			"PerPage":        12,
			"PreviousPage":   prev_page,
			"Search":         req.Search,
		})
	}

	return c.Render("partials/actor_data", fiber.Map{})
}
