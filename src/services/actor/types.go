package actor

type Actor struct {
	ActorID    int    `json:"actor_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	LastUpdate string `json:"last_update"`
}

type Client_GetAllActorResponse struct {
	Data []Actor `json:"data"`
}

type GetAllActorResponse struct {
	*Client_GetAllActorResponse
	Page           int   `json:"page"`
	TotalPages     []int `json:"total_pages"`
	TotalPageCount int   `json:"total_page_count"`
	PerPage        int   `json:"per_page"`
	PreviousPage   int   `json:"previous_page"`
}

type Client_CountActorResponse struct {
	ActorId int `json:"actor_id"`
}

type Client_GetActorResponse struct {
	Count Client_CountActorResponse `json:"_count"`
}

type Client_GetTotalActorResponse struct {
	Data Client_GetActorResponse `json:"data"`
}

type GetAllActorRequest struct {
	Page   int    `query:"page" json:"page"`
	Search string `query:"search" json:"search"`
}

// Film

type GetAllFilmRequest struct {
	Page   int    `query:"page" json:"page"`
	Search string `query:"search" json:"search"`
}

type Film struct {
	Description     string   `json:"description"`
	FilmId          int      `json:"film_id"`
	LanguageId      int      `json:"language_id"`
	LastUpdate      string   `json:"last_update"`
	Length          int      `json:"length"`
	ReleaseYear     int      `json:"release_year"`
	RentalDuration  int      `json:"rental_duration"`
	RentalRate      string   `json:"rental_rate"`
	ReplacementCost string   `json:"replacement_cost"`
	SpecialFeatures []string `json:"special_features"`
	Title           string   `json:"title"`
}

type Client_GetAllFilmResponse struct {
	Data []Film `json:"data"`
}

type GetAllFilmResponse struct {
	*Client_GetAllFilmResponse
	Page           int   `json:"page"`
	TotalPages     []int `json:"total_pages"`
	TotalPageCount int   `json:"total_page_count"`
	PerPage        int   `json:"per_page"`
	PreviousPage   int   `json:"previous_page"`
}

type Client_CountFilmResponse struct {
	FilmId int `json:"film_id"`
}

type Client_GetFilmResponse struct {
	Count Client_CountFilmResponse `json:"_count"`
}

type Client_GetTotalFilmResponse struct {
	Data Client_GetFilmResponse `json:"data"`
}
