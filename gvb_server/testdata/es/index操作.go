package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
)

// Assuming 'client' is initialized elsewhere in your code like this:
// client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
//
//	if err != nil {
//	    logrus.Fatal(err)
//	}
var client *elastic.Client

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	return c
}

func init() {
	core.InitConf()
	core.InitLogger()
	client = EsConnect()
}

type DemoModel struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func (DemoModel) Index() string {
	return "demo_index"
}

func (DemoModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": 100000
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "yyyy-MM-dd HH:mm:ss||strict_date_optional_time||epoch_millis"
      }
    }
  }
}
`
}

func (demo DemoModel) IndexExists() bool {
	exists, err := client.IndexExists(demo.Index()).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

func (demo DemoModel) CreateIndex() error {
	if demo.IndexExists() {
		if err := demo.RemoveIndex(); err != nil {
			return err
		}
	}
	createIndex, err := client.CreateIndex(demo.Index()).BodyString(demo.Mapping()).Do(context.Background())
	if err != nil {
		logrus.Error("Fail to create the index: ", err)
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("Fail to acknowledge the index")
		return err
	}
	logrus.Infof("Index %s creates successfully", demo.Index())
	return nil
}

func (demo DemoModel) RemoveIndex() error {
	logrus.Info("index existed，is going to delete...")
	// context.Background() 提供了一个没有特定超时或取消策略的上下文，用于执行删除索引的操作。
	// 这意味着操作将不会因为超时或被取消而中断，除非 Elasticsearch 服务本身响应了超时或错误。
	indexDelete, err := client.DeleteIndex(demo.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("fail to delete the index: ", err)
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("fail to acknowledge the deleting index")
		return err
	}
	logrus.Info("index deleted successfully")
	return nil
}

/*func Create(data *DemoModel) error {
	// Create 函数将一个 DemoModel 实例的数据索引到 Elasticsearch，
	// 并更新这个实例的 ID 字段为 Elasticsearch 为新索引的文档生成的唯一标识符。
	indexResponse, err := client.Index().Index(data.Index()).BodyJson(data).Do(context.Background())
	// 这里指定了我们想要保存数据的“索引”名称。在 Elasticsearch 里，
	// “索引”就像是传统数据库中的“表”，是保存相关数据的地方。
	// BodyJson 方法告诉客户端，我们将数据以 JSON 格式提供。
	// 这里的 data 就是我们想要保存到 Elasticsearch 中的数据。
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	logrus.Infof("%#v", indexResponse)
	data.ID = indexResponse.Id
	return nil
}*/

func Create(data *DemoModel) (err error) {
	indexResponse, err := client.Index().
		Index(data.Index()).
		BodyJson(data).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	data.ID = indexResponse.Id
	return nil
}

// FindList 列表查询
func FindList(key string, page, limit int) (demoList []DemoModel, count int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := client.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	return demoList, count
}

func FindSourceList(key string, page, limit int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := client.
		Search(DemoModel{}.Index()).
		Source(`{"_source": ["title"]}`).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count := int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []DemoModel{}
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	fmt.Println(demoList, count)
}

func Update(id string, data *DemoModel) error {
	_, err := client.
		Update().
		Index(DemoModel{}.Index()).
		Id(id).
		Doc(map[string]string{
			"title": data.Title,
		}).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	logrus.Info("update demo successfully")
	return nil
}

func Remove(idList []string) (count int, err error) {
	bulkService := client.Bulk().Index(DemoModel{}.Index()).Refresh("true")
	for _, id := range idList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	res, err := bulkService.Do(context.Background())
	return len(res.Succeeded()), err
}

func main() {
	/*DemoModel{}.CreateIndex()*/
	/*Create(&DemoModel{Title: "go语言开发", UserID: 2, CreatedAt: time.Now().Format(time.RFC3339)})*/
	/*list, count := FindList("python", 1, 10)
	fmt.Println(list, count)*/
	/*Update("m1pkLI4B2vl9wLVbigSu", &DemoModel{Title: "python零基础入门"})*/
	// count, err := Remove([]string{"m1pkLI4B2vl9wLVbigSu"})
	// fmt.Println(count, err)
}
