package structs

//Message struct is the message lrs sends to lpts
type Message struct {
	ProfileID string   `json:"profileID"`
	Courses   []Course `json:"courses"`
}
