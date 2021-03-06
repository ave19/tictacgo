package tictacgo

// ErrBadSquare for when someone picks the 10th square of the board
type ErrBadSquare struct {
	message string
}

// NewErrBadSquare creates an err object
func NewErrBadSquare(message string) *ErrBadSquare {
	return &ErrBadSquare{
		message: message,
	}
}

func (e *ErrBadSquare) Error() string {
	return e.message
}

// ValueError What to pass when something's not right
type ValueError struct {
	message string
}

func (e *ValueError) Error() string {
	return e.message
}
