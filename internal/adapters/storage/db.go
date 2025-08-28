package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
    
	_ "github.com/lib/pq"
)

type DbEnvs struct {
    db_host string
    db_port string 
    db_user string
    db_password string
    db_name string
    db_ssl string
}

func OpenConnection() (*sql.DB, error) {
    envs, err := getDbEnvs()
    if err != nil {
        return nil, err
    }

	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", 
        envs.db_host, envs.db_port, envs.db_user, envs.db_password, envs.db_name, envs.db_ssl,
    )


	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func getDbEnvs() (*DbEnvs, error) { 
    db_host := os.Getenv("DB_HOST")
    if db_host == "" {
        return nil, errors.New("db_host is not set")
    }  
    db_port := os.Getenv("DB_PORT")
    if db_port == "" {
        return nil, errors.New("db_port is not set")
    } 
    db_user := os.Getenv("DB_USER")
    if db_user == "" {
        return nil, errors.New("db_user is not set")
    } 
    db_password := os.Getenv("DB_PASSWORD")
    if db_password == "" {
        return nil, errors.New("db_password is not set")
    } 
    db_name := os.Getenv("DB_NAME")
    if db_name == "" {
        return nil, errors.New("db_name is not set")
    } 
    db_ssl := os.Getenv("DB_SSL")
    if db_ssl == "" {
        return nil, errors.New("db_ssl is not set")
    } 
    envs := DbEnvs{
        db_host,
        db_port,
        db_user,
        db_password,
        db_name,
        db_ssl,
    }
    
    return &envs, nil
}

