package api

import ()

type Error interface {
	error
	StatusCode() int
	StatusMessage() string
}
