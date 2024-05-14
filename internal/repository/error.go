package repository

import "errors"

// ErrRecordNotExist is returned when a requested record does not exist in the repository
var ErrRecordNotExist = errors.New("the record does not exist")
// ErrRecordsNotFound is returned when no records are found in the repository.
var ErrRecordsNotFound = errors.New("no records were found")
