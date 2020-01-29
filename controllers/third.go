/**
 * @Author: liubaoshuai3
 * @Date: 2020/1/29 10:52
 * @File: third
 * @Description:
 */

package controllers

import (
	"github.com/astaxie/beego"
	"mygo/models"
	"mygo/tools"
	"strconv"
)

type ThirdController struct {
	beego.Controller
}

func(c *ThirdController) writeJson(r *tools.ResData)  {
	if r.Code == 0 {
		c.Data["json"] = r.Data
	} else {
		c.Data["json"] = r
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *ThirdController) GetThirdList() {
	var res tools.ResData
	var params = make(map[string]interface{})
	var limit, page  = 10, 1
	if c.GetString("limit") != "" {
		limit, _ = strconv.Atoi(c.GetString("limit"))
		params["limit"] = limit
	}
	if c.GetString("page") != "" {
		page, _= strconv.Atoi(c.GetString("page"))
		params["offset"] = (page - 1) * limit
	}
	if company := c.GetString("company"); company != "" {
		params["company"] = company
	}
	if isp := c.GetString("isp"); isp != "" {
		params["isp"] = isp
	}
	if node := c.GetString("node"); node != "" {
		params["node"] = node
	}
	rList, total, err := models.GetCompanyList(params)
	r := map[string]interface{}{
		"data": rList,
		"total": total,
	}
	if err != nil {
		res.Code = 10001
		res.Msg = err.Error()
		res.Data = r
	} else {
		res.Code = 1000
		res.Msg = "success"
		res.Data = r
	}
	c.writeJson(&res)
}