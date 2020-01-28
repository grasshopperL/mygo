/**
 * @Author: liubaoshuai3
 * @Date: 2020/1/28 13:42
 * @File: third
 * @Description:
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Third struct {
	Id       int    `orm:"pk;auto"`
	Node     string `orm:"size(30)"`
	Company  string `orm:"size(50)"`
	Isp      string
	Price    float64
	Ctype    string `orm:"size(30)"`
	Cratio   float64
	Tapetype int
	Project  string
	Cname    string
	Ns       string
	Status   int
	Username string `orm:"size(30)"`
	Tim      string
}

func init()  {
	orm.RegisterModelWithPrefix("", new(Third))
}

// add one third company
func AddThird(third *Third) (int64, error) {
	o := orm.NewOrm()
	_ = o.Using("write")
	return o.Insert(third)
}

// add companies by batch
func AddThirds(thirds []Third) (int64, error) {
	o := orm.NewOrm()
	_ = o.Using("write")
	return o.InsertMulti(len(thirds), thirds)
}

// del one third company
func DelThird(id int) (bool, error)  {
	o := orm.NewOrm()
	_ = o.Using("read")
	_, err := o.Delete(&Third{Id: id})
	if err != nil {
		return false, err
	}
	return true, nil
}

// get one company's info by id
func GetCompanyInfoById(id string) (third Third, err error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("third").
		Where("id = ?")
	sql := qb.String()
	o := orm.NewOrm()
	err = o.Raw(sql, id).QueryRow(&third)
	return
}

// get all company
func GetCompanyList(params map[string]interface{}) ([]Third, int, error) {
	var resultList []Third
	sql := "select * from third where 1=1"
	sqlCount := "select count(*) as count from third where 1=1"
	if params["node"] != "" && params["node"].(string) != "" {
		sql += "and node = '" + params["node"].(string) + "'"
		sqlCount += "and node = '" + params["node"].(string) + "'"
	}
	if params["isp"] != "" && params["isp"].(string) != "" {
		sql += "and isp = '" + params["isp"].(string) + "'"
		sqlCount += "and isp = '" + params["isp"].(string) + "'"
	}
	if params["company"] != "" && params["company"].(string) != "" {
		sql += "and company = '%" + params["company"].(string) + "%' "
		sqlCount += "and company = '%" + params["company"].(string) + "%'"
	}
	var countResult []orm.Params
	o := orm.NewOrm()
	_ = o.Using("read")
	_, err := o.Raw(sqlCount).Values(&countResult)
	if err != nil {
		return resultList, 0, err
	}
	var count int
	if len(countResult) > 0 && countResult[0]["count"] != nil {
		count, _ = strconv.Atoi(countResult[0]["count"].(string))
		//testCount = countResult[0]["count"].(int)
	}
	order := "id"
	by := "desc"
	if params["order"] != nil && params["order"].(string) != "" {
		order = params["order"].(string)
	}
	if params["by"] != nil && params["by"].(string) != "" {
		by = params["by"].(string)
	}
	sql += " order by " + order + " " + by
	_, err = o.Raw(sql).QueryRows(&resultList)
	if err != nil {
		return resultList, count, err
	}
	return resultList, count, nil
}