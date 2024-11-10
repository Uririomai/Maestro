package model

import "errors"

// repository
var (
	ErrEmailRegistered      = errors.New("email is already registered")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
	ErrAliasTaken           = errors.New("alias is already taken")
	ErrEmptyOrder           = errors.New("order is empty")
	ErrInvalidEmail         = errors.New("email is invalid")
	ErrInvalidImagesId      = errors.New("invalid images id")
	ErrAdminHaveWebsite     = errors.New("admin already have website")
	ErrInvalidActive        = errors.New("invalid active status")
	ErrInvalidNotification  = errors.New("invalid notification status")
	ErrWebsiteNotFound      = errors.New("website not found")
)
