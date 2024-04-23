package config

import "database/sql"

// ตั้งค่า database
func DatabaseConnection() *sql.DB {
	database, _ := sql.Open("sqlite3", "./pongsawn.db")
	return database
}
