package telegram

import "errors"

var (
	errCreateArgCount = errors.New("bad argument count for create")
	errBadYearBirth   = errors.New("bad year of birth")
	errBadId          = errors.New("bad request")
)
