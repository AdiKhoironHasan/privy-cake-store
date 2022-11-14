package dto

type ResponseDTO struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CakeResDTO struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
	Created_at  string  `json:"created_at,omitempty"`
	Updated_at  string  `json:"updated_at,omitempty"`
}
