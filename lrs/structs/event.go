package structs

//Event struct is a single learning event
type Event struct {
	Section   int  `json:"section"`
	Completed bool `json:"completed"`
}
