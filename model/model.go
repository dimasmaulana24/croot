package model

type Chat struct {
	Phone_number string `json:"phone_number"`
	Messages     string `json:"messages"`
}

type Response struct {
	Response string `json:"response"`
}
