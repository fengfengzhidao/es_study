package docs

import (
	"context"
	"es_study/global"
	"es_study/models"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func DocDeleteBatch() {
	idList := []string{
		"tGcofYkBWS69Op6QHJ2g",
		"tWcpfYkBWS69Op6Q050w",
	}
	bulk := global.ESClient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
	for _, s := range idList {
		req := elastic.NewBulkDeleteRequest().Id(s)
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Succeeded())
}
