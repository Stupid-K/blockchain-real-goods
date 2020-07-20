package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/Stupid-K/blockchain-real-goods/chaincode/blockchain-real-goods/lib"
	"github.com/Stupid-K/blockchain-real-goods/chaincode/blockchain-real-goods/routers"
	"github.com/Stupid-K/blockchain-real-goods/chaincode/blockchain-real-goods/utils"
	"time"
)

type BlockChainRealGoods struct {
}

//链码初始化
func (t *BlockChainRealGoods) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	var accountIds = [6]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
	}
	var userNames = [6]string{"管理员", "①号顾客", "②号顾客", "③号顾客", "④号顾客", "⑤号顾客"}
	var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}
	//初始化账号数据
	for i, val := range accountIds {
		account := &lib.Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		// 写入账本
		if err := utils.WriteLedger(account, stub, lib.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

//实现Invoke接口调用智能合约
func (t *BlockChainRealGoods) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "queryAccountList":
		return routers.QueryAccountList(stub, args)
	case "createRealGoods":
		return routers.CreateRealGoods(stub, args)
	case "queryRealGoodsList":
		return routers.QueryRealGoodsList(stub, args)
	case "createSelling":
		return routers.CreateSelling(stub, args)
	case "createSellingByBuy":
		return routers.CreateSellingByBuy(stub, args)
	case "querySellingList":
		return routers.QuerySellingList(stub, args)
	case "querySellingListByBuyer":
		return routers.QuerySellingListByBuyer(stub, args)
	case "updateSelling":
		return routers.UpdateSelling(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(BlockChainRealGoods))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
