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
	Db      int    `form:"db" label:"数据库名" json:"db" binding:"required"`
}
