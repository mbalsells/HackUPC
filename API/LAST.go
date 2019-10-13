package main

import (
	// "fmt"
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
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

// this was login and register

type clock struct {
	hour int
	min  int
}

type lesson struct {
	subject    string
	start_time clock
	end_time   clock
	feedback   [11]int
}

type almost_lesson struct {
	subject    string
	start_time clock
	end_time   clock
}

func Newalmost_Lesson(l lesson) almost_lesson {
	_l := almost_lesson{l.subject, l.start_time, l.end_time}
	return _l
}

func min_lesson(l1, l2 lesson) lesson {
	if leq(l1.start_time, l2.start_time) {
		return l1
	}
	return l2
}

func NewLesson(subject string, start_time, end_time clock) lesson {
	var _empty [11]int
	l := lesson{subject, start_time, end_time, _empty}
	return l
}

func NewClock(hour, min int) clock {
	t := clock{hour, min}
	return t
}

func less(a, b clock) bool {
	return (a.hour < b.hour) || (a.hour == b.hour && a.min < b.min)
}

func greater(a, b clock) bool {
	return less(b, a)
}

func leq(a, b clock) bool {
	return less(a, b) || a == b
}

func geq(a, b clock) bool {
	return leq(b, a)
}

type subject struct {
	name     string
	teacher  string
	schedule [7]lesson
}

func NewSubject(name, teacher string, schedule [7]lesson) subject {
	s := subject{name, teacher, schedule}
	return s
}

type user struct {
	username string
	email    string
	name     string
	subjects []string
	feedback []int
}

func NewUser(username, name, email string, subjects []string) user {
	var _zeroes []int
	for _, _ = range subjects {
		_zeroes = append(_zeroes, 0)
	}

	u := user{username, email, name, subjects, _zeroes}
	return u
}

var map_subjects map[string]subject
var map_users map[string]user

func performset_feedback(u user, pointsstring string, subject_name string) {
	points, _ := strconv.Atoi(pointsstring)

	today := (int(time.Now().Weekday()) + 6) % 7
	//today = 1 EOOOOOO MORE INTERESTING

	for i, x := range u.subjects {
		_subj := map_subjects[x]
		//fmt.Println("subj, ", _subj)
		if _subj.schedule[today].start_time.hour > 0 {
			//fmt.Println("now u , ", u)

			if x == subject_name {
				u.feedback[i] = points
				for j, y := range map_subjects {
					if y.name == subject_name {
						temp := map_subjects[j]
						temp.schedule[today].feedback[points]++
						map_subjects[j] = temp
					}
				}
			}
		}
	}
}

func getInfoUser(u string) (map[string][]string, bool) {
	_user, err := map_users[u]

	var m map[string][]string
	m = make(map[string][]string)

	m["username"] = []string{_user.username}
	m["name"] = []string{_user.name}
	m["email"] = []string{_user.email}
	m["subject"] = _user.subjects

	var feed []string

	for _, v := range _user.feedback {
		feed = append(feed, strconv.Itoa(v))
	}

	m["feedback"] = feed

	return m, err
}

func performSchedule(u user) []almost_lesson {
	var ans []almost_lesson
	today := (int(time.Now().Weekday()) + 6) % 7
	for _, name := range u.subjects {
		_subject := map_subjects[name]

		x := _subject.schedule[today]
		if x.start_time.hour > 0 {
			ans = append(ans, Newalmost_Lesson(x))
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		x := ans[i].start_time
		y := ans[j].start_time
		return less(x, y)
	})

	return ans
}

func setfeedback(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	performset_feedback(map_users[req.Form["username"][0]], req.Form["point"][0], req.Form["subjectName"][0])

	var m map[string]string
	m = make(map[string]string)

	m["result"] = "true"
	value, _ := json.Marshal(m)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("content-type", "application/json")
	w.Write(value)
}

func schedule(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	result := performSchedule(map_users[req.Form["username"][0]])

	var _subject []string
	var _start_time []string
	var _end_time []string

	for _, val := range result {
		_subject = append(_subject, val.subject)
		_start_time = append(_start_time, strconv.Itoa(val.start_time.hour)+":"+strconv.Itoa(val.start_time.min))
		_end_time = append(_end_time, strconv.Itoa(val.end_time.hour)+":"+strconv.Itoa(val.end_time.min))
	}

	var m map[string][]string
	m = make(map[string][]string)

	m["subject"] = _subject
	m["start_time"] = _start_time
	m["end_time"] = _end_time

	value, _ := json.Marshal(m)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("content-type", "application/json")
	w.Write(value)
}

func getaverage(subject_name string) float64 {
	count := 0
	sum := 0

	for _, v := range map_users {
		for i, sub := range v.subjects {
			if sub == subject_name && v.feedback[i] > 0 {
				count++
				sum += v.feedback[i]
			}
		}
	}

	den := float64(count)
	num := float64(sum)

	if count == 0 {
		return 0
	} else {
		return num / den
	}
}

func average(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	result := getaverage(req.Form["subjectName"][0])

	var m map[string]string
	m = make(map[string]string)

	m["result"] = strconv.FormatFloat(result, 'f', 6, 64)

	value, _ := json.Marshal(m)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("content-type", "application/json")
	w.Write(value)
}

func infouser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	result, _ := getInfoUser(req.Form["username"][0])
	value, _ := json.Marshal(result)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("content-type", "application/json")
	w.Write(value)
}

