package service

import (
	"errors"
	"regexp"
)

var (
	reEmail = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

//	reUsername = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,17}$`)
//	avatarsDir = path.Join("web", "static", "img", "avatars")
)

var (
	// ErrInvalidUserID denotes an invalid user id; that is not uuid.
	//	ErrInvalidUserID = errors.New("invalid user id")
	// ErrInvalidEmail denotes an invalid email address.
	ErrInvalidEmail = errors.New("invalid email")
	// ErrInvalidUsername denotes an invalid username.
	// ErrInvalidUsername = errors.New("invalid username")
	// // ErrEmailTaken denotes an email already taken.
	// ErrEmailTaken = errors.New("email taken")
	// // ErrUsernameTaken denotes a username already taken.
	// ErrUsernameTaken = errors.New("username taken")
	// // ErrUserNotFound denotes a not found user.
	// ErrUserNotFound = errors.New("user not found")
	// // ErrForbiddenFollow denotes a forbiden follow. Like following yourself.
	// ErrForbiddenFollow = errors.New("forbidden follow")
	// // ErrUnsupportedAvatarFormat denotes an unsupported avatar image format.
	// ErrUnsupportedAvatarFormat = errors.New("unsupported avatar format")
	// ErrUnimplemented denotes a not implemented functionality.
	ErrUnimplemented = errors.New("unimplemented")
	// ErrUnauthenticated denotes no authenticated user in context.
	ErrUnauthenticated = errors.New("unauthenticated")
	// ErrInvalidRedirectURI denotes an invalid redirect uri.
	ErrInvalidRedirectURI = errors.New("invalid redirect uri")
	// ErrInvalidToken denotes an invalid token.
	ErrInvalidToken = errors.New("invalid token")
	// ErrExpiredToken denotes that the token already expired.
	ErrExpiredToken = errors.New("expired token")
	// ErrInvalidVerificationCode denotes an invalid verification code.
	ErrInvalidVerificationCode = errors.New("invalid verification code")
	// ErrVerificationCodeNotFound denotes a not found verification code.
	ErrVerificationCodeNotFound = errors.New("verification code not found")
	reUUID                      = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
)
