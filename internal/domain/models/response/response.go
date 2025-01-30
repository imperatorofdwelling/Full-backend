package response

type (
	ResponseError struct {
		Error string `json:"error"`
	} // @name ResponseError

	ResponseSuccess struct {
		Data interface{} `json:"data"`
	} // @name ResponseSuccess
)
