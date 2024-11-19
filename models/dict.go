package models

type Dict struct {
	Word           string          `json:"word"`
	Pos            []string        `json:"pos"`
	Verbs          []Verb          `json:"verbs"`
	Pronunciations []Pronunciation `json:"pronunciation"`
	Definitions    []Definition    `json:"definition"`
}
type Verb struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}
type Pronunciation struct {
	Pos  string `json:"pos"`
	Lang string `json:"lang"`
	URL  string `json:"url"`
	Pron string `json:"pron"`
}
type Definition struct {
	ID          int       `json:"id"`
	Pos         string    `json:"pos"`
	Text        string    `json:"text"`
	Translation string    `json:"translation"`
	Examples    []Example `json:"example"`
}
type Example struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}
