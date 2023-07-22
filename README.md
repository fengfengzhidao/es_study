本节课讲解es的基本操作

ElasticSearch 简称es，开源的分布式的`全文搜索引擎`，可以近乎实时的存储检索数据

## es安装

建议大家直接使用docker安装es



拉取镜像

```Python
docker pull elasticsearch:7.12.0
```



创建docker容器挂在的目录：

```Python
# linux的命令
mkdir -p /opt/es/config & mkdir -p /opt/es/data & mkdir -p /opt/es/plugins

chmod 777 /opt/es/data

```

配置文件

```Python
echo "http.host: 0.0.0.0" > /opt/es/config/elasticsearch.yml
```



创建容器

```Python
# linux
docker run --name es -p 9200:9200  -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -v /opt/es/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/es/data:/usr/share/elasticsearch/data -v /opt/es/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0

# windows
docker run --name es -p 9200:9200  -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -v H:\\docker\\es\\config\\elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v H:\\docker\\es\\data:/usr/share/elasticsearch/data -v H:\\docker\\es\\plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0
# windows添加目录映射，需要在dockerDesktop里面设置映射目录

```









访问ip:9200能看到东西

![](http://python.fengfengzhidao.com/pic/20230129212040.png)

就说明安装成功了



浏览器可以下载一个 `Multi Elasticsearch Head` es插件



或者这个

ElasticHD

[https://github.com/qax-os/ElasticHD/releases](https://github.com/qax-os/ElasticHD/releases)



## 连接es

```go
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
    elastic.SetBasicAuth("", ""),
  )
  if err != nil {
    fmt.Println(err)
    return
  }
  global.ESClient = client
}

```



## es认证



不需要认证的情况

1. 服务器自己使用，9200，9300端口不对外开放
2. 本身跑在127.0.0.1上

需要认证的情况：

1. es需要对外提供服务的

[https://blog.csdn.net/qq_38669698/article/details/130529829](https://blog.csdn.net/qq_38669698/article/details/130529829)



![](http://python.fengfengzhidao.com/pic/20230722173432.png)



这样就说明成功了

输入用户名和密码就能看到之前的那个页面

或者使用curl进行测试

```go
curl  http://127.0.0.1:9200/
curl -u elastic:xxxxxx http://127.0.0.1:9200/
```



## 索引操作



### mapping

```go
// 查看某个索引的map
/index/_mapping
```

常见的类型

```go
{
  "mappings": {
    "properties": {
      "title": { 
        "type": "text" // 查询的时候是分词匹配
      },
      "key": { 
        "type": "keyword" // 完整匹配
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
```



### 创建索引

```go
func CreateIndex() {
  createIndex, err := global.ESClient.
    CreateIndex("user_index").
    BodyString(models.UserModel{}.Mapping()).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(createIndex)
  fmt.Println("索引创建成功")
}

```

> 注意：索引存在执行创建索引是会报错的



### 判断索引是否存在

```go
// ExistsIndex 判断索引是否存在
func ExistsIndex(index string) bool {
  exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
  return exists
}

```



### 删除索引

```go
func DeleteIndex(index string) {
  _, err := global.ESClient.
    DeleteIndex(index).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(index, "索引删除成功")
}

```





## 文档操作

### 添加文档

#### 添加单个文档

```go
func DocCreate() {
  user := models.UserModel{
    ID:        12,
    UserName:  "lisi",
    Age:       23,
    NickName:  "夜空中最亮的lisi",
    CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
    Title:     "今天天气很不错",
  }
  indexResponse, err := global.ESClient.Index().Index(user.Index()).BodyJson(user).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("%#v\n", indexResponse)
}

```

添加文档

如果是mapping里面没有的字段，那么es会自动创建这个字段对应的mapping



#### 批量添加

```go
func DocCreateBatch() {

  list := []models.UserModel{
    {
      ID:        12,
      UserName:  "fengfeng",
      NickName:  "夜空中最亮的枫枫",
      CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
    },
    {
      ID:        13,
      UserName:  "lisa",
      NickName:  "夜空中最亮的丽萨",
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
```



### 删除文档

#### 根据id删除

```go
func DocDelete() {

  deleteResponse, err := global.ESClient.Delete().
    Index(models.UserModel{}.Index()).Id("tmcqfYkBWS69Op6Q4Z0t").Refresh("true").Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(deleteResponse)
}

```

如果文档不存在，会报404的错误



#### 根据id批量删除

```go
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
  fmt.Println(res.Succeeded())  // 实际删除的文档切片
}

```

如果文档不存在，不会有错误， `res.Succeeded()` 为空



### 文档查询



#### 列表查询

```go
func DocFind() {

  limit := 2
  page := 4
  from := (page - 1) * limit

  query := elastic.NewBoolQuery()
  res, err := global.ESClient.Search(models.UserModel{}.Index()).Query(query).From(from).Size(limit).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  count := res.Hits.TotalHits.Value  // 总数
  fmt.Println(count)
  for _, hit := range res.Hits.Hits {
    fmt.Println(string(hit.Source))
  }
}

```

#### 精确匹配

针对keyword字段

```go
query := elastic.NewTermQuery("user_name", "fengfeng")
```

#### 模糊匹配

主要是查text，也能查keyword

模糊匹配keyword字段，是需要查完整的

匹配text字段则不用，搜完整的也会搜出很多

```go
query := elastic.NewMatchQuery("nick_name", "夜空中最亮的枫枫")
```





#### 嵌套字段的搜索

```go
"title": {
    "type": "text",
    "fields": {
        "keyword": {
            "type": "keyword",
            "ignore_above": 256
        }
    }
},
```

因为title是text类型，只能模糊匹配，但是需要精确匹配的时候，也能通过title.keyword的形式进行精确匹配



```go
query := elastic.NewTermQuery("title.keyword", "这是我的枫枫") // 精确匹配
//query := elastic.NewMatchQuery("title", "这是我的枫枫")  // 模糊匹配
```





### 文档更新

```go
func DocUpdate() {
  res, err := global.ESClient.Update().Index(models.UserModel{}.Index()).
    Id("vmdnfYkBWS69Op6QEp2Y").Doc(map[string]any{
    "user_name": "你好枫枫",
  }).Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("%#v\n", res)
}

```