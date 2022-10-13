package routers

import (
	"ginblog/config"
	"ginblog/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_getConfig(t *testing.T) {
	Convey("test getConfig", t, func() {
		var err error
		config.C, err = test.LoadConfig()
		So(err, ShouldBeNil)
		val, err := getConfig("String", "globaltitle", "")
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "阿汤哥的博客")
	})
}
