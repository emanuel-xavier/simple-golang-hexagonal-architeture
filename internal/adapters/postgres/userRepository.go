package adapters

import (
	"log"

	"github.com/emanuel-xavier/hexagonal-architerure/internal/core/domain"
	"github.com/emanuel-xavier/hexagonal-architerure/internal/core/port"
	"github.com/emanuel-xavier/hexagonal-architerure/internal/db/postgres"
)

func NewPostgresUserRepository() (psqlRepo port.UserRepository, err error) {
	_, err = postgres.GetConnection()
	if err != nil {
		panic(err)
	}

	psqlRepo = PostgresRespository{}
	return
}

type PostgresRespository struct{}

func (pr PostgresRespository) Save(user domain.User) error {
	conn, err := postgres.GetConnection()
	if err != nil {
		log.Println("GET CONNECTION", err)
		return err
	}

	sql := `INSERT INTO users (username, id) VALUES ($1, $2)`
	_, err = conn.Exec(sql, user.Username, user.Id)

	if err != nil {
		log.Println("INSERT", err)
	}

	return err
}

func (pr PostgresRespository) GetUserById(id string) (user *domain.User, err error) {
	conn, err := postgres.GetConnection()
	if err != nil {
		log.Println("GET CONNECTION", err)
		return
	}

	sql := `SELECT username, id FROM users WHERE id = $1` // Fix the SQL query

	user = &domain.User{}

	err = conn.QueryRow(sql, id).Scan(&user.Username, &user.Id)
	if err != nil {
		log.Println("SELECT", err)
	}

	return
}

func (pr PostgresRespository) GetAll() (users []domain.User, err error) {
	conn, err := postgres.GetConnection()
	if err != nil {
		log.Println("GET CONNECTION", err)
		return
	}

	rows, err := conn.Query("SELECT * FROM users")
	if err != nil {
		log.Println("SELECT", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.Username, &user.Id)
		if err != nil {
			log.Println("SCAN ROW", err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
