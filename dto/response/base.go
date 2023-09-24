package response

// FailureResponse Response - Application response failure struct
type FailureResponse struct {
	Success bool   `json:"success" example:"false"`
	Failure string `json:"failure"`
}

// SuccessResponse Response - Application response success struct
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Result  interface{} `json:"result"`
}
