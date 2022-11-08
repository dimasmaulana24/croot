package model

type Chat struct {
	Phone_number string `json:"phone_number"`
	Messages     string `json:"messages"`
}

type Response struct {
	Response string `json:"response"`
}

type Gis struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name"`
}
