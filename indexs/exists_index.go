package indexs

import (
	"context"
	"es_study/global"
)

// ExistsIndex 判断索引是否存在
func ExistsIndex(index string) bool {
	exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
	return exists
}
