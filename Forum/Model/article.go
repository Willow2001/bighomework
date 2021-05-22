package Model

import "time"

type ArticleInfo struct{//主页显示的文章的结构体
	//id、分类id、标题、文章摘要、作者、点赞数、收藏数、评论数、发表时间、更新时间
	ArticleId int `db:"article_id"`
	CategoryId int `db:"category_id"`
	Title string `db:"title"`
	Summary string `db:"summary"`
	AuthorName string `db:"author_name"`
	ApproveCount uint32 `db:"approvecount"`
	CollectCount uint32 `db:"collectcount"`
	AnswerCount uint32 `db:"answercount"`
	LoadTime time.Time `db:"loadtime"`
	UpdateTime time.Time `db:"updatetime"`
}
type ArticleDetail struct{//文章详情的结构体
	//基本信息、正文、分类
	ArticleInfo
	Content string `db:"article_content"`
	Category
}
