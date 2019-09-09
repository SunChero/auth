package service

import (
	"crypto/rsa"
	"database/sql"
)

//Service service backend
type Service struct {
	db      *sql.DB
	signKey *rsa.PrivateKey
}

//New create a new service
func New(db *sql.DB, signKey *rsa.PrivateKey) *Service {
	s := &Service{db, signKey}

	return s
}
