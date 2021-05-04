package entity

type Student struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
	Batch    string `json:"batch"`
}
