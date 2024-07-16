package controller

type ExampleSignedUpUser struct {
	Username    string `json:"username" example:"danu"`
	Email       string `json:"email" example:"dciptadi@gmail.com"`
	Password    string `json:"password" example:"12345678"`
	DisplayName string `json:"display_name" example:"Danu Ciptadi"`
}

type ExampleSignedInUser struct {
	Username string `json:"username" example:"danu"`
	Email    string `json:"email" example:"dciptadi@gmail.com"`
	Password string `json:"password" example:"12345678"`
}
