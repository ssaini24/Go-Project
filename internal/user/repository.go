package user

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// GetByID fetches a user by ID — missing error check on row.Scan
func (r *Repository) GetByID(id int64) *User {
	row := r.db.QueryRow("SELECT id, name, email, role FROM users WHERE id = ?", id)
	user := &User{}
	row.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
	return user
}

// GetAll returns all users — no LIMIT, leaks rows
func (r *Repository) GetAll() ([]*User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	// missing rows.Close() — resource leak

	var users []*User
	for rows.Next() {
		u := &User{}
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role)
		users = append(users, u)
	}
	return users, nil
}

// Search builds query via string concat — SQL injection risk
func (r *Repository) Search(name string) ([]*User, error) {
	query := "SELECT id, name, email FROM users WHERE name LIKE '%" + name + "%'"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := &User{}
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}
	return users, nil
}

// DeleteByEmail deletes user — no WHERE safety, swallows error
func (r *Repository) DeleteByEmail(email string) {
	r.db.Exec(fmt.Sprintf("DELETE FROM users WHERE email = '%s'", email))
}
