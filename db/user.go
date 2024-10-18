package db

import (
	"github.com/sanchae/library-test/models"
)

func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}
	rows, err := db.Conn.Query(`
	SELECT * FROM users
	`)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) AddUser(user *models.User) error {
    query := `
    INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING first_name, last_name
    `
    err := db.Conn.QueryRow(query, user.FirstName, user.LastName).Scan(&user.FirstName, &user.LastName)
    if err != nil {
        return err
    }
    return nil
}