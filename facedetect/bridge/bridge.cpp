#include "bridge.h"

#include <string.h>

CascadeClassifier CascadeClassifier_New() {
  return new cv::CascadeClassifier();
}

void CascadeClassifier_Delete(CascadeClassifier cs) {
  delete cs;
}

int CascadeClassifier_Load(CascadeClassifier cs, const char* name) {
  return cs->load(name);
}

struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, MatVec3b img) {
  std::vector<cv::Rect> faces;
  cs->detectMultiScale(*img, faces); // TODO control default parameter
  Rect* rects = new Rect[faces.size()];
  for (size_t i = 0; i < faces.size(); ++i) {
    Rect r = {faces[i].x, faces[i].y, faces[i].width, faces[i].height};
    rects[i] = r;
  }
  Rects ret = {rects, (int)faces.size()};
  return ret;
}

void Rects_Delete(struct Rects rs) {
  delete rs.rects;
}

void DrawRectsToImage(MatVec3b img, struct Rects rects) {
  for (int i = 0; i < rects.length; ++i) {
    Rect r = rects.rects[i];
    cv::rectangle(*img, cv::Point(r.x, r.y), cv::Point(r.x+r.width, r.y+r.height),
      cv::Scalar(0, 200, 0), 3, CV_AA);
  }
}
