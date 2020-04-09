// Package calculate contains the functions which calculate progress
package calculate

import (
	"encoding/json"
	"fmt"

	"lpts/structs"
)

// All calculates progress for all courses
func All(body []byte) (structs.AllMessage, error) {

	data := struct {
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
	}{}

	if err := json.Unmarshal(body, &data); err != nil {
		return structs.AllMessage{}, fmt.Errorf("calculate/all unmarshalling failed: %v", err)
	}

	// Loop through data and calculate progress for each course
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

	return progress, nil
}
