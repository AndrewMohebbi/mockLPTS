package structs

//Course struct contains the progress made on a course
type Course struct {
	CourseCode string  `json:"coursecode"`
	CourseName string  `json:"courseName"`
	Version    string  `json:"version"`
	Language   string  `json:"language"`
	Percentage float64 `json:"percentage"`
}
