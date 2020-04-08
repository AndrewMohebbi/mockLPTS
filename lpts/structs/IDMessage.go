package structs

//IDMessage struct returns progress for a profile ID
type IDMessage struct {
	CourseProgress []Progress `json:"courseProgress"`
}
