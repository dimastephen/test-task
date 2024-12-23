package auth

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	interf "test-task/internal/repository"
	"test-task/internal/repository/model"
	model2 "test-task/internal/service/model"
	"time"
)

type repo struct {
	db *pgxpool.Pool
}

func (r *repo) Create(ctx context.Context, refreshToken string) (int, error) {
	ip := ctx.Value("ip")
	guid := ctx.Value("guid")
	expires_at := time.Now().Add(12 * time.Hour)
	var id int

	hashedToken, err := hashRefreshToken(refreshToken)
	if err != nil {
		return 0, err
	}

	builder := sq.Insert("refresh").Columns("user_ip", "user_id", "token_hash", "expires_at").
		Values(ip, guid, hashedToken, expires_at).PlaceholderFormat(sq.Dollar).Suffix("RETURNING ID")
	query, _, err := builder.ToSql()
	err = r.db.QueryRow(ctx, query, ip, guid, hashedToken, expires_at).Scan(&id)
	if err != nil {
		return 0, err
	}
	log.Printf("Inserted new info about user with guid=%v, ip=%v, returning id=%v", guid, ip, id)
	return id, nil
}

func (r *repo) Get(ctx context.Context, info model.RefreshTokenInfo) (*model.RefreshTokenInfo, error) {
	var response model.RefreshTokenInfo
	guid, err := strconv.Atoi(info.Guid)
	if err != nil {
		return nil, errors.New("failed to parse guid")
	}
	log.Printf("%v - guid", guid)

	query := "SELECT user_ip, token_hash FROM refresh WHERE user_id=$1"
	log.Printf("Executing query: %s with args: %v", query, guid)
	row := r.db.QueryRow(ctx, query, guid)
	err = row.Scan(&response.Ip, &response.RefreshToken)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, errors.New("No user with guid match")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(response.RefreshToken), []byte(info.RefreshToken))
	if err != nil {
		return nil, errors.New("Wrong refresh token")
	}

	response.Guid = info.Guid
	return &response, nil
}

func (r *repo) Update(ctx context.Context, update model2.RefreshUpdate) (id int, err error) {
	hashedRefresh, err := bcrypt.GenerateFromPassword([]byte(update.NewRefreshToken), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	query := "UPDATE refresh SET user_ip=$1, token_hash=$2, expires_at=$3, updated_at=$4 WHERE user_id=$5"
	rows, err := r.db.Exec(ctx, query, update.Ip, hashedRefresh, time.Now().Add(12*time.Hour), time.Now(), update.Guid)
	ids := rows.RowsAffected()
	return int(ids), nil
}

func NewAuthRepository(db *pgxpool.Pool) interf.AuthRepository {
	return &repo{db: db}
}
