/* Go Restful API endpoint for Sherpa

PLACEHOLDER

*/

package main

//ftmt -i/o stuff
//html - ?
//github.com/gorilla/mux - Go route handler - used in our server for  wildcard handling

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//mock data to test API functionality
var usersTest = `[{"userid":1,"first":"Quin","last":"Kinser","email":"quinkinser@gmail.com","username":"quink","lessonsCompleted":["lid1","lid2","lid3"]},{"userid":2,"first":"Wayne","last":"Adams","email":"quinkinser@gmail.com","username":"wayney","lessonsCompleted":["lid1","lid2"]},{"userid":3,"first":"Jeremy","last":"Toce","email":"quinkinser@gmail.com","username":"toasty","lessonsCompleted":["lid1"]}]`

//Array of structs representing our data
type userStruct []struct {
	Userid           int      `json:"userid"`
	First            string   `json:"first"`
	Last             string   `json:"last"`
	Email            string   `json:"email"`
	Username         string   `json:"username"`
	LessonsCompleted []string `json:"lessonsCompleted"`
}

// listAllUsers: Returns JSON object with multiple user information
// endpoint: /users
//Method: GET

func listAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, usersTest)
}

// createUpdateUser
// endpoint: /users
// Method: POST

//returnUser: Returns JSON object for single user specified by id number
// endpoint: /users/{userid}
// Method: GET

func returnUser(w http.ResponseWriter, r *http.Request) {
	users := userStruct{}                     // creates a go object from userStruct
	json.Unmarshal([]byte(usersTest), &users) // transforms json(bytes) into the users go object

	// search for the user you want

	b, _ := json.Marshal(users[0]) // transform the user object into json bytes
	s := string(b)                 // transforms json bytes into string
	addUser()
	fmt.Fprint(w, s)
}

func addUser() {
	db, err := sql.Open("mysql",
		"root:sherpa@tcp(127.0.0.1:3306)/sherpa")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	res, err := db.Query("INSERT into users (first, last, email, username, lessonsCompleted)values ('Wayne', 'Adams', 'wadams19@gmail.com', 'wadams', 'all')")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	// lastId, err := res.LastInsertId()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(lastID)
}

// Reading from MYSQL table and printing all rows to console
func readTable() {
	var (
		userid           int
		first            string
		last             string
		email            string
		username         string
		lessonsCompleted string
	)
	db, err := sql.Open("mysql",
		"root:sherpa@tcp(127.0.0.1:3306)/sherpa")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userid, &first, &last, &email, &username, &lessonsCompleted)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//Mysql

func mysqlTest() {
	db, err := sql.Open("mysql",
		"root:sherpa@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES(?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

}

// markLesson
// endpoint: /users/{userid}
// Method: POST

func main() {

	route := mux.NewRouter()

	route.HandleFunc("/users", listAllUsers)
	route.HandleFunc("/users/{userId}", returnUser)

	http.Handle("/", route)

	//Start server
	log.Fatal(http.ListenAndServe(":8081", route))

}
