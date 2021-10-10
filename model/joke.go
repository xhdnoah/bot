package model

type JokeModel struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func (joke JokeModel) TableName() string {
	return "jokes"
}

func (joke JokeModel) Take() JokeModel {
	DB.Self.Order("random()").Limit(1).Find(&joke)
	return joke
}
