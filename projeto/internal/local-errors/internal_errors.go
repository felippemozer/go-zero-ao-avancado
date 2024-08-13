package localerrors

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrNotFound error = errors.New("not found")
