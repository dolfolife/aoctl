package puzzle

import (
    "errors"
)

const (
    Unsolved      PuzzleStatus = "UNSOLVED"
    Solved        PuzzleStatus = "SOLVED"
    Unreachable   PuzzleStatus = "UNREACHABLE"
)

var (
    ErrInvalidStatus = errors.New("invalid status")
)

type InvalidPuzzle struct {
    appError  error
    statusErr error
}

// to implement the error interface we need to implement the Error() method
func (e *InvalidPuzzle) Error() string {
    return errors.Join(e.statusErr, e.appError).Error()
}

func NewError(appError error, statusErr error) *InvalidPuzzle {
    return &InvalidPuzzle{
        appError: appError,
        statusErr: statusErr,
    }
}
