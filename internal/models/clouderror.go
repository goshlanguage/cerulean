package models

import "strconv"

// CloudError represents the structure for how the Azure SDK formats error responses
// TODO: Is this used in other services/should it be in a more generalized folder location
// TODO: Finish implementing the Info and Details attributes (not needed for MVP)
type CloudError struct {
	Error struct {
		AdditionalInfo []struct { // The error additional info.
			// Info object `json:"info" // The additional info.
			Type string `json:"type"` // The additional info type.
		} `json:"additionalInfo"`
		Code string `json:"code"` // The error code.
		// Details ErrorResponse[] `json:"details" // The error details.
		Message string `json:"message"` // The error message.
		Target  string `json:"target"`  // The error target.
	} `json:"error"`
}

// NewCloudError constructs a new Azure-style CloudError to return on API error
func NewCloudError(httpErrorCode int, err error) CloudError {
	var errorResponse CloudError

	errorResponse.Error.Code = strconv.Itoa(httpErrorCode)
	errorResponse.Error.Message = err.Error()

	return errorResponse
}
