// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// TbUser is the golang structure for table tb_user.
type TbUser struct {
    Id       uint64      `orm:"id,primary" json:"id"`        // 消息给过来的ID                 
    Mobile   string      `orm:"mobile"     json:"mobile"`    // 手机号                         
    Name     string      `orm:"name"       json:"name"`      // 姓名                           
    Password string      `orm:"password"   json:"password"`  // 密码                           
    IsDelete int         `orm:"is_delete"  json:"is_delete"` // 是否删除 1：已删除；0：未删除  
    CreateAt *gtime.Time `orm:"create_at"  json:"create_at"` //                                
    UpdateAt *gtime.Time `orm:"update_at"  json:"update_at"` //                                
}