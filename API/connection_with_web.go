package main

import (
    "fmt"
    //"time"
    //"sort"
    "bufio"
	//"encoding/json"
	"errors"
	//"net/http"
	"os"
	"strconv"

)

type clock struct {
    hour int
    min int
}

type lesson struct {
    subject         string
    start_time      clock
    end_time        clock
}

func min_lesson (l1, l2 lesson) lesson {
    if leq(l1.start_time, l2.start_time) {
        return l1
    }
    return l2
}

func NewLesson (subject string, start_time, end_time clock) lesson {
    l := lesson {subject, start_time, end_time}
    return l
}

func NewClock(hour, min int) clock {
    t := clock {hour, min}
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
    name string
    teacher string
    schedule [7]lesson
}

func NewSubject (name, teacher string, schedule [7] lesson) subject {
    s := subject {name, teacher, schedule}
    return s
}


type user struct {
    username string
    email string
    name string
    subjects [] string
}

func NewUser (username, name, email string, subjects []string) user {
    u := user {username, name, email, subjects}
    return u
}


var map_subjects map[string]subject
var map_users map[string]user

/*
func getInfoUser (username string) (map[string] string, error) {
	v := map[string]string{"hello": "world"}
	return v
}

*//*
func daily_schedule(u user) []lesson {
    var ans []lesson
    today := int(time.Now().Weekday()) - 1
    for _, name := range u.subjects {
        subject := map_subjects[name]
        x := subject.schedule[today]
        if x.start_time.hour != -1 {
            ans = append(ans, x)
        }
    }
    sort.Slice(ans, func(i, j int) bool {
        x := ans[i].start_time
        y := ans[j].start_time
        return less(x, y)
        })
    return ans
}
*/
/*
func schedule(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	result, err := performSchedule(req.Form["username"][0])

	value, _ := json.Marshal(map[string]interface{}{
		"": result,
		"error":   err.Error(),
	})

	w.Header().Set("content-type", "application/json")
	w.Write(value)
}*/

/*
func infouser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	result, err := getInfoUser(req.Form["username"][0])

	jsonString, _ := json.Marshal(datas)

	value, _ := json.Marshal(map[string]interface{}{
		"": value,
		"error":   err.Error(),
	})

	w.Header().Set("content-type", "application/json")
	w.Write(value)
}


// SetupHandlers initiates the servers HTTP endpoints
func SetupHandlers() {
	//http.HandleFunc("/schedule", login)
	http.HandleFunc("/infouser", register)
	http.ListenAndServe(":8080", nil)
}
*/



func init_everything () error {
	map_subjects = make (map[string]subject)
	map_users = make (map[string]user)

	filesUsernames, err := os.OpenFile("usernames.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesUsernames.Close()

	filesNames, err := os.OpenFile("names.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesNames.Close()

	filesEmails, err := os.OpenFile("emails.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	
	filesAssignments, err := os.OpenFile("Assignments.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
		map_users[_username] = temp_user
	}

	filesSubject, err := os.OpenFile("subject.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	
	filesSchedules, err := os.OpenFile("schedule.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

