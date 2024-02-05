package main
 
import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
 
    _ "github.com/go-sql-driver/mysql"
)
 
const (
    dbDriver   = "mysql"
    dbUser     = "root"
    dbPassword = "123"
    dbName     = "authentication_db"
)
 
var db *sql.DB
 
func main() {
    // Connect to the database
    var err error
    db, err = sql.Open(dbDriver, fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName))
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	//signup page
	http.HandleFunc("/register", registerHandler)

    // Serve login page and handle login requests
    http.HandleFunc("/login", loginHandler)
 
    // Start server
    fmt.Println("Server started on :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }

    // Get username and password from form
    username := r.Form.Get("username")
    password := r.Form.Get("password")

	//Check if both username and password are provided
	if username == "" || password == "" {
        http.Error(w, "Username and password are required", http.StatusBadRequest)
        return
    }

    // Check if the user already exists
    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        log.Println("Error querying user:", err)
        return
    }

    if count > 0 {
        http.Error(w, "Username already exists", http.StatusConflict)
        return
    }

    // Insert the new user into the database
    _, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        log.Println("Error creating user:", err)
        return
    }

    // Registration successful
    fmt.Fprintf(w, "User %s created successfully!", username)
}
 
func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
 
    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }
 
    // Get username and password from form
    username := r.Form.Get("username")
    password := r.Form.Get("password")
 
    // Query the database for the user
    var storedPassword string
    err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Database error", http.StatusInternalServerError)
        log.Println("Error querying user:", err)
        return
    }
 
    // Compare stored password with the provided password
    if storedPassword != password {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }
 
    // Authentication successful
    fmt.Fprintf(w, "Welcome, %s!", username)
}
 