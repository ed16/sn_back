package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// The model.
type User struct {
	gorm.Model
	Name string `json:"name"`
}

func main() {
	// Open a new connection to our sqlite database.
	db, err := gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("Failed to open the SQLite database.")
	}
	// Create the table from our struct.
	db.AutoMigrate(&User{})

	// Create a new user in our database.
	db.Create(&User{
		Name: "Craig",
	})

	// Find all of our users.
	var users []User
	db.Find(&users)

	// Output the users from the DB json encoded
	jsonEncoded, _ := json.Marshal(&users)
	fmt.Println(string(jsonEncoded))

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("index.html", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"userName": "Daniel Fridman",
		})
	})

	r.POST("/sendMessage", func(c *gin.Context) {
		x, _ := ioutil.ReadAll(c.Request.Body)

		c.JSON(200, gin.H{
			"message": string(x),
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
