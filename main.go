package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "84422664"
// 	dbname   = "message"
// )

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", indexPage)
	r.GET("/next", nextPage)
	r.POST("/submit", submitProcess)
	// r.GET("/show", conDB)

	r.Run()

}

// func conDB() {

// 	//connect

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	//insert

// 	stmt := `

// 	INSERT INTO message (message_input)
// 	VALUES ($1)

// 	`

// 	id := 0
// 	err = db.QueryRow(stmt, "some").Scan(&id)

// 	if err != nil {
// 		fmt.Println(err)

// 	}

// 	//delete (future feature)

// 	//query

// 	rows, err := db.Query("SELECT id_message, message_input FROM message ORDER BY id_message ASC")

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for rows.Next() {

// 		var id int
// 		var text string

// 		err = rows.Scan(&id, &text)
// 		fmt.Printf("id : %d Text %s \n", id, text)
// 	}
// }

//index.html

func indexPage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
	})

}

//next.html

func nextPage(c *gin.Context) {

	c.HTML(http.StatusOK, "next.html", gin.H{
		"title": "Second Page",
	})

}

func submitProcess(c *gin.Context) {

	// fmt.Println("Method", r.Method)

	inputValue := c.PostForm("text_input")
	formContent := c.PostForm("inputForm")

	c.JSON(http.StatusOK, gin.H{

		"status":  "Nice",
		"message": inputValue,
		"form":    formContent,
	})
}
