/**
 * @Author: 网红电商组
 * @Email: 233_666@qq.com
 * @Date: 2020/7/18 12:33 上午
 * @Description: 商品相关合约路由
 */
package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/Stupid-K/blockchain-real-goods/chaincode/blockchain-real-goods/lib"
	"github.com/Stupid-K/blockchain-real-goods/chaincode/blockchain-real-goods/utils"
	"strconv"
	"time"
)

//新建商品
func CreateRealGoods(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 4 {
		return shim.Error("参数个数不满足")
	}
	accountId := args[0] //accountId用于验证是否为管理员
	Brand := args[1]
	isBrand := args[2]
	isSeller := args[3]
	if accountId == "" || Brand== "" || isBrand == "" || isSeller == "" {
		return shim.Error("参数存在空值")
	}
	if accountId == Brand{
		return shim.Error("操作人应为管理员且与所有人不能相同")
	}
	// 参数数据格式转换
	var formattedisBrand float64
	if val, err := strconv.ParseFloat(isBrand, 64); err != nil {
		return shim.Error(fmt.Sprintf("isBrand参数格式转换出错: %s", err))
	} else {
		formattedisBrand = val
	}
	var formattedisSeller float64
	if val, err := strconv.ParseFloat(isSeller, 64); err != nil {
		return shim.Error(fmt.Sprintf("isSeller参数格式转换出错: %s", err))
	} else {
		formattedisSeller = val
	}
	//判断是否管理员操作
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, []string{accountId})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("操作人权限验证失败%s", err))
	}
	var account lib.Account
	if err = json.Unmarshal(resultsAccount[0], &account); err != nil {
		return shim.Error(fmt.Sprintf("查询操作人信息-反序列化出错: %s", err))
	}
	if account.UserName != "管理员" {
		return shim.Error(fmt.Sprintf("操作人权限不足%s", err))
	}
	//判断品牌是否存在
	resultsProprietor, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, []string{proprietor})
	if err != nil || len(resultsProprietor) != 1 {
		return shim.Error(fmt.Sprintf("业主proprietor信息验证失败%s", err))
	}
	realGoods := &lib.RealGoods{
		RealGoodsID: fmt.Sprintf("%d", time.Now().Local().UnixNano()),
		Proprietor:   proprietor,
		Encumbrance:  false,
		isBrand:    formattedisBrand,
		isSeller:  formattedisSeller,
	}
	// 写入账本
	if err := utils.WriteLedger(realGoods, stub, lib.RealGoodsKey, []string{realGoods.Proprietor, realGoods.RealGoodsID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	realGoodsByte, err := json.Marshal(realGoods)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(realGoodsByte)
}

//查询商品(可查询所有，也可根据所有人查询其下商品)
func QueryRealGoodsList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var realGoodsList []lib.RealGoods
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.RealGoodsKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var realGoods lib.RealGoods
			err := json.Unmarshal(v, &realGoods)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealGoodsList-反序列化出错: %s", err))
			}
			realGoodsList = append(realGoodsList, realGoods)
		}
	}
	realGoodsListByte, err := json.Marshal(realGoodsList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealGoodsList-序列化出错: %s", err))
	}
	return shim.Success(realGoodsListByte)
}
