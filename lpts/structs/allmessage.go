package structs

//AllMessage struct returns progress for all courses
type AllMessage struct {
	CourseProgress []Course `json:"courseProgress"`
}
