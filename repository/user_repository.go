package repository

import (
	"database/sql"
	"fmt"

	"github.com/sidz111/user-management-crud/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create
func (r *UserRepository) Create(user model.User) error {
	query := "Insert into student(name, address) values(?,?)"
	_, err := r.db.Exec(query, user.Name, user.Address)
	if err != nil {
		return fmt.Errorf("Failed to create User %w", err)
	}
	return nil
}

// get by id
func (r *UserRepository) GetByid(id int) (*model.User, error) {
	query := "select id, name, address from user where id=?"
	user := &model.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found")
		}
		return nil, fmt.Errorf("Failed to get User %w", err)
	}
	return user, nil
}

// delete by id
func (r *UserRepository) DeleteById(id int) error {
	query := "delete from usser where id =?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete user %w", err)
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("User not found")
	}
	return nil
}

// Get all Users
func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User

	query := "Select id, name, address from user"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Users Not Found %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Address)
		if err != nil {
			return nil, fmt.Errorf("Failed to Scan User")
		}
		users = append(users, user)
	}
	return users, nil
}

// Update user
func (r *UserRepository) UpdateUser(id int, user *model.User) error {
	query := "Update user set name=? address =? where id =?"
	result, err := r.db.Exec(query, user.Name, user.Address, user.ID)
	if err != nil {
		return fmt.Errorf("Fail to Update User %w", err)
	}

	row, _ := result.RowsAffected()
	if row == 0 {
		return fmt.Errorf("User Not Found")
	}
	return nil
}
