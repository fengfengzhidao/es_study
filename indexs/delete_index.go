package indexs

import (
	"context"
	"es_study/global"
	"fmt"
)

func DeleteIndex(index string) {
	_, err := global.ESClient.
		DeleteIndex(index).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(index, "索引删除成功")
}
