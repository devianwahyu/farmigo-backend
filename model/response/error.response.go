package response

type ErrorResponse struct {
	FailedFields string
	Tag          string
	Value        string
}
