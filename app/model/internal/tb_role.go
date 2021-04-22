// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// TbRole is the golang structure for table tb_role.
type TbRole struct {
    Id       uint64      `orm:"id,primary" json:"id"`        // 主键                           
    Code     string      `orm:"code"       json:"code"`      // 编码                           
    Name     string      `orm:"name"       json:"name"`      // 名称                           
    IsDelete int         `orm:"is_delete"  json:"is_delete"` // 是否删除 1：已删除；0：未删除  
    CreateAt *gtime.Time `orm:"create_at"  json:"create_at"` //                                
    UpdateAt *gtime.Time `orm:"update_at"  json:"update_at"` //                                
}