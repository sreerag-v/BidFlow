package response

type ErrResponse struct {
	Data  interface{}
	Error interface{}
	StatusCode int
}

type SuccResponse struct {
	Data       interface{}
	StatusCode int
}
