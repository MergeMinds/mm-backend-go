package user

import (
	"context"
	"errors"
	"time"

	"github.com/InTeam-Russia/go-backend-template/internal/auth/password"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PGRepo struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewPGRepo(db *pgxpool.Pool, logger *zap.Logger) Repo {
	return &PGRepo{db, logger}
}

const createUserSql = `
	INSERT INTO users (first_name, last_name, username, email, role, password_hash, password_salt, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, created_at
`

func (r *PGRepo) Create(user *CreateModel) (*Model, error) {
	passwordSalt, err := password.GenerateSalt()
	if err != nil {
		return nil, nil
	}

	passwordHash := password.Hash(user.Password, passwordSalt)
	r.logger.Debug("Executing query", zap.String("query", createUserSql))

	var newUser Model
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.Username = user.Username
	newUser.Email = user.Email
	newUser.Role = user.Role
	newUser.PasswordHash = passwordHash
	newUser.PasswordSalt = passwordSalt

	err = r.db.QueryRow(
		context.Background(),
		createUserSql,
		newUser.FirstName,
		newUser.LastName,
		newUser.Username,
		newUser.Email,
		newUser.Role,
		newUser.PasswordHash,
		newUser.PasswordSalt,
		time.Now(),
	).Scan(&newUser.Id, &newUser.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

const getByIdSql = `
	SELECT id, first_name, last_name, username, email, role, password_hash, password_salt, created_at
	FROM users
	WHERE id = $1
`

func (r *PGRepo) GetById(id int64) (*Model, error) {
	r.logger.Debug("Executing query", zap.String("query", getByIdSql))

	var user Model
	row := r.db.QueryRow(context.Background(), getByIdSql, id)

	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
		&user.PasswordSalt,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

const getByUsernameSql = `
	SELECT id, first_name, last_name, username, email, role, password_hash, password_salt, created_at
	FROM users
	WHERE username = $1
`

func (r *PGRepo) GetByUsername(username string) (*Model, error) {
	r.logger.Debug("Executing query", zap.String("query", getByUsernameSql))

	var user Model
	row := r.db.QueryRow(context.Background(), getByUsernameSql, username)

	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
		&user.PasswordSalt,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

const deleteByIdSql = `
	DELETE FROM users
	WHERE id = $1
`

func (r *PGRepo) DeleteById(id int64) error {
	r.logger.Debug("Executing query", zap.String("query", deleteByIdSql))
	_, err := r.db.Exec(context.Background(), deleteByIdSql, id)
	return err
}
