/**
 * @Author: liubaoshuai3
 * @Date: 2020/1/29 11:16
 * @File: tool
 * @Description:
 */

package tools

import "github.com/astaxie/beego"

// response data content
type ResData struct {
	Code int
	Msg  string
	Data interface{}
}

// get username info in request header
func GetHeader(c *beego.Controller) map[string]string {
	info := make(map[string]string)
	info["username"] = c.Ctx.Input.Header("USERNAME")
	return info
}

// send json to response
func WriteJson(c *beego.Controller, r *ResData)  {
	if r.Code == 0 {
		c.Data["json"] = r.Data
	} else {
		c.Data["json"] = r
	}
	c.ServeJSON()
	c.StopRun()
}