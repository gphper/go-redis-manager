/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-21 13:22:54
 */
package model

type ServiceConfigReq struct {
	ServiceName string `form:"service_name" label:"连接名称" json:"service_name" binding:"required"`
	Host        string `form:"host" label:"IP地址" json:"host" binding:"required"`
	Port        string `form:"port" label:"端口号" json:"port" binding:"required"`
	Password    string `form:"password" label:"密码" json:"password"`
}

type ServiceSwitchReq struct {
	Service string `form:"service" label:"连接名称" json:"service" binding:"required"`
	Db      int    `form:"db" label:"数据库名" json:"db" default:"0"`
}

type RedisKeyReq struct {
	SearchKey  string `form:"key" label:"key值" json:"key" binding:"required"`
	SearchType int    `form:"type" label:"搜索模式" json:"type" binding:"required"`
}

type AddKeyReq struct {
	Key   string `form:"key" label:"字段值" json:"key" binding:"required"`
	Pre   string `form:"prekey" json:"prekey"`
	Type  string `form:"type" label:"类型" json:"type" binding:"required"`
	Field string `form:"field" label:"字段名" json:"field" default:""`
	Value string `form:"value" label:"字段值" json:"value" binding:"required" default:""`
	Score int    `form:"score" label:"分数" json:"score" default:"1"`
	Id    string `form:"id" label:"ID" json:"id" default:"*"`
}

type DelKeyReq struct {
	Key  string `form:"key" label:"字段值" json:"key" binding:"required"`
	Type string `form:"type" label:"类型" json:"type" binding:"required"`
}
