package main

import (
	"context"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/olivere/elastic/v7"	"io"
)

type Users struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Email    string `json:"email"`
}


func main() {
	// Create a client
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	defer client.Stop()

	ctx := context.Background()

	query := elastic.NewTermQuery("email", "fourth@gmail.com")
	get1, err := client.Search().
		Index("users").
		Query(query).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	for _, hit := range get1.Hits.Hits {
		var u Users
		err = json.Unmarshal(hit.Source, &u)
		if err != nil {
			panic(err)
		}
		fmt.Printf(u.User, u.Email, u.Password)
	}
}












//
//const (
//	host     = "localhost"
//	port     = 5432
//	userName = "postgres"
//	password = "Saharok#24"
//	dbname   = "postgres"
//)
//
//func main() {
//	r := gin.Default()
//	//r.GET("/ping", func(c *gin.Context) {
//	//	c.JSON(200, gin.H{
//	//		"message": "hello",
//	//	})
//	//})
//
//	r.POST("/user", func(c *gin.Context) {
//		var user Users
//		err := c.BindJSON(&user)
//		if err != nil {
//			log.Fatal(err)
//		}
//		if strings.TrimSpace(user.User) == "" {
//			c.JSON(http.StatusBadRequest, nil)
//			return
//		}
//		if len(user.Password) < 8 || len(user.Password) > 20 {
//			c.JSON(http.StatusBadRequest, nil)
//			return
//		}
//
//		matched, _ := regexp.MatchString(`[\w-]+@[\w]+(?:\.[\w]+)+`, user.Email)
//		if matched != true {
//			c.JSON(http.StatusBadRequest, "wrong email")
//			return
//		}
//
//		connStr := fmt.Sprintf(
//			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//			host,
//			port,
//			userName,
//			password,
//			dbname,
//		)
//		db, err := sql.Open("postgres", connStr)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, err)
//			return
//		}
//		err = db.Ping()
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, err)
//			return
//		}
//		log.Println("psql connected successfully")
//
//		query := `insert into "feamouser"("login","email", "password") values($1, $2, $3)`
//		_, err = db.Exec(
//			query,
//			user.User,
//			user.Email,
//			user.Password,
//		)
//		if err != nil {
//			return
//		}
//
//		c.JSON(http.StatusOK, user)
//
//	})
//
//	err := r.Run(":8080")
//	if err != nil {
//		return
//	} // () --- listen and serve on 0.0.0.0:8080- default
//}
