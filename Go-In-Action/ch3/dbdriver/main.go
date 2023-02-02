package main

import "database/sql"
import _ "LetUsGo/Go-In-Action/ch3/dbdriver/postgres"

func main() {
	_, err := sql.Open("postgres", "mydb")
	if err != nil {
		return
	}
}
