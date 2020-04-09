// Package calculate contains the functions which calculate progress
package calculate

import (
	"encoding/json"
	"log"

	"lpts/structs"
)

// All calculates progress for all courses
func All(body []byte) structs.AllMessage {

	type message struct {
		ProfileID string
		Courses   []struct {
			CourseCode    string
			CourseName    string
			Version       string
			Language      string
			TotalSections int
			Events        []struct {
				Section   int
				Completed bool
			}
		}
	}

	data := message{}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("CalculateAll: unmarshalling failed")
		return structs.AllMessage{}
	}

	var courses []structs.Course
	for _, c := range data.Courses {
		var completed int
		for _, e := range c.Events {
			if e.Completed == true {
				completed++
			}
		}
		percent := float64(completed) / float64(c.TotalSections) * 100
		course := structs.Course{
			CourseCode: c.CourseCode, CourseName: c.CourseName,
			Version: c.Version, Language: c.Language, Percentage: percent,
		}
		courses = append(courses, course)
	}

	progress := structs.AllMessage{CourseProgress: courses}

	return progress
}
