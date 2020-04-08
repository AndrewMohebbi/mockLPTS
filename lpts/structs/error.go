package structs

//Error struct is what you send when a request fails
type Error struct {
	Error string `json:"error"`
}
