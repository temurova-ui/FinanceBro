package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/temurova-ui/FinanceBro/internal/model"
)

type UserRepo struct{
	db *pgxpool.Pool
}

func NewUserRepo (db *pgxpool.Pool)*UserRepo{
	return &UserRepo{db: db,}
}

func (u *UserRepo) Create (ctx context.Context, user model.User)(int, error){
	const query = `INSERT INTO users (name, email, password_hash) 
					VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := u.db.QueryRow(ctx, query, user.Name, user.Email, user.PasswordHash).Scan(&id)
	if err != nil{
		return 0, err
	}
	return id, nil
}

func (u *UserRepo) ExistByEmail(ctx context.Context, email string)(bool, error){
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1) BY email `
	var exists bool
	err := u.db.QueryRow(ctx, query, email).Scan(&exists)
	if err !=  nil{
		return false, fmt.Errorf("u.db.QueryRow: %w", err)
	}
	return exists, nil
}