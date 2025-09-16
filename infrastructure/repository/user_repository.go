package repository

import (
	"context"
	"database/sql"
	"nofu-be/domain"
	"nofu-be/domain/user"
	"nofu-be/infrastructure/database"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository() domain.UserRepository {
	return &UserRepositoryImpl{DB: database.DB}
}

func (r *UserRepositoryImpl) ListAll() ([]user.User, error) {
	query := `SELECT c_username, c_password, c_fullname FROM user_master`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		err = rows.Scan(&u.CUsername, &u.CPassword, &u.CFullname)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Create menyimpan user baru ke database
func (r *UserRepositoryImpl) InsertUser(u user.User) error {
	query := `INSERT INTO user_master (c_username, c_password, c_fullname) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(context.Background(), query, u.CUsername, u.CPassword, u.CFullname)
	return err
}
