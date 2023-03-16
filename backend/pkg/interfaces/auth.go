package interfaces

type AuthInterface struct {
	Password string `json:"password" validate:"required,min=8"`
	Mail     string `json:"mail" validate:"required,email"`
}
