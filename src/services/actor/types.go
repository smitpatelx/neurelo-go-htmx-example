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
