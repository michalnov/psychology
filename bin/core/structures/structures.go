package structures

//Req -
type Req struct {
	UserID     int  `json:"userid"`
	TestID     int  `json:"testid"`
	QuestionID int  `json:"questionid"`
	AnswerID   int  `json:"answerid,ommitempty"`
	NoAnswer   bool `json:"noanswer,ommitempty"`
	GoBack     bool `json:"goback,ommitempty"`
}

//Test -
type Test struct {
	Description string     `json:"description,ommitempty"`
	Key         string     `json:"testkey,ommitempty"`
	TestID      int        `json:"testid,ommitempty"`
	Que         []Question `json:"que,ommitempty"`
}

//Question -
type Question struct {
	Text       string   `json:"text,ommitempty"`
	QuestionID int      `json:"questionid,ommitempty"`
	Ans        []Answer `json:"ans,ommitempty"`
}

//Answer -
type Answer struct {
	Text     string `json:"text,ommitempty"`
	AnswerID int    `json:"answerid,ommitempty"`
	Reverse  bool   `json:"reverse,ommitempty"`
}
