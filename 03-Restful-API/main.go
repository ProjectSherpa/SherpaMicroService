/* Go Restful API endpoint for Sherpa

PLACEHOLDER

*/

package main

//ftmt -i/o stuff
//html - ?
//github.com/gorilla/mux - Go route handler - used in our server for  wildcard handling

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var usersTest = `[{"userid":1,"first":"Quin","last":"Kinser","email":"quinkinser@gmail.com","username":"quink","lessonsCompleted":["lid1","lid2","lid3"]},{"userid":2,"first":"Wayne","last":"Adams","email":"quinkinser@gmail.com","username":"wayney","lessonsCompleted":["lid1","lid2"]},{"userid":3,"first":"Jeremy","last":"Toce","email":"quinkinser@gmail.com","username":"toasty","lessonsCompleted":["lid1"]}]`

type userStruct []struct {
	Userid int `json:"userid"`
	First string `json:"first"`
	Last string `json:"last"`
	Email string `json:"email"`
	Username string `json:"username"`
	LessonsCompleted []string `json:"lessonsCompleted"`
}
type PublicKey struct {
	Id  int
	Key string
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
  users := userStruct{} // creates a go object from userStruct
	json.Unmarshal([]byte(usersTest), &users) // transforms json(bytes) into the users go object

  // search for the user you want

  b, _ := json.Marshal(users[0]) // transform the user object into json bytes
  s := string(b) // transforms json bytes into string
  fmt.Fprint(w, "found me the first user ", s)
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
