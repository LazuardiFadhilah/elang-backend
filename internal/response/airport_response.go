package response

type Airport struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type AirportResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Airport `json:"data"`
}
