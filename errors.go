// https://www.vaultproject.io/api#http-status-codes
package govault

import "fmt"

type (
	ErrSuccessNoData       struct{}
	ErrInvalidRequest      struct{}
	ErrForbidden           struct{}
	ErrInvalidPath         struct{}
	ErrStandby             struct{}
	ErrPerformanceStandby  struct{}
	ErrInternalServerError struct{}
	ErrThirdPartyError     struct{}
	ErrSealed              struct{}

	ErrUnknownStatusCode struct {
		StatusCode int
	}
)

func (e *ErrSuccessNoData) Error() string {
	return "Success, no data returned."
}

func (e *ErrInvalidRequest) Error() string {
	return "Invalid request, missing or invalid data."
}

func (e *ErrForbidden) Error() string {
	return "Forbidden, your authentication details are either incorrect, you don't have access to this feature, or - if CORS is enabled - you made a cross-origin request from an origin that is not allowed to make such requests."
}

func (e *ErrInvalidPath) Error() string {
	return "Invalid path. This can both mean that the path truly doesn't exist or that you don't have permission to view a specific path. We use 404 in some cases to avoid state leakage."
}

func (e *ErrStandby) Error() string {
	return "Default return code for health status of standby nodes. This will likely change in the future."
}

func (e *ErrPerformanceStandby) Error() string {
	return "Default return code for health status of performance standby nodes."
}

func (e *ErrInternalServerError) Error() string {
	return "Internal server error. An internal error has occurred, try again later. If the error persists, report a bug."
}

func (e *ErrThirdPartyError) Error() string {
	return "A request to Vault required Vault making a request to a third party; the third party responded with an error of some kind."
}

func (e *ErrSealed) Error() string {
	return "Vault is down for maintenance or is currently sealed. Try again later."
}

func (e *ErrUnknownStatusCode) Error() string {
	return fmt.Sprintf("Unknown status code: %d", e.StatusCode)
}

func checkStatus(code int) error {
	switch code {
	case 200:
		return nil
	case 204:
		return &ErrSuccessNoData{}
	case 400:
		return &ErrInvalidRequest{}
	case 403:
		return &ErrForbidden{}
	case 404:
		return &ErrInvalidPath{}
	case 429:
		return &ErrStandby{}
	case 473:
		return &ErrPerformanceStandby{}
	case 500:
		return &ErrInternalServerError{}
	case 502:
		return &ErrThirdPartyError{}
	case 503:
		return &ErrSealed{}
	default:
		return &ErrUnknownStatusCode{StatusCode: code}
	}
}
