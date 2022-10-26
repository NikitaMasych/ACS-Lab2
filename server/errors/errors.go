package errors

import "github.com/pkg/errors"

var (
	ErrNoMessage         = errors.New("has no open or closed quote for message")
	ErrNoSortType        = errors.New("has no specified sort type")
	ErrInvalidSortType   = errors.New("undefined sort type")
	ErrNoStartIndex      = errors.New("has no start index")
	ErrInvalidLength     = errors.New("invalid length")
	ErrInvalidStartIndex = errors.New("invalid start index")
)
