package models

import (
	"errors"
	"time"
)



var ErrNoRecords = errors.New("models: no matching records found")


type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}