package domain

type Opt struct {
	
}

type OptService interface {
	GenerateOTP(email string) (string, error)
}

type PasswordForgotEvent struct {
	Email string `json:"email"`
}