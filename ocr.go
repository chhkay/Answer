package Answer

import (
	"github.com/chhkay/Answer/ocr"
	"github.com/chhkay/Answer/util"
)

//Ocr ocr 识别图片文字
type Ocr interface {
	GetText(imgPath string) (string, error)
}

//NewOcr 使用哪种ocr识别
func NewOcr(cfg *util.Config) Ocr {
	if cfg.OCR == "tesseract" {
		return ocr.NewTesseract()
	}
	return ocr.NewBaidu(cfg)
}

func tesseractOCR() Ocr {
	return ocr.NewTesseract()
}

func baiduOCR(cfg *util.Config) Ocr {
	return ocr.NewBaidu(cfg)
}
