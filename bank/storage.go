package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(100),
		last_name varchar(100),
		number serial,
		
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}
func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into account (
			first_name , last_name , number , balance , created_at
		) values (
			$1 , $2 , $3 , $4 , $5
		)`

	_, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	query := `select * from account`
	res, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for res.Next() {
		account, err := scanIntoAccounts(res)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)

	}
	return accounts, nil

}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query(`delete from account where id=$1`, id)
	return err

}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) GetAccountById(id int) (*Account, error) {

	query := `select * from account where id=$1`
	res, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		return scanIntoAccounts(res)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func scanIntoAccounts(rows *sql.Rows) (*Account, error) {
	account := Account{}
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	return &account, err
}
