package requests

type UbiquitiLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
