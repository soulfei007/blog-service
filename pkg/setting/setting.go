package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type Setting struct {
	vp *ini.File
}

func NewSetting() (*Setting, error) {

	Cfg, err := ini.Load("configs/app.ini")
	if err != nil {

		log.Fatalf("Fail to parse 'confings/app.ini': %v", err)

	}

	return &Setting{vp: Cfg}, nil
}
