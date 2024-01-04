package response

type ErrResponse struct {
	Response  interface{}
	Error interface{}
	StatusCode int
}

type SuccResponse struct {
	Response       interface{}
	StatusCode int
}

type LoginRes struct{
	TokenString interface{}
	StatusCode int
}