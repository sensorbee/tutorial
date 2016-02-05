package facedetect

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"pfi/sensorbee/sensorbee/core"
	"pfi/sensorbee/sensorbee/data"
	"testing"
)

func TestNewCascadeClassifier(t *testing.T) {
	Convey("Given a SensorBee's core.Context", t, func() {
		ctx := &core.Context{}
		Convey("When create state with empty map", func() {
			params := data.Map{}
			_, err := NewCascadeClassifier(ctx, params)
			Convey("Then should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Convey("When create state with not exist file name", func() {
			params := data.Map{
				"file": data.String("not_exist_file"),
			}
			_, err := NewCascadeClassifier(ctx, params)
			Convey("Then should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})
		Convey("When create state with file name", func() {
			xml := `<?xml version="1.0"?>
<opencv_storage>
<cascade>
  <stageType>BOOST</stageType>
  <featureType>LBP</featureType>
  <height>34</height>
  <width>20</width>
  <stageParams>
    <boostType>GAB</boostType>
    <minHitRate>0.1</minHitRate>
    <maxFalseAlarm>0.1</maxFalseAlarm>
    <weightTrimRate>0.1</weightTrimRate>
    <maxDepth>1</maxDepth>
    <maxWeakCount>100</maxWeakCount></stageParams>
  <featureParams>
    <maxCatCount>256</maxCatCount>
    <featSize>1</featSize></featureParams>
  <stageNum>0</stageNum>
  <stages></stages>
  <features>
    <_>
      <rect></rect></_>
  </features>
</cascade>
</opencv_storage>
`
			err := ioutil.WriteFile("_test_for_face_detect.xml", []byte(xml), 0644)
			So(err, ShouldBeNil)
			Reset(func() {
				os.Remove("_test_for_face_detect.xml")
			})
			params := data.Map{
				"file": data.String("_test_for_face_detect.xml"),
			}
			st, err := NewCascadeClassifier(ctx, params)
			So(err, ShouldBeNil)
			Reset(func() {
				st.Terminate(ctx)
			})
			Convey("Then state should be created", func() {
				cc, ok := st.(*cascadeClassifier)
				So(ok, ShouldBeTrue)
				So(cc.classifier, ShouldNotBeNil)
			})
		})
	})
}
