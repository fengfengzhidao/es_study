package models

type UserModel struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	CreatedAt string `json:"created_at"`
	//Age       int    `json:"age"`
	Title string `json:"title"`
}

func (UserModel) Index() string {
	return "user_index"
}

func (UserModel) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "nick_name": { 
        "type": "text"
      },
      "user_name": { 
        "type": "keyword" // 完整匹配
      },
      "age": { 
        "type": "integer" // 完整匹配
      },
      "id": {
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
`
}
