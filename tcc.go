package tcc

import (
	"net/http"
	"time"
	"github.com/jinzhu/gorm"
)

type tcc interface {

}

type TCC struct {
	httpCli *http.Client
	db *gorm.DB
	tccConfig http_tcc.Config
}

func NewTCC(mysqlConfig database.Config, tccConfig http_tcc.Config) error {
	httpCli := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * 30,
	}
	return &DefaultTCC{httpCli: httpCli}
}
