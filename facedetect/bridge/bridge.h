#ifndef _FACEDETECT_BRIDGE_H_
#define _FACEDETECT_BRIDGE_H_

#include "../../../opencv/bridge/opencv_bridge.h"

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

typedef struct Rect {
  int x;
  int y;
  int width;
  int height;
} Rect;
typedef struct Rects {
  Rect* rects;
  int length;
} Rects;

#ifdef __cplusplus
typedef cv::CascadeClassifier* CascadeClassifier;
#else
typedef void* CascadeClassifier;
#endif

CascadeClassifier CascadeClassifier_New();
void CascadeClassifier_Delete(CascadeClassifier cs);
int CascadeClassifier_Load(CascadeClassifier cs, const char* name);
struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, MatVec3b img);
void Rects_Delete(struct Rects rs);
void DrawRectsToImage(MatVec3b img, struct Rects rects);

#ifdef __cplusplus
}
#endif

#endif //_FACEDETECT_BRIDGE_H_
