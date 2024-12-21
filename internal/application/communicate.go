package application

type Request struct {
	Expression string `json:"expression"`
}

type ErrResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Result float64 `json:"result"`
}
