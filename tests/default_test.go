package test

import (
	"fmt"
	"mygo/models"
	"mygo/others"
	_ "mygo/routers"
	"mygo/tools"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}

func TestAddThird(t *testing.T) {
	third := models.Third{
		Id:       0,
		Node:     "",
		Company:  "",
		Isp:      "",
		Price:    0,
		Ctype:    "",
		Cratio:   0,
		Tapetype: 0,
		Project:  "",
		Cname:    "",
		Ns:       "",
		Status:   0,
		Username: "",
		Tim:      "",
	}
	num, err := models.AddThird(&third)
	fmt.Print(num, err)
}

func TestSha1Encode(t *testing.T) {
	t.Log(tools.Sha1Encode("lbs"))
}

func TestDayTime(t *testing.T) {
	t.Log(tools.GetEndTime(time.Now()), tools.GetStartTime(time.Now()))
}

func TestMonthTime(t *testing.T) {
	t.Log(tools.GetMonthStart(time.Now()), tools.GetMonthEnd(time.Now()))
}

func TestStrToTime(t *testing.T) {
	t.Log(tools.StrToTime("2020-1"))
}

func TestIfCidrCorrect(t *testing.T) {
	t.Log(tools.IfCidrCorrect("192.0.2.1/24"))
}

func TestIpToInt(t *testing.T) {
	t.Log(tools.IpToInt("127.0.0.1"))
}

func TestMain1(t *testing.T) {
	others.Main()
}


