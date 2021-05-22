package controller

import (
	"Forum/Model"
	"Forum/dao"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserRegister(c *gin.Context){//注册
	name := c.PostForm("username")
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")
	if password1 != password2 {
		err := errors.New("注册失败！密码不一致.")
		fmt.Print(err)

	}
	if dao.GetUerName(name) != nil {
		err := errors.New("注册失败！用户名重复.")
		fmt.Print(err)

	}
	user :=&Model.UserInfo{}
	user.Name = name
	user.Password = password2
	dao.InsertUser(user)
	c.String(http.StatusOK,fmt.Sprintf("注册成功\n"))
}
func UserLogin(c *gin.Context ){//登录
	name := c.PostForm("username")
	password := c.PostForm("password")
	passwordTrue := dao.GetUerName(name).Password
	Id := dao.GetUerName(name).Id
	if password == passwordTrue {
		c.String(http.StatusOK,fmt.Sprintf("登录成功！\n"))
		cookie, err := c.Cookie("usercookie") // 获取Cookie
		if err != nil {
			cookie = name
			// 设置Cookie
			c.SetCookie("usercookie", string(Id), 3600, "/", "localhost", true, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	} else{
		c.String(501,fmt.Sprintf("登录失败!\n"))
	}
}
func UserLogout(c *gin.Context ){//注销
	_, err := c.Cookie("usercookie") // 获取Cookie
	if err != nil {
		c.String(501,fmt.Sprintf("请先登录!\n"))
		return
	}
	c.SetCookie("usercookie","", -1, "/", "localhost", true, true)

}

func UserPostArticle(c *gin.Context ){//发表文章
	cookie, err := c.Cookie("usercookie") // 获取Cookie
	if err != nil {
		c.String(501,fmt.Sprintf("请先登录!\n"))
		return
	}
	article := &Model.ArticleDetail{}
	authorid,_ := strconv.ParseInt(cookie, 10, 64)
	article.AuthorName= dao.GetUerId(authorid).Name
	article.Title= c.PostForm("title")
	article.Summary= c.PostForm("summary")
	article.Category.CategoryName = c.PostForm("category")
	article.Content = c.PostForm("content")
	dao.InsertArticle(article)
}
//发表评论
func UserPostAnswer(c *gin.Context ){
	cookie, err := c.Cookie("usercookie") // 获取Cookie
	if err != nil {
		c.String(501,fmt.Sprintf("请先登录!\n"))
		return
	}
	articleId,_ := strconv.ParseInt(c.Param("articleid"), 10, 64)
	answertoid,_ := strconv.Atoi(c.DefaultQuery("answerid", "0")) //无输入则默认直接做为文章的评论
	answer := &Model.Answer{}
	authorid,_ := strconv.ParseInt(cookie, 10, 64)
	answer.ResponserName = dao.GetUerId(authorid).Name
	answer.AnswerContent = c.PostForm("content")
	answer.ArticleId = articleId
	answer.Responseanswer =answertoid
	dao.AnswerArticle(answer)
}

func ReadArticle(c *gin.Context ){//阅读文章
	articleId,_ := strconv.ParseInt(c.Param("articleid"), 10, 64)
	article := dao.GetArticleIdDetail(articleId)
	//输出json结果给调用方
	c.JSON(http.StatusOK,article)
}
func SearchArticle(c *gin.Context ){//模糊搜索文章列表
	articleName := c.Param("articlename")
	articleList := dao.GetArticleName(articleName)
	//输出json结果给调用方
	c.JSON(http.StatusOK,articleList)
}
func MostAnswerArticle(c *gin.Context ){//搜讨论度前number的文章
	postnumber := c.DefaultQuery("number", "10")//无输入则默认前10
	number,_ := strconv.Atoi(postnumber)
	articleList := dao.GetArticleListAnswer(1,number)
	//输出json结果给调用方
	c.JSON(http.StatusOK,articleList)
}
//给文章点赞或者取消点赞
func ApproveArticle(c *gin.Context ){
	cookie, err := c.Cookie("usercookie") // 获取Cookie
	if err != nil {
		c.String(501,fmt.Sprintf("请先登录!\n"))
		return
	}
	userId,_ := strconv.ParseInt(cookie, 10, 64)
	articleId,_ := strconv.ParseInt(c.Param("articleid"), 10, 64)
	dao.UserApproveArticle(userId,articleId)
}
//给文章收藏或者取消收藏
func CollectArticle(c *gin.Context ){
	cookie, err := c.Cookie("usercookie") // 获取Cookie
	if err != nil {
		c.String(501,fmt.Sprintf("请先登录!\n"))
		return
	}
	userId,_ := strconv.ParseInt(cookie, 10, 64)
	articleId,_ := strconv.ParseInt(c.Param("articleid"), 10, 64)
	dao.UserCollectArticle(userId,articleId)
}
//搜给文章点过赞的用户
func SearchArticleApprover(c *gin.Context ) {

	var (
		userName     string
		userNameList []string
	)
	articleId,_ := strconv.ParseInt(c.Param("articleid"), 10, 64)
	userIdList := dao.GetApproverId(articleId)
	for _, userId := range userIdList {
		{
			userName = dao.GetUerId(userId).Name
			userNameList = append(userNameList, userName)
		}
		//输出json结果给调用方
		c.JSON(http.StatusOK, userNameList)
	}
}
//搜文章的所有评论
func SearchArticleAnswers(c *gin.Context ) {

	var (
		answerList []*Model.Answer
	)
	articleId,_ := strconv.ParseInt(c.Param("articleid"), 10, 64)
	answerList = dao.GetArticleAnswer(articleId)
	//输出json结果给调用方
	c.JSON(http.StatusOK, answerList)
	}
//获取用户的收藏夹
func SearchUsersCollection(c *gin.Context ) {
	var (
		article *Model.ArticleInfo
		articleList []*Model.ArticleInfo
	)
	cookie, err := c.Cookie("usercookie") // 获取Cookie,验证登录
	if err != nil {
		c.String(501,fmt.Sprintf("请先登录!\n"))
		return
	}
	userId,_ := strconv.ParseInt(cookie, 10, 64)
	articleIdList := dao.GetUserCollection(userId)
	for _, articleId := range articleIdList {
		{
			article = dao.GetArticleId(articleId)
			articleList = append(articleList, article)
		}
		//输出json结果给调用方
		c.JSON(http.StatusOK, articleList)
	}
}