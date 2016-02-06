package plugin

import (
	"pfi/sensorbee/sensorbee/bql/udf"
	"pfi/sensorbee/tutorial/facedetect"
)

// initialize scouter components. this init method will be called by
// SensorBee customized main.go.
//
//  import(
//      _ "pfi/sensorbee/tutorial/facedetect/plugin"
//  )
func init() {
	// cascade classifier
	udf.MustRegisterGlobalUDSCreator("facedetect_cascade_classifier",
		udf.UDSCreatorFunc(facedetect.NewCascadeClassifier))
	udf.MustRegisterGlobalUDF("facedetect_detect_multi_scale",
		udf.MustConvertGeneric(facedetect.DetectMultiScale))
	udf.MustRegisterGlobalUDF("facedetect_draw_rects",
		udf.MustConvertGeneric(facedetect.DrawRectsToImage))

	// mount image
	udf.MustRegisterGlobalUDSCreator("facedetect_shared_image",
		udf.UDSCreatorFunc(facedetect.NewSharedImage))
	udf.MustRegisterGlobalUDF("facedetect_mount_image",
		udf.MustConvertGeneric(facedetect.MountAlphaImage))
}
