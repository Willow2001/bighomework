package dao

import (
	"Forum/Model"
	"fmt"
	"time"
)
//插入用户（注册）
func InsertUser(user *Model.UserInfo) (userid int64) {
	sqlStr := `insert into user(user_name,user_password,createtime) values (?,?,?)`//插入语句
	username :=user.Name
	password :=user.Password
	createtime := time.Now()
	user.CreateTime =createtime
	ret, err := db.Exec(sqlStr,username,password,createtime)
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


//根据Id查找用户
func GetUerId(Id int64) (user *Model.UserInfo) {
	if Id <0{
		return
	}
	sqlStr := `select *
				from user 
				where user_id=?
				`
	user = &Model.UserInfo{}
	err := db.Get(user,sqlStr,Id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	return
}
//根据名字查找用户
func GetUerName(Name string) (user *Model.UserInfo) {
	sqlStr := `select *
				from user 
				where user_name=?
				`
	user = &Model.UserInfo{}
	err := db.Get(user,sqlStr,Name)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	return
}
//给文章点赞或者取消点赞
func UserApproveArticle(userId,articleId int64){
	useridback := 0
	//首先查询是否有点赞记录
	sqlStrQuery := `select user_id
				from userapprove 
				Where user_id =? and article_id =?`
	err := db.Get(useridback,sqlStrQuery,userId,articleId)
	if err != nil {//查找失败，记录不存在
		sqlStrInsert := `insert into 
			userapprove(user_id,article_id,createtime) 
			values (?,?,?)`//插入语句
		createtime := time.Now()
		_, err = db.Exec(sqlStrInsert,userId,articleId,createtime)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
		sqlStrUpdate := `Update article set approvecount = approvecount+1 where article_id =?`//更新点赞数
		_, err = db.Exec(sqlStrUpdate,articleId)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
	}else{//取消给文章的点赞
		sqlStrDelete := `Delete From userapprove Where user_id = ? and article_id=?`//删除点赞记录
		_, err := db.Exec(sqlStrDelete,userId,articleId)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
		sqlStrUpdate := `Update article set approvecount = approvecount-1 where article_id = ?`//更新点赞数

		_, err = db.Exec(sqlStrUpdate,articleId)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
	}

}
//给文章收藏或者取消收藏
func UserCollectArticle(userId,articleId int64){
	useridback := 0
	//首先查询是否有收藏记录
	sqlStrQuery := `select user_id
				from usercollection 
				Where user_id = ? and article_id = ?`
	err := db.Get(useridback,sqlStrQuery,userId,articleId)
	if err != nil {//查找失败，记录不存在
		sqlStrInsert := `insert into 
			usercollection(user_id,article_id,createtime) 
			values (?,?,?)`//插入语句
		createtime := time.Now()
		_, err = db.Exec(sqlStrInsert,userId,articleId,createtime)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
		sqlStrUpdate := `Update article set collectcount = collectcount+1 where article_id = ?`//更新收藏数
		_, err = db.Exec(sqlStrUpdate,articleId)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
	}else{//取消给文章的收藏
		sqlStrDelete := `Delete From usercollection Where user_id = ? and article_id = ?`//删除收藏记录
		_, err := db.Exec(sqlStrDelete,userId,articleId)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
		sqlStrUpdate := `Update article set collectcount = collectcount-1 where article_id = ?`//更新收藏数
		_, err = db.Exec(sqlStrUpdate,articleId)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
		}
	}

}
//根据用户id查找用户的收藏夹
func GetUserCollection(userId int64) (articleIdList []int64) {
	if userId <0{
		return
	}
	sqlStr := `select article_id
				from usercollection
				where user_id=?
				`
	err := db.Select(&articleIdList,sqlStr,userId)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	return
}