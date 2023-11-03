package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/emanuel-xavier/hexagonal-architerure/configs"
	_ "github.com/lib/pq"
)

var (
	singletonDB *sql.DB
	once        sync.Once
	err         error
)

// GetConnection initializes and returns a connection to the PostgreSQL database.
func GetConnection() (*sql.DB, error) {
	// Try to initialize the connection only once
	once.Do(initializeConnection)

	if singletonDB == nil {
		return nil, fmt.Errorf("failed to initialize the database connection")
	}

	return singletonDB, nil
}

func Close() {
	singletonDB.Close()
}

func initializeConnection() {
	conf := configs.GetDB()

	connStr := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=%s",
		conf.Host, conf.User, conf.Database, conf.Port, conf.Pass, conf.SSLMode)

	// Attempt connection initialization with retries
	var err error
	for retries := 3; retries > 0; retries-- {
		singletonDB, err = sql.Open("postgres", connStr)
		if err != nil {
			// Sleep for a moment before the next retry
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	if err == nil {
		// Successfully initialized the connection
		singletonDB.SetConnMaxIdleTime(time.Duration(conf.MaxIdleConnections))
		singletonDB.SetMaxOpenConns(conf.MaxOpenConnections)

		// Check the connection by pinging the database
		if err := singletonDB.Ping(); err != nil {
			// If the ping fails, close the connection and set it to nil
			singletonDB.Close()
			singletonDB = nil
			log.Fatal(err)
			log.Fatal("Could not ping postgres")
		}
	} else {
		log.Fatal("Database initialization failed")
	}
}

func ReleaseConnection() {
	if singletonDB != nil {
		singletonDB.Close()
		singletonDB = nil
	}
}
