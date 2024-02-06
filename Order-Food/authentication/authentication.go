package authentication

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const usersFile = "users.csv"

func Init() {
	// Create users file if it doesn't exist
	file, err := os.OpenFile(usersFile, os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

// signupUser adds a new user to the file if the username is not already taken.
func SignupUser(username, password string) (string, error) {
	// Check if the user already exists
	users, err := ReadUsersFile()
	if err != nil {
		return "", err
	}

	for _, user := range users {
		if user[0] == username {
			return "Username already exists", nil
		}
	}

	// Add the new user to the file
	err = WriteUserToFile(username, password)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("User %s created successfully!", username), nil
}

// loginUser checks if the provided username and password match any existing user.
func LoginUser(username, password string) (string, error) {
	// Check if the user exists and the password matches
	users, err := ReadUsersFile()
	if err != nil {
		return "", err
	}

	for _, user := range users {
		if user[0] == username && user[1] == password {
			return fmt.Sprintf("Welcome, %s!", username), nil
		}
	}

	return "Invalid username or password", nil
}

func ReadUsersFile() ([][]string, error) {
	file, err := os.Open(usersFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	users, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func WriteUserToFile(username, password string) error {
	file, err := os.OpenFile(usersFile, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{username, password})
	if err != nil {
		return err
	}

	return nil
}