package request

type SignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
