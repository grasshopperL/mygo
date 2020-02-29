/**
 * @Author: liubaoshuai3
 * @Date: 2020/1/29 11:16
 * @File: tool
 * @Description:
 */

package tools

import (
	"bytes"
	"cdnocs/tool"
	"crypto/md5"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	_ "go/types"
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

// MD5 encode
func Md5Encode(s string) string {
	if s == "" {
		return ""
	}
	c := []byte(s)
	ms := md5.New()
	ms.Write(c)
	cs := ms.Sum(nil)
	return hex.EncodeToString(cs)
}

// base64 encode
func b64Encode(s string) string {
	var e bytes.Buffer
	c := base64.NewEncoder(base64.StdEncoding, &e)
	_, _ = c.Write([]byte(s))
	_ = c.Close()
	return e.String()
}

// base64 decode
func b64Decode(s string) string {
	var d bytes.Buffer
	c := base64.NewDecoder(base64.StdEncoding, &d)
	con := make([]byte, 215)
	_, _ = c.Read(con)
	return string(con)
}

// get ip address of local
func GetLocalIp(n int) string {
	addrs, _ := net.InterfaceAddrs()
	ips := make([]string, 0)
	for _, add := range addrs {
		if ip, ok := add.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				ips = append(ips, ip.IP.String())
			}
		}
	}
	return ips[n]
}

// send http request
func Curl(method string, url string, params map[string]interface{}, jsonparam string, header map[string]string, isHttps bool) (string, error) {
	req := new(httplib.BeegoHTTPRequest)
	if method == "get" {
		if params != nil {
			for k, v := range params {
				if strings.Contains(url, "?") {
					url += fmt.Sprintf("&%s=%s", k, InToStr(v))
				} else {
					url += fmt.Sprint("?%s=%s", k, InToStr(v))
				}
			}
		}
		req = httplib.Get(url)
	} else if method == "post" {
		req := httplib.Post(url)
		if params != nil && len(params) > 0 {
			for k, v := range params{
				req.Param(k, InToStr(v))
			}
		}  else if jsonparam != "" {
			req.Body(jsonparam)
		}
	} else if method == "put" {
		req = httplib.Put(url)
		if params != nil && len(params) > 0 {
			for k, v := range params{
				req.Param(k, InToStr(v))
			}
		} else if jsonparam != "" {
			req.Body(jsonparam)
		}
	} else if method == "delete" {
		if params != nil {
			for k, v := range params{
				if strings.Contains(url, "?") {
					url += fmt.Sprintf("&%s=%s", k, InToStr(v))
				} else {
					url += fmt.Sprintf("?%s=%s", k, InToStr(v))
				}
			}
		}
		req = httplib.Delete(url)
	}
	req.SetTimeout(120 * time.Second, 120 * time.Second)
	if header != nil {
		for k, v := range header {
			req.Header(k, v)
		}
	} else {
		req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	}
	if isHttps {
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return req.String()
}

// interface convert to string
func InToStr(v interface{}) string {
	switch t := v.(type) {
	case int:
		return strconv.Itoa(t)
	case int64:
		return strconv.FormatInt(t, 10)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case bool:
		if t == true {
			return "1"
		} else {
			return "0"
		}
	case string:
		return t
	default:
		return ""
	}
}
// email send successfully or not
func IfSendSuccessful(to string, title string, content string, cc string, mode string,) bool {
	if to == "" || title == "" || content == "" {
		return false
	}
	if mode == "" {
		mode = "happy"
	}
	params := make(map[string]interface{})
	params["_to"] = to
	params["_title"] = title
	params["_content"] = content
	params["_cc"] = cc
	params["_mode"] = mode
	url := "https://lbs_alan.com/send"
	data, _ := json.Marshal(params)
	header := make(map[string]string)
	header["Content-Type"] = "application/json;charset=utf-8"
	_, err := tool.Curl("post", url, nil, string(data), header, true)
	if err != nil {
		return false
	}
	return true
}

func GetCidrIpList(cidr string) []string {
	ips := make([]string, 0)
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return ips
	}
	ip := fmt.Sprintf("%s", ipNet.IP)
	ipInt := IpToInt(ip)
	n, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	s := ""
	for i:= 1; i <= (32 - n) ; i++  {
		if i == n {
			s += "0"
		} else {
			s += "1"
		}
	}
	maxNum := int(BinToDec(s) + ipInt)
	for i := int(ipInt) + 2; i < maxNum; i++ {
		ips = append(ips, IntToIp(i))
	}
	return ips
}