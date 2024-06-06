package dal

import (
	"database/sql"
	"tutorials/dynamic_db_conn_project/models"
)

func RegisterUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func GetUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(db *sql.DB, userID string, user models.User) error {
	_, err := db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, userID)
	return err
}

func DeleteUser(db *sql.DB, userID string) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", userID)
	return err
}