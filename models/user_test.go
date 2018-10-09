package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_hash(t *testing.T) {
	Convey("test hash", t, func() {
		pwd, err := hash("123456", 16, 1000)
		So(err, ShouldBeNil)
		res, err := hashCompare("123456", pwd)
		So(err, ShouldBeNil)
		// t.Error("password:", pwd)
		So(res, ShouldEqual, true)
	})
}
