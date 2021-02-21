package condb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "84422664"
	dbname   = "message"
)

func conDB() {

	//declare

	// inputValue := c.PostForm("text_input")
	// formContent := c.PostForm("inputForm")

	//connect
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert

	stmt := `

	INSERT INTO message (message_input)
	VALUES ($1)

	`

	id := 0
	err = db.QueryRow(stmt, "some").Scan(&id)

	if err != nil {
		fmt.Println(err)

	}

	//delete (future feature)

	//query

	rows, err := db.Query("SELECT id_message, message_input FROM message ORDER BY id_message ASC")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {

		var id int
		var text string

		err = rows.Scan(&id, &text)
		fmt.Printf("id : %d Text : %s \n", id, text)
	}
}
