package models

// Satellite is the model
type Satellite struct {
	Message  []string `json:"message" validate:"required"`
	Name     string   `json:"name" validate:"required"`
	Distance float64  `json:"distance" validate:"required"`
}

type Satellites struct {
	Satellites []Satellite `json:"satellites" validate:"required"`
}
