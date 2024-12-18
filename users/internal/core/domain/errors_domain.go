package domain

import "errors"

var (
	ErrorDataNotFound  = errors.New("data not found")
	ErrorConflictData  = errors.New("data conflicts with existing data")
	ErrorNoUpdatedData = errors.New("no data to update")
	ErrorInternal      = errors.New("internal server error")
)
