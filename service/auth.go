package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	verificationCodeLifespan = time.Minute * 15
	tokenLifespan            = time.Hour * 24 * 14
)

//SendVerificationCode send magic link
func (s *Service) CreateVerificationCode(ctx context.Context, email string) (string, error) {
	email = strings.TrimSpace(email)
	if !reEmail.MatchString(email) {
		return "", ErrInvalidEmail
	}
	return s.sqlCreateVerificationCode(ctx, email)
}

// AuthURI to be redirected to and complete the login flow.
// It contains the token in the hash fragment.
func (s *Service) LoginWithCode(ctx context.Context, verificationCode string) (string, error) {
	log.Printf(`verification code %v `, verificationCode)
	verificationCode = strings.TrimSpace(verificationCode)
	if !reUUID.MatchString(verificationCode) {
		return "", ErrInvalidVerificationCode
	}
	var uid string
	var createdAt time.Time
	err := s.db.QueryRowContext(ctx, `
		DELETE FROM verification_codes WHERE id = $1
		RETURNING auth_id, created_at`, verificationCode).Scan(&uid, &createdAt)
	// if err == sql.ErrNoRows {
	// 	return "", ErrVerificationCodeNotFound
	// }

	if err != nil {
		return "", fmt.Errorf("could not delete verification code: %v", err)
	}

	now := time.Now()
	exp := createdAt.Add(verificationCodeLifespan)
	if exp.Equal(now) || exp.Before(now) {
		return "", ErrExpiredToken
	}
	return s.createJWT(uid)
}
