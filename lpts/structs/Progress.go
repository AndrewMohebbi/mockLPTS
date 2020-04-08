package structs

//Progress struct contains the progress made on a course
type Progress struct {
	CourseCode string  `json:"coursecode"`
	CourseName string  `json:"courseName"`
	Version    string  `json:"version"`
	Language   string  `json:"language"`
	Percentage float64 `json:"percentage"`
}
