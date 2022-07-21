package telegram

import "errors"

var (
	errArgCount     = errors.New("bad argument count")
	errBadYearBirth = errors.New("bad year of birth")
	errBadId        = errors.New("bad request")
)
