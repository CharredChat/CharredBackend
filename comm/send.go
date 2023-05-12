package comm

type TokenResponse struct {
	Token string `json:"token"`
}

type MessageResponse struct {
	ID uint64 `json:"id"`
}
