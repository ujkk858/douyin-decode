package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/opesun/goquery"
)

type HomeController struct {
	Ctx iris.Context
}

func (c *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "home/index.html",
		Data: iris.Map{
			"Title":       "抖音视频解析!",
			"AnalysisUrl": "分析测试",
		},
	}
}

func (c *HomeController) PostAnalysis() mvc.Result {

	var apis = [7]string{
		"https://aweme.snssdk.com/aweme/v1/aweme/detail/?origin_type=link&retry_type=no_retry&iid=74655440239&device_id=57318346369&ac=wifi&channel=wandoujia&aid=1128&app_name=aweme&version_code=140&version_name=1.4.0&device_platform=android&ssmix=a&device_type=MI+8&device_brand=xiaomi&os_api=22&os_version=5.1.1&uuid=865166029463703&openudid=ec6d541a2f7350cd&manifest_version_code=140&resolution=1080*1920&dpi=1080&update_version_code=1400&as=a13520b0e9c40d9cbd&cp=064fdf579fdd07cae1&aweme_id=",
		"https://aweme.snssdk.com/aweme/v1/aweme/detail/?origin_type=link&retry_type=no_retry&iid=74655440239&device_id=57318346369&ac=wifi&channel=wandoujia&aid=1128&app_name=aweme&version_code=140&version_name=1.4.0&device_platform=android&ssmix=a&device_type=MI+8&device_brand=xiaomi&os_api=22&os_version=5.1.1&uuid=865166029463703&openudid=ec6d541a2f7350cd&manifest_version_code=140&resolution=1080*1920&dpi=1080&update_version_code=1400&as=a13510902a54ed1cad&cp=0a40dc5ba5db09cee1&aweme_id=",
		"https://aweme.snssdk.com/aweme/v1/aweme/detail/?origin_type=link&retry_type=no_retry&iid=43619087057&device_id=57318346369&ac=wifi&channel=update&aid=1128&app_name=aweme&version_code=251&version_name=2.5.1&device_platform=android&ssmix=a&device_type=MI+8&device_brand=xiaomi&language=zh&os_api=22&os_version=5.1.1&uuid=865166029463703&openudid=ec6d541a2f7350cd&manifest_version_code=251&resolution=1080*1920&dpi=480&update_version_code=2512&as=a1e500706c54fd8c8d&cp=004ad55fc8d60ac4e1&aweme_id=",
		"https://aweme.snssdk.com/aweme/v1/aweme/detail/?origin_type=link&retry_type=no_retry&iid=43619087057&device_id=57318346369&ac=wifi&channel=update&aid=1128&app_name=aweme&version_code=251&version_name=2.5.1&device_platform=android&ssmix=a&device_type=MI+8&device_brand=xiaomi&language=zh&os_api=22&os_version=5.1.1&uuid=865166029463703&openudid=ec6d541a2f7350cd&manifest_version_code=251&resolution=1080*1920&dpi=480&update_version_code=2512&as=a10500409d74bdec1d&cp=0a4ed456dedf0acee1&aweme_id=",
		"https://aweme.snssdk.com/aweme/v1/aweme/detail/?origin_type=link&retry_type=no_retry&iid=75364831157&device_id=68299559251&ac=wifi&channel=wandoujia&aid=1128&app_name=aweme&version_code=650&version_name=6.5.0&device_platform=android&ssmix=a&device_type=xiaomi+8&device_brand=xiaomi&language=zh&os_api=22&os_version=5.1.1&openudid=2e5c5ff4ce710faf&manifest_version_code=660&resolution=1080*1920&dpi=480&update_version_code=6602&mcc_mnc=46000&js_sdk_version=1.16.2.7&as=a1257080aec45ddcad&cp=0b4cd25fe4d00ccfe1&aweme_id=",
		"https://aweme.snssdk.com/aweme/v1/aweme/detail/?origin_type=link&retry_type=no_retry&iid=75364831157&device_id=68299559251&ac=wifi&channel=wandoujia&aid=1128&app_name=aweme&version_code=650&version_name=6.5.0&device_platform=android&ssmix=a&device_type=xiaomi+8&device_brand=xiaomi&language=zh&os_api=22&os_version=5.1.1&openudid=2e5c5ff4ce710faf&manifest_version_code=660&resolution=1080*1920&dpi=480&update_version_code=6602&mcc_mnc=46000&js_sdk_version=1.16.2.7&as=a125a0b01f946d2cdd&cp=0744d553ffd60cc3e1&aweme_id=",
		"https://aweme.snssdk.com/aweme/v1/aweme/detail/?retry_type=no_retry&iid=74655440239&device_id=57318346369&ac=wifi&channel=wandoujia&aid=1128&app_name=aweme&version_code=140&version_name=1.4.0&device_platform=android&ssmix=a&device_type=MI+8&device_brand=xiaomi&os_api=22&os_version=5.1.1&uuid=865166029463703&openudid=ec6d541a2f7350cd&manifest_version_code=140&resolution=1080*1920&dpi=1080&update_version_code=1400&as=a125372f1c487cb50f&cp=728dcc5bc7f4f558e1&aweme_id=",
	}

	var inputText = c.Ctx.FormValue("AnalysisUrl")
	var url = decodeHttpUrl(inputText)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	var req, err = http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "cookie", Value: "tt_webid=6711334817457341965; _ga=GA1.2.611157811.1562604418; _gid=GA1.2.1578330356.1562604418; _ba=BA0.2-20190709-51", HttpOnly: true})
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 5.0; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var respStr = string(body)

	html, err := goquery.ParseString(respStr)

	if err != nil {
		fmt.Println(err)
	}

	var str = html.Find("script").Text()

	var start = strings.Index(str, "itemId: \"")
	var end = strings.LastIndex(str, "\",\n"+"            test_group")

	var itemId = strings.ReplaceAll(str[start:end], "itemId: \"", "")

	var req1, err1 = http.NewRequest("GET", apis[4]+itemId, nil)
	if err1 != nil {
		fmt.Println(err1)
	}
	req1.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 5.0; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1")

	resp1, err := client.Do(req1)
	if err != nil {
		fmt.Println(err)
	}

	defer resp1.Body.Close()

	body1, err := ioutil.ReadAll(resp1.Body)

	//var jsonResult = string(body1)

	var f interface{}

	json.Unmarshal(body1, &f)

	m := f.(map[string]interface{})["aweme_detail"].(map[string]interface{})["video"].(map[string]interface{})["play_addr"].(map[string]interface{})["url_list"]

	var b, err3 = json.Marshal(m)

	if err3 != nil {
		fmt.Println(err3)
	}

	var arr []string
	json.Unmarshal(b, &arr)

	//var result = fr[2]

	return mvc.View{
		Name: "home/index.html",
		Data: iris.Map{
			"Title":       "抖音视频解析!",
			"AnalysisUrl": inputText,
			"Url":         strings.ReplaceAll(arr[2], "\"", ""),
		},
	}
}

func decodeHttpUrl(url string) string {
	var containChinese = isContainChinese(url)
	if containChinese {
		var start = strings.Index(url, "http")
		var end = strings.LastIndex(url, "/")
		var decodeurl = url[start:end]
		return decodeurl
	} else {
		return url
	}
}

//检测是否包含中文
func isContainChinese(s string) bool {
	var r = []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			return true
		}
	}
	return false
}
