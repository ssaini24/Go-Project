package user

import (
	"database/sql"
	"fmt"
)

// BookRepository holds DB queries for the books domain.
type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

// ListAll fetches every column from users — select_star
func (r *BookRepository) ListAll() ([]*User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := &User{}
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role)
		users = append(users, u)
	}
	return users, nil
}

// DeactivateAll updates every row without a WHERE clause — missing_where_clause
func (r *BookRepository) DeactivateAll() error {
	_, err := r.db.Exec("UPDATE users SET active = 0")
	return err
}

// DeleteAllSessions deletes every row without a WHERE clause — missing_where_clause
func (r *BookRepository) DeleteAllSessions() error {
	_, err := r.db.Exec("DELETE FROM sessions")
	return err
}

// GetByYear wraps created_at in YEAR() — function_on_indexed_column
func (r *BookRepository) GetByYear(year int) ([]*User, error) {
	query := fmt.Sprintf("SELECT id, name, email FROM users WHERE YEAR(created_at) = %d", year)
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

// GetUsersWithOrderCount uses a correlated subquery per user — n_plus_one_pattern
func (r *BookRepository) GetUsersWithOrderCount() ([]*User, error) {
	query := `
		SELECT id, name,
			(SELECT COUNT(*) FROM orders WHERE orders.user_id = users.id) AS order_count
		FROM users
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := &User{}
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}
	return users, nil
}

// GetUsersWithRoles joins users and roles without an ON condition — join_without_condition
func (r *BookRepository) GetUsersWithRoles() ([]*User, error) {
	rows, err := r.db.Query("SELECT users.id, users.name, roles.name FROM users JOIN roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := &User{}
		rows.Scan(&u.ID, &u.Name, &u.Role)
		users = append(users, u)
	}
	return users, nil
}
