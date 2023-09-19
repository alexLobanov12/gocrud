package models

type Message struct {
	Message string `json:"message"`
}

type Login struct {
	Login string `json:login`
	Password string `json:password`
}

type LoginError struct {
	Errors [string] `errors`
	Status int 
}

/*
func (l *Login) ValidateLogin() {
	if l.Login == 
}
*/
