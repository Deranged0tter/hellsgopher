package hellsgopher

import "errors"

var (
	ErrNotExist         = errors.New("file does not exist")
	ErrUnknownGOOS      = errors.New("unknown GOOS")
	ErrNotWin           = errors.New("only works on windows")
	ErrFuncNotSupported = errors.New("this function is not supported")
	ErrBindAccept       = errors.New("unable to accept bind connection")
)
