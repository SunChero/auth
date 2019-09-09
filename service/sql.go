package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/markbates/goth"
)

func (s *Service) checkUser(ctx context.Context, guser *goth.User) (string, error) {
	var id string
	log.Printf(`%#v`, guser)
	loginID := fmt.Sprintf(`%v$%v`, guser.Provider, guser.Email)
	log.Printf(`this is the loginID :  %v`, loginID)
	query := "SELECT login_id FROM  auth WHERE auth_id = $1"
	err := s.db.QueryRowContext(ctx, query, loginID).Scan(&id)
	if err == sql.ErrNoRows {
		query = "INSERT INTO auth (auth_id) VALUES ($1)"
		tx, err := s.db.BeginTx(ctx, nil)
		if err != nil {
			return "", fmt.Errorf("tx transaction error: %v", err)
		}
		_, err = tx.ExecContext(ctx, query, loginID)
		if err != nil {
			return "", fmt.Errorf("could not insert user: %v", err)
		}
		if err = tx.Commit(); err != nil {
			return "", fmt.Errorf("could not commit insert user: %v", err)
		}

		return loginID, nil
	}
	if err != nil {
		log.Println(err)
	}
	if id != "" {
		fmt.Printf(`this is the current id : %v`, id)
		return loginID, nil

	}
	return "", nil

}

func (s *Service) sqlCreateVerificationCode(ctx context.Context, email string) (string, error) {
	var code string
	err := s.db.QueryRowContext(ctx, `INSERT INTO verification_codes (email, auth_id) VALUES ($1,
		(SELECT auth_id from auth  WHERE email = $1)
		) RETURNING id`, email).Scan(&code)
	// if isForeignKeyViolation(err) {
	// 	return ErrUserNotFound
	// }
	if err != nil {
		return "", fmt.Errorf("could not insert verification code: %v", err)
	}

	return code, nil
}
