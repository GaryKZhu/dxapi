package setting

import (
	"log"
	"time"
	"strconv"
	"github.com/go-ini/ini"
)


var (
	Cfg *ini.File

	RunMode string

	HTTPPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration

)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.conf")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.conf': %v", err)
	}

	LoadBase()
	LoadServer()
	//LoadDB()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = ":" + strconv.Itoa(sec.Key("HTTP_PORT").MustInt(8000))
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

