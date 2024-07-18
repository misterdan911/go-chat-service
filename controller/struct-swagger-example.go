package controller

type ExampleSignedUpUser struct {
	Username       string `json:"username" example:"danu"`
	Email          string `json:"email" example:"dciptadi@gmail.com"`
	FirstName      string `json:"firstName" example:"Danu"`
	LastName       string `json:"lastName" example:"Ciptadi"`
	Password       string `json:"password" example:"12345678"`
	RepeatPassword string `json:"repeatPassword" example:"12345678"`
}

type ExampleSignedInUser struct {
	Username string `json:"username" example:"danu"`
	Email    string `json:"email" example:"dciptadi@gmail.com"`
	Password string `json:"password" example:"12345678"`
}
