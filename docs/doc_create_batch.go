package docs

import (
	"context"
	"es_study/global"
	"es_study/models"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

func DocCreateBatch() {

	list := []models.UserModel{
		{
			//ID:        13,
			//UserName:  "lisi",
			//NickName:  "夜空中最亮的李四",
			Title:     "这是我的生活",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			//ID:        14,
			//UserName:  "zhangsan",
			//NickName:  "张三",
			Title:     "你好啊，枫枫",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			//ID:        14,
			//UserName:  "zhangsan",
			//NickName:  "张三",
			Title:     "这是我的枫枫",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	bulk := global.ESClient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
	for _, model := range list {
		req := elastic.NewBulkCreateRequest().Doc(model)
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Succeeded())
}
