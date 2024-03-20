package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
    "github.com/rs/cors"
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var db *sql.DB

func initDB() {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "123"
    dbName := "my_friends"
    var err error
    db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
}

func post(c web.C, w http.ResponseWriter, r *http.Request) {
    friend := c.URLParams["name"]
    _, err := db.Exec("INSERT INTO friends (name) VALUES (?)", friend)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "Added friend: %s", friend)
}

func get(c web.C, w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT name FROM friends")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var friends []string
    for rows.Next() {
        var name string
        err := rows.Scan(&name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        friends = append(friends, name)
    }

    json.NewEncoder(w).Encode(friends)
}

func deleteFriend(c web.C, w http.ResponseWriter, r *http.Request) {
    friend := c.URLParams["name"]
    _, err := db.Exec("DELETE FROM friends WHERE name=?", friend)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "Removed friend: %s", friend)
}

func main() {
    initDB()
    defer db.Close()

    // CORS setup
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "DELETE"},
    })

    // Routing
    goji.Use(c.Handler)
    goji.Get("/friend/", get)
    goji.Post("/friend/:name", post)
    goji.Delete("/friend/:name", deleteFriend)

    // Start server
    goji.Serve()
}


// package main

// import (
//   "fmt"
//   "io/ioutil"
//   "encoding/json"
//   "net/http"
//   "strings"
//   "github.com/rs/cors"
//   "github.com/zenazn/goji"
//   "github.com/zenazn/goji/web"
// )

// func addFriend(friend string) {
//   fileName := "data.txt"
//   contents,_ := ioutil.ReadFile(fileName)
//   result := string(contents) + "\n" + friend
//   //println(string(contents))
//   ioutil.WriteFile(fileName, []byte(result), 0x777)
// }

// func getFriends() ([]string)  {
//   fileName := "data.txt"
//   contents,_ := ioutil.ReadFile(fileName)
//   result := strings.Split(string(contents),"\n")
//   //fmt.Println(string(contents))
//   return result
// }

// func removeFriend(joke string){
//   jokes := getFriends();
//   //fmt.Printf("%q\n", jokes)
//   index := findFriend(joke, jokes)

//   if index != 999 {
//     jokes := append(jokes[:index],jokes[index+1:]...)
//     //fmt.Printf("%q\n", jokes)

//     fileName := "data.txt"
//     ioutil.WriteFile(fileName, []byte(strings.Join(jokes,"\n")), 0x777)
//   }
//   //fmt.Printf("%v\n", index)

// }

// func findFriend(a string, list []string) int {
//   for key, val := range list {
//     if val == a {
//       return key
//     }
//   }
//   return 999
// }

// func post(c web.C, w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "Add, %s!", c.URLParams["name"])
//   addFriend(c.URLParams["name"])
// }

// func get(c web.C, w http.ResponseWriter, r *http.Request) {
//   encoder := json.NewEncoder(w)
//   encoder.Encode(getFriends())
// }

// func delete(c web.C, w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "Remove, %s!", c.URLParams["name"])
//   removeFriend(c.URLParams["name"])
// }

// func main() {
//   //cross browser origin
//   c := cors.New(cors.Options{
//     AllowedOrigins: []string{"*"},
//     AllowedMethods: []string{"GET","POST","DELETE",},
//   })

//   //routing
//   goji.Use(c.Handler)
//   goji.Get("/friend/", get)
//   goji.Post("/friend/:name", post)
//   goji.Delete("/friend/:name", delete)
//   goji.Serve()
// }