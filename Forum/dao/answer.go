package dao

import (
	"Forum/Model"
	"fmt"
	"time"
)

//发表评论
func AnswerArticle (answer *Model.Answer) (AnswerId int64) {
	if answer == nil{
		return
	}//防空验证
	sqlStr := `insert into 
			answer(answer_id,answer_articleid,answer_content,responseanswerid,responser_author,answercount,loadtime) 
			values (?,?,?,?,?,?,?)`//插入语句
	answercount :=0
	loadtime := time.Now()
	ret, err := db.Exec(sqlStr,answer.AnswerId,answer.ArticleId,answer.AnswerContent,answer.Responseanswer,answer.ResponserName,answercount,loadtime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	sqlStrUpdateArticle := `Update article set answercount = answercount+1 where article_id = ?`//更新文章评论数
	_, err = db.Exec(sqlStrUpdateArticle,answer.ArticleId)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
	}
	if answer.Responseanswer != 0{//这是一条楼中楼回复
		sqlStrUpdateAnswer := `Update answer set answercount = answercount+1 where answer_id = ?`//更新评论的评论数
		_, err = db.Exec(sqlStrUpdateAnswer,answer.AnswerId)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	return theID
}
