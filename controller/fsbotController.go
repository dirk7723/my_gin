package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FsbotController struct{}

type Card struct {
	Title      string `form:"title" json:"title"`
	Content    string `form:"content" json:"content"`
	Link       string `form:"link" json:"link"`
	Product_id int    `form:"product_id" json:"product_id"`
}

type Ret map[string]interface{}

//type Result map[string]interface{}

func (con FsbotController) Index(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var ca Card
	err := c.ShouldBind(&ca)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	} else {
		product_id := ca.Product_id
		bot := getBot(product_id)
		bot_img := bot["bot_img"].(string)
		post_data_origin := buildCard(ca, bot_img)
		post_data, _ := json.Marshal(post_data_origin)
		var data = bytes.NewReader(post_data)
		client := &http.Client{}
		url := bot["bot_url"].(string)
		req, err := http.NewRequest("POST", url, data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		// bodyText, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%s\n", bodyText)
		var ret Ret
		json.NewDecoder(resp.Body).Decode(&ret)
		var result Ret
		result = make(Ret)
		if StatusCode, ok := ret["StatusCode"].(float64); ok {
			if StatusCode == 0 {
				result["code"] = 0
				result["msg"] = "success"
			} else {
				result["code"] = 1006
				result["msg"] = "发送失败"
			}
		} else {
			result["code"] = 1007
			result["msg"] = "未知错误，联系管理员:" + ret["msg"].(string)
		}
		c.JSON(http.StatusOK, result)
	}
}

func buildCard(card Card, img_key string) map[string]interface{} {
	/* 声明局部变量 */
	post_data := map[string]interface{}{
		"msg_type": "interactive",
		"card": map[string]interface{}{
			"config": map[string]interface{}{
				"wide_screen_mode": true,
			},
			"header": map[string]interface{}{
				"template": "green",
				"title": map[string]interface{}{
					"content": "🔔" + card.Title,
					"tag":     "plain_text",
				},
			},
			"elements": []map[string]interface{}{
				{
					"alt": map[string]interface{}{
						"content": "",
						"tag":     "plain_text",
					},
					"img_key": img_key,
					"tag":     "img",
				},
				{
					"tag": "div",
					"text": map[string]interface{}{
						"content": card.Content,
						"tag":     "lark_md",
					},
				},
				{
					"actions": []map[string]interface{}{
						{
							"tag": "button",
							"text": map[string]interface{}{
								"content": "查看详情",
								"tag":     "plain_text",
							},
							"type": "primary",
							"url":  card.Link,
						},
					},
					"tag": "action",
				},
			},
		},
	}
	return post_data
}

func getBot(product_id int) map[string]interface{} {
	//dev
	Feishu_bot := map[string]interface{}{
		"Zeus":   "https://open.feishu.cn/open-apis/bot/v2/hook/3ebc3e66-1221-44a1-9018-ceeb91b0a771",
		"Nowsee": "https://open.feishu.cn/open-apis/bot/v2/hook/bbff02c0-d331-4dbc-b9c0-a2887e6b3517",
		"Seers":  "https://open.feishu.cn/open-apis/bot/v2/hook/577bfd80-82c8-490a-9bd4-fea9142d9250",
	}
	//prod
	// Feishu_bot := map[string]interface{}{
	// 	"Zeus":   "https://open.feishu.cn/open-apis/bot/v2/hook/474bf8ef-a4e7-4949-99c9-b016e32203bd",
	// 	"Nowsee": "https://open.feishu.cn/open-apis/bot/v2/hook/4c1dc39d-571c-46da-85b5-4027c5171f19",
	// 	"Seers":  "https://open.feishu.cn/open-apis/bot/v2/hook/4f31aeca-b118-4843-b1f9-b1e2c2e8c174",
	// }

	bot := map[string]interface{}{
		"bot_url": Feishu_bot["Zeus"],
		"bot_img": "img_v2_dbdc1379-039a-487a-b92e-7c09d0dd0d1g",
	}
	switch product_id {
	case 1:
		bot["bot_url"] = Feishu_bot["Nowsee"]
		bot["bot_img"] = "img_v2_02a08ae8-d38f-415e-a8a9-5e2c7f93b93g"
	case 2:
		bot["bot_url"] = Feishu_bot["Nowsee"]
		bot["bot_img"] = "img_v2_02a08ae8-d38f-415e-a8a9-5e2c7f93b93g"
	case 3:
		bot["bot_url"] = Feishu_bot["Seers"]
		bot["bot_img"] = "img_v2_fd441874-3c8b-418e-a21c-3601d6c0033g"
	default:
		bot["bot_url"] = Feishu_bot["Zeus"]
		bot["bot_img"] = "img_v2_dbdc1379-039a-487a-b92e-7c09d0dd0d1g"
	}
	return bot
}
