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
	// "strings"
)

// var usersTest = `{"users":[{"userid":1,"first":"Quin","last":"Kinser","email":"quinkinser@gmail.com","username":"quink","lessonsCompleted":["lid1","lid2","lid3"]},{"userid":2,"first":"Wayne","last":"Adams","email":"quinkinser@gmail.com","username":"wayney","lessonsCompleted":["lid1","lid2"]},{"userid":3,"first":"Jeremy","last":"Toce","email":"quinkinser@gmail.com","username":"toasty","lessonsCompleted":["lid1"]}]}`

var usersTest = `[{"userid":1,"first":"Quin","last":"Kinser","email":"quinkinser@gmail.com","username":"quink","lessonsCompleted":["lid1","lid2","lid3"]},{"userid":2,"first":"Wayne","last":"Adams","email":"quinkinser@gmail.com","username":"wayney","lessonsCompleted":["lid1","lid2"]},{"userid":3,"first":"Jeremy","last":"Toce","email":"quinkinser@gmail.com","username":"toasty","lessonsCompleted":["lid1"]}]`

type userStruct struct {
	userid           int
	first            string
	last             string
	email            string
	username         string
	lessonsCompleted []string
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
	// urlPart := strings.Split(r.URL.Path, "/")
	// bytes := []byte(usersTest)

	// // parseTest, _ := json.Marshal(usersTest)
	// users := make([]userStruct, 0)
	// json.Unmarshal(bytes, &users)

	// var first = users[:1]

	// b, _ := json.Marshal(first)
	// s := string(b)

	// // fmt.Printf(users)
	// fmt.Fprint(w, "found me the first user ", s)
	keysBody := []byte(`[{"id": 1,"key": "-"},{"id": 2,"key": "-"},{"id": 3,"key": "-"}]`)
	keys := make([]PublicKey, 0)
	json.Unmarshal(keysBody, &keys)
	fmt.Printf("%#v", keys)

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
