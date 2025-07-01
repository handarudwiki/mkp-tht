package dto

type CreateTerminal struct {
	Name      string  `json:"name" validate:"required"`
	Location  string  `json:"location" validate:"required"`
	Address   string  `json:"address" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}
