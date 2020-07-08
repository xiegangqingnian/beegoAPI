package controllers

type DuplicateTrxErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   Error  `json:"error"`
}
type Details struct {
	Message    string `json:"message"`
	File       string `json:"file "`
	LineNumber int    `json:"line_number "`
	Method     string `json:"method "`
}
type Error struct {
	Code    int       `json:"code"`
	Name    string    `json:"name"`
	What    string    `json:"what"`
	Details []Details `json:"details"`
}
