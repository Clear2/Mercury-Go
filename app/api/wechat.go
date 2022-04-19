package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mercury/app/utils"
)

//https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx7bd870672e9700b2&secret=3bcc18dbf19f3a7384504275f54eb7d2
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessToken(ctx *gin.Context) {
	body, _ := utils.HTTPGet("", nil)
	resAccessToken := new(AccessToken)
	err := json.Unmarshal(body, &resAccessToken)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(resAccessToken.AccessToken)
}
