package docs

import (
	"context"
	"es_study/global"
	"es_study/models"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func DocFind() {

	limit := 10
	page := 1
	from := (page - 1) * limit

	query := elastic.NewTermQuery("title.keyword", "这是我的枫枫")
	//query := elastic.NewMatchQuery("title", "夜空中最亮的枫枫")

	res, err := global.ESClient.Search(models.UserModel{}.Index()).
		Query(query).
		From(from).Size(limit).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	count := res.Hits.TotalHits.Value
	fmt.Println(count)
	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}
