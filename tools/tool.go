/**
 * @Author: liubaoshuai3
 * @Date: 2020/1/29 11:16
 * @File: tool
 * @Description:
 */

package tools

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/astaxie/beego"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

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

// hash encode
func Sha1Encode(s string) string {
	c := []byte(s)
	sha := sha1.New()
	sha.Write(c)
	cipher := sha.Sum(nil)
	return hex.EncodeToString(cipher)
}

// judge a file is exited or not
func FileIsExits(p string) bool {
	if _, err := os.Stat(p); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// any struct to map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var d = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		d[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return d
}

// convert string time to time by layout
func StrToTime(tim string) time.Time  {
	layout := ""
	if len(tim) <= 10  {
		layout = "2006-01-02"
	} else {
		layout = "2006-01-02 15:04:05"
	}
	loc, _ := time.LoadLocation("Local")
	sTime, _ := time.ParseInLocation(layout, tim, loc)
	return sTime
}

// get the end time of one day
func GetEndTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

// get the start time of one day
func GetStartTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// get the start time of the month
func GetMonthEnd(d time.Time) time.Time {
	return GetEndTime(GetMonthStart(d).AddDate(0, 1, -1))
}

// get the end time of the month
func GetMonthStart(d time.Time) time.Time {
	return GetStartTime(d.AddDate(0, 0, -d.Day() + 1))
}

// judge the cidr is correct or not
func IfCidrCorrect(c string) bool {
	_, ipNet, _ := net.ParseCIDR(c)
	ipS := ipNet.IP.String()
	ipC := strings.Split(c, "/")[0]
	if ipC != ipS {
		return false
	}
	return true
}

// ip address convert to num
func IpToInt(ip string) int64 {
	bits := strings.Split(ip, ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])
	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	return sum
}

//func IfCidrHaveCommon(c1, c2 string) bool {
//	_, ipNet1, _ := net.ParseCIDR(c1)
//	_, ipNet2, _ := net.ParseCIDR(c2)
//	ip1S := ipNet1.IP.String()
//	ip2S := ipNet2.IP.String()
//
//	return false
//}

