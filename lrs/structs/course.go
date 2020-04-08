package structs

//Course struct contains course info as well as a list of learning events
type Course struct {
	CourseCode    string  `json:"coursecode"`
	CourseName    string  `json:"courseName"`
	Version       string  `json:"version"`
	Language      string  `json:"language"`
	TotalSections int     `json:"totalSections"`
	Events        []Event `json:"events"`
}
