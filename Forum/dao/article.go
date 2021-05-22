package dao

import (
	"Forum/Model"
	"fmt"
	"time"
)

// 插入（添加）文章
func InsertArticle (article *Model.ArticleDetail) (articleid int64) {
	if article == nil{
		return
	}//防空验证
	sqlStr := `insert into 
			article(category_id,title,summary,author_name,approvecount,collectcount,answercount,article_content,loadtime,updatetime) 
			values (?,?,?,?,?,?,?,?,?,?)`//插入语句
	approvecount :=0
	collectcount :=0
	answercount :=0
	loadtime := time.Now()
	updatetime := loadtime

	ret, err := db.Exec(sqlStr,article.ArticleInfo.CategoryId,article.Title,article.Summary,article.AuthorName,
		approvecount,collectcount,answercount,article.Content,loadtime,updatetime)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	return theID
}

//根据文章id获取单个文章信息
func GetArticleId(articleid int64) (article *Model.ArticleInfo) {
	if articleid <0{
		return
	}
	article = &Model.ArticleInfo{}
	sqlStr := `select article_id,title,summary,author_name,approvecount,collectcount,answercount,loadtime,updatetime
				from article 
				where article_id=?
				`
	err := db.Get(article,sqlStr,articleid)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	return
}

//根据文章id获取单个文章（阅读原文）
func GetArticleIdDetail(articleid int64) (article *Model.ArticleDetail) {
	if articleid <0{
		return
	}
	sqlStr := `select *
				from article 
				where article_id=?
				`
	article = &Model.ArticleDetail{}
	err := db.Get(article,sqlStr,articleid)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	return
}
//根据文章名称模糊搜索单个文章
func GetArticleName(articlename string) (article []*Model.ArticleInfo) {

	sqlStr := `select article_id,title,summary,author_name,approvecount,collectcount,answercount,loadtime,updatetime
				from article 
				where title like ?
				`
	err := db.Select(&article,sqlStr,"%"+ articlename+"%")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	return
}
////讨论热度（评论数）降序排序获取文章分页
func GetArticleListAnswer(pageNum,pageSize int) (articleList []*Model.ArticleInfo) {
	if pageNum < 0 || pageSize < 0 {
		fmt.Printf("get failed")
		return
	}
	sqlStrAnswer := `select article_id,title,summary,author_name,approvecount,collectcount,answercount,loadtime,updatetime
				from article
				order by answercount desc
				limit ?,?`

	errAnswer := db.Select(&articleList, sqlStrAnswer, pageNum, pageSize)
	if errAnswer != nil {
		fmt.Printf("get failed, err:%v\n", errAnswer)
		return
	}
	return
}

//根据文章id获取点过赞的用户id
func GetApproverId(articleid int64) (userIdList []int64) {
	if articleid <0{
		return
	}
	sqlStr := `select user_id
				from userapprove 
				where article_id=?
				`
	err := db.Select(&userIdList,sqlStr,articleid)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	return
}
//根据文章id获取所有评论
func GetArticleAnswer(articleid int64) (answerList []*Model.Answer) {
	if articleid <0{
		return
	}
	sqlStr := `select *
				from answer 
				where answer_articleid=?
				`
	err := db.Select(&answerList,sqlStr,articleid)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	return
}