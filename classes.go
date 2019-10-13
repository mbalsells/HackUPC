package main

import (
    "fmt"
    "time"
    "sort"
)

type clock struct {
    hour int
    min int
}

type lesson struct {
    subject         string
    start_time      clock
    end_time        clock
    feedback        [11]int
}

func min_lesson (l1, l2 lesson) lesson {
    if leq(l1.start_time, l2.start_time) {
        return l1
    }
    return l2
}

func NewLesson (subject string, start_time, end_time clock) lesson {
    var feedback [11]int
    for i := range feedback {
      feedback[i] = 0
    }
    l := lesson {subject, start_time, end_time, feedback}
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
    // dni for student
    name string
    teacher string
    schedule [7] lesson
}

func NewSubject (name, teacher string, schedule [7] lesson) subject {
    s := subject {name, teacher, schedule}
    return s
}

var map_subjects map[string]subject


type user struct {
    dni int
    name string
    subjects [] string
    feedback [] int
}

func NewUser (dni int, name string, subjects []string) user {
    n := len(subjects)
    var feedback [] int
    for i := 0; i < n; i++ {
      feedback = append(feedback, -1)
    }
    u := user {dni, name, subjects, feedback}
    return u
}

func getInfoUser (u user) (dni int, name string, subjects []string, lessons []lesson, feedback []int) {
  return u.dni, u.name, u.subjects, daily_schedule(u), u.feedback
}

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

func set_feedback(u user, points int, subject_name string) {
      today := int(time.Now().Weekday()) - 1
      for i, x := range u.subjects {
          if x == subject_name {
              u.feedback[i] = points
              for j, y := range list_subjects {
                  if y.name == subject_name {
                      list_subjects[j].schedule[today].feedback[points]++
                  }
              }
          }
      }
}

/*
func nextLesson(u user) lesson {
    var today = int(time.Now().Weekday()) - 1
    next_lesson := NewLesson("fail", 99, 99)

    hour, min, _ := time.Now().Clock()
    current_time := NewClock(hour, min)

    for _, temp_subject_string := range u.subject{
        temp_lesson = map_subjects[temp_subject_string].subject_lessons[day]

        if geq(temp_lesson.end_time, current_time) {
            next_lesson = min_lesson(next_lesson, temp_lesson)
        }
    }

    return next_lesson
}


// updates past lessons of the current day
func update_past_lessons(u user) user {
    var today = int(time.Now().Weekday())-1

    hour, min, _ := time.Now().Clock()
    current_time := NewClock(hour, min)

 && subject.sche
    u.past_lessons = nil
    for _, name := range u.subjects {
        subject, ok := map_subjects[name]
        if okdule[today] {
            if leq(subject.end_time[today], current_time) {
                l := NewLesson(subject.name, subject.start_time[today], subject.end_time[today])
                u.past_lessons = append(u.past_lessons, l)
            }
        }
    }
    return u
}
*/

var list_subjects []subject

func main() {
  map_subjects = make(map[string]subject)


 algebra := NewSubject("algebra", "Casanellas",
                      [7]lesson{NewLesson("algebra", clock{8, 0}, clock{9, 0}),
                                NewLesson("algebra", clock{8, 0}, clock{9, 0}),
                                NewLesson("algebra", clock{8, 0}, clock{9, 0}),
                                NewLesson("algebra", clock{8, 0}, clock{9, 0}),
                                NewLesson("algebra", clock{8, 0}, clock{9, 0}),
                                NewLesson("algebra", clock{8, 0}, clock{9, 0}),
                                NewLesson("algebra", clock{8, 0}, clock{9, 0})})
 list_subjects = append(list_subjects, algebra)
 map_subjects["algebra"] = algebra
 calcul := NewSubject("calcul", "Noy",
                      [7] lesson {NewLesson("calcul", clock{10, 0}, clock{11, 0}),
                                  NewLesson("calcul", clock{-1, -1}, clock{-1, -1}),
                                  NewLesson("calcul", clock{10, 0}, clock{11, 0}),
                                  NewLesson("calcul", clock{-1, -1}, clock{-1, -1}),
                                  NewLesson("calcul", clock{10, 0}, clock{11, 0}),
                                  NewLesson("calcul", clock{11, 30}, clock{12, 30}),
                                  NewLesson("calcul", clock{-1, -1}, clock{-1, -1})})
 list_subjects = append(list_subjects, calcul)

 map_subjects["calcul"] = calcul
 u1 := NewUser(1234, "Max", []string{"algebra", "calcul", "io"})
 fmt.Println(u1.dni)
 //fmt.Println(nextLesson(u1))
 fmt.Println(getInfoUser(u1))
 set_feedback(u1, 7, "calcul")
 set_feedback(u1, 9, "algebra")
 fmt.Println("feedback user: ", u1.feedback, " \nfeedback lesson: ", list_subjects)
}
