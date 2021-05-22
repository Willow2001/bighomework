package Model

import "time"

type Answer struct{//评论的结构体
	//id、回复的文章、正文、回复的评论的id、作者、评论数、发表时间
	AnswerId int `db:"answer_id"`
	ArticleId int64 `db:"answer_articleid"`
	AnswerContent string `db:"answer_content"`
	Responseanswer int `db:"responseanswerid"`
	ResponserName string `db:"responser_author"`
	AnswerAnswerCount uint32 `db:"answercount"`
	AnswerLoadTime time.Time `db:"loadtime"`
}
