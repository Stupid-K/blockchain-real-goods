/**
 * @Author: 网红电商组
 * @Email: 233_666@qq.com
 * @Date: 2020/7/19 4:18 下午
 * @Description: 账户相关合约路由
 */
package routers

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/Stupid-K/blockchain-real-goods/chaincode/blockchain-real-estate/lib"
	"github.com/Stupid-K/blockchain-real-goods/chaincode/blockchain-real-estate/utils"
)

//查询账户列表
func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []lib.Account
	results, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var account lib.Account
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}
