package models

type Terminal struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Location  string  `json:"location"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UpdatedAt string  `json:"updated_at"`
	CreatedAt string  `json:"created_at"`
}
