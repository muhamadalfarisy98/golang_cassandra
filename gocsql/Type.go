package main

// structure that represent our data
type Student struct {
	// exported fields
	ID        int    `json:"ID"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Age       int    `json:"Age"`
}
