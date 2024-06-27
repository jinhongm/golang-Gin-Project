package core

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	c, err := elastic.NewClient(
		elastic.SetURL(global.Config.Es.URL()),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.Es.User, global.Config.Es.Password),
	)
	if err != nil {
		logrus.Fatalf("Fail to connect to the elastic search %s", err.Error())
	}
	return c
}
