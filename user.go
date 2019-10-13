package main

import (
  "fmt"
  "time"
)

type lesson struct {
  subject     string
  start_time  clock
  end_time    clock
}

func NewLesson (subject string, start_time, end_time clock) lesson {
  l := lesson {subject, start_time, end_time}
  return l
}

type clock struct {
  hour int
  min int
}

func NewClock(hour, min int) clock {
  t := clock {hour, min}
  return t
}

func less(a, b *clock) bool {
  return (a.hour < b.hour) || (a.hour == b.hour && a.min < b.min)
}

func greater(a, b *clock) bool {
  return less(b, a)
}

func leq(a, b *clock) bool {
  return less(a, b) || a == b
}

func geq(a, b *clock) bool {
  return leq(b, a)
}

type subject struct {
  name string
  teacher string
  schedule [7] bool
  start_time [7] clock
  end_time [7] clock
}

func NewSubject (name, teacher string, schedule [7] bool, start_time, end_time [7] clock) subject {
  s := subject {name, teacher, schedule, start_time, end_time}
  return s
}

var map_subjects map[string]subject

type user struct {
  name string
  subjects [] string
  past_lessons [] lesson
}

func NewUser (name string, subjects []string, past_lessons []lesson) user {
  u := user {name, subjects, past_lessons}
  return u
}

func nextLesson(u user) lesson {
  var today = int(time.Now().Weekday())-1
  next_subject := ""
  min_time := NewClock(99, 99)
  hour, min, _ := time.Now().Clock()
  current_time := NewClock(hour, min)
  for _, name := range u.subjects {
    subject, ok := map_subjects[name]
    if ok {
      if subject.schedule[today] {
        if (geq(&subject.end_time[today], &current_time)) {
           if(leq(&subject.start_time[today], &min_time)) {
             min_time = subject.start_time[today]
             next_subject = subject.name
           }
        }
      }
    }
  }
  if next_subject == "" {
    return NewLesson("fail", NewClock(-1, -1), NewClock(-1, -1))
  }
  subject := map_subjects[next_subject]
  var ans = NewLesson(subject.name, subject.start_time[today], subject.end_time[today])
  return ans
}

// updates past lessons of the current day
func update_past_lessons(u user) user {
  var today = int(time.Now().Weekday())-1
  hour, min, _ := time.Now().Clock()
  current_time := NewClock(hour, min)
  u.past_lessons = nil
  for _, name := range u.subjects {
    subject, ok := map_subjects[name]
    if ok && subject.schedule[today] {
      if leq(&subject.end_time[today], &current_time) {
        l := NewLesson(subject.name, subject.start_time[today], subject.end_time[today])
        u.past_lessons = append(u.past_lessons, l)
      }
    }
  }
  return u
}

func main() {
  map_subjects = make(map[string]subject)
  var list_subjects []subject
  algebra := NewSubject("algebra", "Casanellas",
              [7]bool{true, true, true, true, true, true, false},
              [7]clock{{8, 0}, {8, 0}, {8, 0}, {8, 0}, {8, 0}, {11, 0}, {-1, -1}},
              [7]clock{{9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {11, 30}, {-1, -1}})
  list_subjects = append(list_subjects, algebra)
  map_subjects["algebra"] = algebra
  calcul := NewSubject("calcul", "Noy",
              [7]bool{true, false, true, false, true, true, false},
              [7]clock{{10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {11, 30}, {-1, -1}},
              [7]clock{{11, 0}, {11, 0}, {11, 0}, {11, 0}, {11, 0}, {12, 30}, {-1, -1}})
  list_subjects = append(list_subjects, calcul)
  map_subjects["calcul"] = calcul
  u1 := NewUser("Max", []string{"algebra", "calcul", "io"}, []lesson{})
  fmt.Println(u1.name)
  fmt.Println(nextLesson(u1))
}
