package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

// given a file (named fileName) and a text, inserts the text in the file with end of line
func add(fileName, text string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	if _, err := file.WriteString(text + "\n"); err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func performRegister(name, email, username, password string) (bool, error) {
	fileUsernames, err := os.OpenFile("data/usernames.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false, err
	}
	defer fileUsernames.Close()

	scanner := bufio.NewScanner(fileUsernames)
	for scanner.Scan() { // cheks that the username is not taken
		if scanner.Text() == username {
			return false, errors.New("username is already being used")
		}
	}

	//add all the information
	err = add("data/usernames.txt", username)
	if err != nil {
		return false, err
	}
	err = add("data/names.txt", name)
	if err != nil {
		return false, err
	}
	err = add("data/passwords.txt", password)
	if err != nil {
		return false, err
	}
	err = add("data/emails.txt", email)
	if err != nil {
		return false, err
	}

	return true, errors.New("Everything okay :)")
}

func performLogin(username, password string) (bool, error) {
	fileUsernames, err := os.OpenFile("data/usernames.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false, err
	}
	defer fileUsernames.Close()

	filePasswords, err := os.OpenFile("data/passwords.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false, err
	}
	defer filePasswords.Close()

	// cheks that the username exists and that its password is correct
	scannerUsernames := bufio.NewScanner(fileUsernames)
	scannerPasswords := bufio.NewScanner(filePasswords)

	for scannerUsernames.Scan() {
		scannerPasswords.Scan()
		key := scannerPasswords.Text()

		if scannerUsernames.Text() == username {
			if key == password {
				return true, errors.New("Everything okay :)")
			} else {
				return false, errors.New("Incorrect password")
			}
		}
	}

	return false, errors.New("user not registered")
}

func login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	succ, err := performLogin(req.Form["username"][0], req.Form["password"][0])
	value, _ := json.Marshal(map[string]interface{}{
		"success": succ,
		"error":   err.Error(),
	})

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("content-type", "application/json")
	w.Write(value)
}

func register(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	succ, err := performRegister(req.Form["name"][0], req.Form["email"][0], req.Form["username"][0], req.Form["password"][0])

	value, _ := json.Marshal(map[string]interface{}{
		"success": succ,
		"error":   err.Error(),
	})

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("content-type", "application/json")
	w.Write(value)
}

// SetupHandlers initiates the servers HTTP endpoints
func SetupHandlers() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.ListenAndServe(":8080", nil)
}

func main() {
	SetupHandlers()
}
