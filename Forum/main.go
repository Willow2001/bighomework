package main

import (
	"Forum/controller"
	"Forum/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dns := "root:151413Mn++@tcp(127.0.0.1:3306)/forum?charset=utf8mb4&parseTime=True"
	err := dao.InitDB(dns)//账号数据库
	if err !=nil{
		fmt.Print(err)
	}
	articleGroup := r.Group("/article")
	{
		articleGroup.GET("/list",controller.MostAnswerArticle)//利用querystring搜索讨论度高的文章
		articleGroup.GET("/read/:articleid",controller.ReadArticle)//利用path参数搜索文章阅读
		articleGroup.GET("/search/:articlename",controller.SearchArticle)//利用path参数关键词模糊搜索文章
		articleGroup.GET("/approver/:articleid",controller.SearchArticleApprover)//利用path参数查看给文章点过赞的用户名单
		articleGroup.GET("/answers/:articleid",controller.SearchArticleAnswers)//利用path参数查看文章的所有评论

	}
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register",controller.UserRegister)//注册
		userGroup.POST("/login",controller.UserLogin)//登录
		userGroup.GET("/logout",controller.UserLogout)//注销
		userGroup.POST("/postblog",controller.UserPostArticle)//发表文章
		userGroup.GET("/collection",controller.SearchUsersCollection)//查看自己的收藏夹
		userGroup.GET("/approve/:articleid",controller.ApproveArticle)//利用path参数给文章点赞或者取消点赞
		userGroup.GET("/collect/:articleid",controller.CollectArticle)//利用path参数给文章收藏或者取消收藏
		userGroup.POST("/postanswer/:articleid",controller.UserPostAnswer)//发表评论
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
