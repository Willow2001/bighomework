package Model
type Category struct{//分类的结构体
	//分类id、名字、排序
	CategoryId int64 `db:"category_id"`
	CategoryName string `db:"category_name"`
}
