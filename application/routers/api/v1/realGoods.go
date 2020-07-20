/**
 * @Author: 网红电商组
 * @Email: 233_666@qq.com
 * @Date: 2020/7/19 4:43 下午
 * @Description: 商品信息相关接口
 */
package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/Stupid-K/blockchain-real-goods/application/blockchain"
	"github.com/Stupid-K/blockchain-real-goods/application/pkg/app"
	"net/http"
	"strconv"
)

type RealGoodsRequestBody struct {
	AccountId   string  `json:"accountId"`     //操作人ID
	Brand       string  `json:"brand"`          //品牌所有(品牌的AccountId)
	isBrand     float64 `json:"isBrand"`       //是否品牌保证
	isSeller    float64 `json:"isSeller"`      //是否卖家认证
}

type RealGoodsQueryRequestBody struct {
	Brand      string  `json:"brand"`  //卖家(网红)(网红AccountId)
}

// @Summary 新建商品
// @Param realGoods body RealGoodsRequestBody true "realGoods"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/createRealGoods [post]
func CreateRealGoods(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealGoodsRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.AccountId))
	bodyBytes = append(bodyBytes, []byte(body.Brand))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.isBrand, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.isSeller, 'E', -1, 64)))
	//调用智能合约
	resp, err := bc.ChannelExecute("createRealGoods", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 获取商品信息(空json{}可以查询所有，指定brand可以查询指定品牌旗下的商品)
// @Param realGoodsQuery body RealGoodsQueryRequestBody true "realGoodsQuery"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryRealGoodsList [post]
func QueryRealGoodsList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealGoodsQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Brand != "" {
		bodyBytes = append(bodyBytes, []byte(body.Brand))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryRealGoodsList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
