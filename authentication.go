package main
 
import (
    "fmt"
    "net/http"
 
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)
 
var (
    secretKey   = []byte("auth")
    cookieStore = sessions.NewCookieStore(secretKey)
    users       = map[string]string{"user": "pass", "admin": "adminpass"}
)
 
func main() {
    r := mux.NewRouter()
 
    // Define routes
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/login", LoginHandler).Methods("POST")
    r.HandleFunc("/dashboard", AuthMiddleware(DashboardHandler))
 
    http.Handle("/", r)
 
    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", nil)
}
 
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Home Page!")
}
 
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }
 
    // Get username and password from form data
    username := r.FormValue("username")
    password := r.FormValue("password")
 
    // Check if the provided credentials are valid
    if isValidCredentials(username, password) {
        // Authentication successful, create a session
        session, _ := cookieStore.Get(r, "session-name")
        session.Values["authenticated"] = true
        session.Values["username"] = username
        session.Save(r, w)
 
        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
        return
    }
    // Authentication failed
    fmt.Fprintln(w, "Invalid credentials. Please try again.")
}
 
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
    // Authorization check based on user role (replace with your own logic)
    session, _ := cookieStore.Get(r, "session-name")
    username, ok := session.Values["username"].(string)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
 
    // Authorized user, display the dashboard
    fmt.Fprintf(w, "Welcome to the Dashboard, %s!", username)
}
 
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Authorization check based on user session (replace with your own logic)
        session, _ := cookieStore.Get(r, "session-name")
        if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
 
        // Call the next handler
        next.ServeHTTP(w, r)
    }
}
 
func isValidCredentials(username, password string) bool {
    // Simulate user authentication (replace with your own logic)
    // In a real application, you would check credentials against a database
    storedPassword, ok := users[username]
    return ok && storedPassword == password
}