// SetupHandlers initiates the servers HTTP endpoints
func SetupHandlers() {
	http.HandleFunc("/setfeedback", setfeedback)
	http.HandleFunc("/schedule", schedule)
	http.HandleFunc("/infouser", infouser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/average", average)
	http.ListenAndServe(":8080", nil)
}

func init_everything() error {
	map_subjects = make(map[string]subject)
	map_users = make(map[string]user)

	filesUsernames, err := os.OpenFile("data/usernames.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesUsernames.Close()

	filesNames, err := os.OpenFile("data/names.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesNames.Close()

	filesEmails, err := os.OpenFile("data/emails.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesEmails.Close()

	ScannerUsernames := bufio.NewScanner(filesUsernames)
	ScannerNames := bufio.NewScanner(filesNames)
	ScannerEmails := bufio.NewScanner(filesEmails)

	for ScannerUsernames.Scan() { // cheks that the username is not taken
		ScannerNames.Scan()
		ScannerEmails.Scan()

		_username := ScannerUsernames.Text()
		_name := ScannerNames.Text()
		_email := ScannerEmails.Text()

		var _empty []string
		map_users[_username] = NewUser(_username, _name, _email, _empty)
	}

	filesAssignments, err := os.OpenFile("data/Assignments.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesAssignments.Close()

	ScannerAssignments := bufio.NewScanner(filesAssignments)

	for ScannerAssignments.Scan() { // cheks that the username is not taken
		_username := ScannerAssignments.Text()
		ScannerAssignments.Scan()
		_assignment := ScannerAssignments.Text()

		temp_user := map_users[_username] //EOOO
		temp_user.subjects = append(temp_user.subjects, _assignment)
		temp_user.feedback = append(temp_user.feedback, 0)
		map_users[_username] = temp_user
	}

	filesSubject, err := os.OpenFile("data/subject.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesSubject.Close()

	ScannerSubject := bufio.NewScanner(filesSubject)

	for ScannerSubject.Scan() { // cheks that the username is not taken
		_assignment := ScannerSubject.Text()
		ScannerSubject.Scan()
		_teacher := ScannerSubject.Text()

		var _empty [7]lesson
		map_subjects[_assignment] = NewSubject(_assignment, _teacher, _empty)
	}

	filesSchedules, err := os.OpenFile("data/schedule.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesSchedules.Close()

	ScannerSchedules := bufio.NewScanner(filesSchedules)

	for ScannerSchedules.Scan() { // cheks that the username is not taken
		_assignment := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_day := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_starth := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_startm := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_endh := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_endm := ScannerSchedules.Text()

		temp_subject := map_subjects[_assignment]

		int_starth, _ := strconv.Atoi(_starth)
		int_startm, _ := strconv.Atoi(_startm)
		int_endh, _ := strconv.Atoi(_endh)
		int_endm, _ := strconv.Atoi(_endm)
		int_day, _ := strconv.Atoi(_day)

		tmp_less := NewLesson(_assignment, NewClock(int_starth, int_startm), NewClock(int_endh, int_endm))

		temp_subject.schedule[int_day] = tmp_less
		map_subjects[_assignment] = temp_subject
	}

	return errors.New("Everything okay :)")
}

func main() {
	init_everything()
	SetupHandlers()
}
