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
type LoginResult struct {
	Token    string `json:"token"`
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
