
// main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

type User struct {
    Id      string    `json:"Id"`
    First_Name      string    `json:"first_name"`
    Last_Name      string    `json:"last_name"`
    DOB string `json:"dob,omitempty"`
    Email string `json:"email"`
    Phone_num string `json:"phone"`
}

var Users []User

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllUsers")
    json.NewEncoder(w).Encode(Users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, user := range Users {
        if user.Id == key {
            json.NewEncoder(w).Encode(user)
        }
    }
}


func createNewUser(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new User struct
    // append this to our Users array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var user User 
    json.Unmarshal(reqBody, &user)
    // update our global Articles array to include
    // our new User
    Users = append(Users, user)

    json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, user := range Users {
        if user.Id == id {
            Users = append(Users[:index], Users[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/users", returnAllUsers)
    myRouter.HandleFunc("/user", createNewUser).Methods("POST")
    myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
    myRouter.HandleFunc("/user/{id}", returnSingleUser)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

    Users = []User{
        User{Id: "1",First_Name: "Ramu",Last_Name: "Kaka",DOB: "11/12/1992",Email: "ramu@kaka.com", Phone_num: "+915522501201"},
        User{Id: "2",First_Name: "Nattu",Last_Name: "Kaka",DOB: "11/12/1986",Email: "nattu@kaka.com", Phone_num: "+916622606206"},
    }
    handleRequests()
}