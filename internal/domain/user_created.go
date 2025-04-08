package domain

type UserCreated struct {
	Username      string
	Email         string
	Code          string
	CodeExpiresAt string
}
