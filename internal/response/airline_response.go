package response

type Airline struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Logo_url string `json:"logo_url"`
}

type AirlineResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Airline `json:"data"`
}

type AirlinesResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Count   int       `json:"count"`
	Data    []Airline `json:"data"`
}
