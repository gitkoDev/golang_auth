package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	db *pgx.Conn
}

func NewStorage(db *pgx.Conn) *Storage {
	return &Storage{db: db}
}

func (s *Storage) IsExistingUser(c *gin.Context, uuid string) bool {
	var id int

	query := `SELECT 1 FROM users where id=$1`
	err := s.db.QueryRow(c, query, uuid).Scan(&id)
	return err == nil
}

func (s *Storage) IsExistingHash(c *gin.Context, hash []byte) bool {
	var refresh_token_hash string

	query := `SELECT 1 FROM users where refresh_token_hash=$1`
	err := s.db.QueryRow(c, query, hash).Scan(&refresh_token_hash)

	return err != pgx.ErrNoRows
}

func (s *Storage) IsExistingIp(c *gin.Context, user_ip string) bool {
	var ip string

	query := `SELECT 1 FROM users where ip=$1`
	err := s.db.QueryRow(c, query, user_ip).Scan(&ip)

	return err != pgx.ErrNoRows
}

func (s *Storage) AddUser(c *gin.Context, user_id string, user_ip string, refresh_token_hash []byte) {
	query := `INSERT INTO users(id, ip, refresh_token_hash) VALUES($1, $2, $3)`
	_, err := s.db.Exec(c, query, user_id, user_ip, refresh_token_hash)
	if err != nil {
		logrus.Fatalln("error adding user to database", err)
	}
}

func (s *Storage) UpdateRefreshTokenHash(c *gin.Context, uuid string, new_hash []byte) {
	query := `UPDATE users 
		SET refresh_token_hash = $1
		WHERE id = $2
	`
	_, err := s.db.Exec(c, query, new_hash, uuid)
	if err != nil {
		logrus.Fatalln("error updating refresh token hash")
	}
}
