package user

type RegisterCommand struct {
	Username string
	Email    string
	Password string
}
type LoginCommand struct {
	Email    string
	Password string
}
