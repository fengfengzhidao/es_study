package core

import (
	"es_study/global"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func EsConnect() {

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("elastic", "xxxxxx"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	global.ESClient = client
}
