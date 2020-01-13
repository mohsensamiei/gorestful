package restful

import "github.com/pkg/errors"

var (
	ErrNotAcceptable        = errors.New("unsupported request accept header")
	ErrUnsupportedMediaType = errors.New("unsupported request content-type header")
	ErrUnauthorized         = errors.New("invalid request authorization header")
	ErrValueDoesNotExist    = errors.New("request value does not exists")
)
