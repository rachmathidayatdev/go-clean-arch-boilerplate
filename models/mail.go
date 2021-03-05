package models

//SignupMailData struct
type SignupMailData struct {
	Email          string `json:"email"`
	ButtonLink     string `json:"button_link"`
	Token          string `json:"token"`
	CodeActivation string `json:"codeActivation"`
	Template       string `json:"template"`
}

//ForgotPasswordMailData struct
type ForgotPasswordMailData struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	ButtonLink string `json:"button_link"`
	Token      string `json:"token"`
	Template   string `json:"template"`
}
