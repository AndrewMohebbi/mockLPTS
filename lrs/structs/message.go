package structs

//Message struct is the message lrs sends to lpts
type Message struct {
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}
