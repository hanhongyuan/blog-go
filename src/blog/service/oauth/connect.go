package oauth

import (
	"blog/fox"
	"blog/fox/db"
	"blog/model"
	"fmt"
)
//第三方
type Connect struct {

}
//快速初始化好
func NewConnect() *Connect {
	return new(Connect)
}
//后台绑定
func (t *Connect)Admin(type_id int, val string, is_uid bool) (*model.Connect, error) {
	if len(val) < 1 {
		return nil,fox.NewError("val 值不能为空:")
	}
	if type_id < 1 {
		return nil,fox.NewError("type_id 值错误:")
	}
	con := model.NewConnect()
	fmt.Println(type_id, val)
	maps := make(map[string]interface{})
	maps["type_id"] = type_id
	//不是UID登陆
	if is_uid {
		maps["uid"] = val
	} else {
		maps["open_id"] = val
	}
	_,err := db.Filter(maps).Get(con)
	if err != nil {
		return nil,fox.NewError("查询错误:"+ err.Error())
	}
	if con.Uid < 1 {
		return nil,fox.NewError("未绑定")
	}
	fmt.Println("con", con)
	return con, nil
}