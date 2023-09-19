package model

type Student struct {
	ID            string `json:"id"`
	Code          string `json:"student_code"`
	Name          string `json:"student_name"`
	Address       string `json:"student_address"`
	Occupation    string `json:"student_occupation"`
	JoiningReason string `json:"joining_reason"`
}

type Students struct {
	Students []Student `json:"participants"`
}
