package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/Stupid-K/blockchain-real-goods/chaincode/blockchain-real-goods/lib"
	"testing"
)

func initTest(t *testing.T) *shim.MockStub {
	scc := new(BlockChainRealGoods)
	stub := shim.NewMockStub("ex01", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	return stub
}

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) pb.Response {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
	return res
}

// 测试链码初始化
func TestBlockChainRealGoods_Init(t *testing.T) {
	initTest(t)
}

// 测试获取账户信息
func Test_QueryAccountList(t *testing.T) {
	stub := initTest(t)
	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取多个数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("5feceb66ffc8"),
			[]byte("6b86b273ff34"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取单个数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("4e07408562be"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("4、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("0"),
		}).Payload)))
}

// 测试创建商品
func Test_CreateRealGoods(t *testing.T) {
	stub := initTest(t)
	//成功
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("6b86b273ff34"), //品牌
		[]byte("1"),            //是否品牌认证
		[]byte("1"),            //是否商家认证
	})
	//操作人权限不足
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("6b86b273ff34"), //操作人
		[]byte("4e07408562be"), //品牌
		[]byte("1"),            //是否品牌认证
		[]byte("1"),            //是否商家认证
	})
	//操作人应为管理员且与所有人不能相同
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("5feceb66ffc8"), //品牌
		[]byte("1"),            //是否品牌认证
		[]byte("1"),            //是否商家认证
	})
	//业主proprietor信息验证失败
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"),    //操作人
		[]byte("6b86b273ff34555"), //品牌
		[]byte("1"),            //是否品牌认证
		[]byte("1"),            //是否商家认证
	})
	//参数个数不满足
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("6b86b273ff34"), //品牌
		[]byte("1"),            //是否品牌认证
	})
	//参数格式转换出错
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("6b86b273ff34"), //品牌
		[]byte("1a"),            //是否品牌认证
		[]byte("1"),             //是否商家认证
	})
}

//手动创建一些商品
func checkCreateRealGoods(stub *shim.MockStub, t *testing.T) []lib.RealGoods {
	var realGoodsList []lib.RealGoods
	var realGoods lib.RealGoods
	//成功
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("6b86b273ff34"), //品牌
        []byte("1"),            //是否品牌认证
        []byte("1"),             //是否商家认证
	})
	resp2 := checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("6b86b273ff34"), //品牌
        []byte("1"),            //是否品牌认证
        []byte("1"),             //是否商家认证
	})
	resp3 := checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("4e07408562be"), //品牌
        []byte("1"),            //是否品牌认证
        []byte("0"),             //是否商家认证
	})
	resp4 := checkInvoke(t, stub, [][]byte{
		[]byte("createRealGoods"),
		[]byte("5feceb66ffc8"), //操作人
		[]byte("ef2d127de37b"), //品牌
        []byte("0"),            //是否品牌认证
        []byte("1"),             //是否商家认证
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &realGoods)
	realGoodsList = append(realGoodsList, realGoods)
	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &realGoods)
	realGoodsList = append(realGoodsList, realGoods)
	json.Unmarshal(bytes.NewBuffer(resp3.Payload).Bytes(), &realGoods)
	realGoodsList = append(realGoodsList, realGoods)
	json.Unmarshal(bytes.NewBuffer(resp4.Payload).Bytes(), &realGoods)
	realGoodsList = append(realGoodsList, realGoods)
	return realGoodsList
}

// 测试获取商品信息
func Test_QueryRealGoodsList(t *testing.T) {
	stub := initTest(t)
	realGoodsList := checkCreateRealGoods(stub, t)

	fmt.Println(fmt.Sprintf("1、测试获取所有数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealGoodsList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、测试获取指定数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealGoodsList"),
			[]byte(realGoodsList[0].Proprietor),
			[]byte(realGoodsList[0].RealGoodsID),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、测试获取无效数据\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealGoodsList"),
			[]byte("0"),
		}).Payload)))
}

// 测试发起销售
func Test_CreateSelling(t *testing.T) {
	stub := initTest(t)
	realGoodsList := checkCreateRealGoods(stub, t)
	//成功
	checkInvoke(t, stub, [][]byte{
		[]byte("createSelling"),
		[]byte(realGoodsList[0].RealGoodsID), //销售对象(正在出售的商品RealGoodsID)
		[]byte(realGoodsList[0].Proprietor),   //卖家(卖家AccountId)
		[]byte("50"),                           //价格
		[]byte("30"),                           //智能合约的有效期(单位为天)
	})
	//验证销售对象objectOfSale属于卖家seller失败
	checkInvoke(t, stub, [][]byte{
		[]byte("createSelling"),
		[]byte(realGoodsList[0].RealGoodsID), //销售对象(正在出售的商品RealGoodsID)
		[]byte(realGoodsList[2].Proprietor),   //卖家(卖家AccountId)
		[]byte("50"),                           //价格
		[]byte("30"),                           //智能合约的有效期(单位为天)
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("createSelling"),
		[]byte("123"),                        //销售对象(正在出售的商品RealGoodsID)
		[]byte(realGoodsList[0].Proprietor), //卖家(卖家AccountId)
		[]byte("50"),                         //价格
		[]byte("30"),                         //智能合约的有效期(单位为天)
	})
	//参数错误
	checkInvoke(t, stub, [][]byte{
		[]byte("createSelling"),
		[]byte(realGoodsList[0].RealGoodsID), //销售对象(正在出售的商品RealGoodsID)
		[]byte(realGoodsList[0].Proprietor),   //卖家(卖家AccountId)
		[]byte("50"),                           //价格
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("createSelling"),
		[]byte(""),                           //销售对象(正在出售的商品RealGoodsID)
		[]byte(realGoodsList[0].Proprietor), //卖家(卖家AccountId)
		[]byte("50"),                         //价格
		[]byte("30"),                         //智能合约的有效期(单位为天)
	})
}

// 测试销售发起、购买等操作
func Test_QuerySellingList(t *testing.T) {
	stub := initTest(t)
	realGoodsList := checkCreateRealGoods(stub, t)
	//先发起
	fmt.Println(fmt.Sprintf("发起\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("createSelling"),
		[]byte(realGoodsList[0].RealGoodsID), //销售对象(正在出售的商品RealGoodsID)
		[]byte(realGoodsList[0].Proprietor),   //卖家(卖家AccountId)
		[]byte("500000"),                       //价格
		[]byte("30"),                           //智能合约的有效期(单位为天)
	}).Payload)))
	fmt.Println(fmt.Sprintf("发起\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("createSelling"),
		[]byte(realGoodsList[2].RealGoodsID), //销售对象(正在出售的商品RealGoodsID)
		[]byte(realGoodsList[2].Proprietor),   //卖家(卖家AccountId)
		[]byte("600000"),                       //价格
		[]byte("40"),                           //智能合约的有效期(单位为天)
	}).Payload)))
	//查询成功
	fmt.Println(fmt.Sprintf("1、查询所有\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("querySellingList"),
	}).Payload)))
	fmt.Println(fmt.Sprintf("2、查询指定%s\n%s", realGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("querySellingList"),
		[]byte(realGoodsList[0].Proprietor),
	}).Payload)))
	//购买
	fmt.Println(fmt.Sprintf("3、购买前先查询%s的账户余额\n%s", realGoodsList[2].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realGoodsList[2].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("4、开始购买\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("createSellingByBuy"),
		[]byte(realGoodsList[0].RealGoodsID), //销售对象(正在出售的商品RealGoodsID)
		[]byte(realGoodsList[0].Proprietor),   //卖家(卖家AccountId)
		[]byte(realGoodsList[2].Proprietor),   //买家(买家AccountId)
	}).Payload)))
	fmt.Println(fmt.Sprintf("》购买后再次查询%s的账户余额\n%s", realGoodsList[2].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realGoodsList[2].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》卖家查询购买成功信息\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("querySellingList"),
		[]byte(realGoodsList[0].Proprietor), //买家(买家AccountId)
	}).Payload)))
	fmt.Println(fmt.Sprintf("》买家查询购买成功信息\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("querySellingListByBuyer"),
		[]byte(realGoodsList[2].Proprietor), //买家(买家AccountId)
	}).Payload)))
	fmt.Println(fmt.Sprintf("》确认收款前卖家%s的账户余额\n%s", realGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realGoodsList[0].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》确认收款前买家%s的账户余额\n%s", realGoodsList[2].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realGoodsList[2].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》确认收款前卖家%s的房产信息\n%s", realGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryRealGoodsList"),
		[]byte(realGoodsList[0].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》确认收款前买家%s的房产信息\n%s", realGoodsList[2].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryRealGoodsList"),
		[]byte(realGoodsList[2].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》卖家确认收款\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("updateSelling"),
		[]byte(realGoodsList[0].RealGoodsID), //销售对象(正在出售的商品RealGoodsGoodsID)
		[]byte(realGoodsGoodsList[0].Proprietor),   //卖家(卖家AccountId)
		[]byte(realGoodsGoodsList[2].Proprietor),   //买家(买家AccountId)
		[]byte("done"),                         //确认收款
	}).Payload)))
	//fmt.Println(fmt.Sprintf("》卖家取消收款\n%s", string(checkInvoke(t, stub, [][]byte{
	//	[]byte("updateSelling"),
	//	[]byte(realGoodsGoodsList[0].RealGoodsGoodsID), //销售对象(正在出售的商品RealGoodsGoodsID)
	//	[]byte(realGoodsGoodsList[0].Proprietor),   //卖家(卖家AccountId)
	//	[]byte(realGoodsGoodsList[2].Proprietor),   //买家(买家AccountId)
	//	[]byte("cancelled"),                    //取消收款
	//}).Payload)))
	fmt.Println(fmt.Sprintf("》确认收款后卖家%s的账户余额\n%s", realGoodsGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realGoodsGoodsList[0].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》确认收款后买家%s的账户余额\n%s", realGoodsGoodsList[2].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realGoodsGoodsList[2].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》确认收款后卖家%s的房产信息\n%s", realGoodsGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryRealGoodsGoodsList"),
		[]byte(realGoodsGoodsList[0].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》确认收款后买家%s的房产信息\n%s", realGoodsGoodsList[2].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryRealGoodsGoodsList"),
		[]byte(realGoodsGoodsList[2].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》卖家查询购买成功信息\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("querySellingList"),
		[]byte(realGoodsGoodsList[0].Proprietor), //买家(买家AccountId)
	}).Payload)))
	fmt.Println(fmt.Sprintf("》买家查询购买成功信息\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("querySellingListByBuyer"),
		[]byte(realGoodsGoodsList[2].Proprietor), //买家(买家AccountId)
	}).Payload)))
}

// 测试捐赠合约
func Test_Donating(t *testing.T) {
	stub := initTest(t)
	realGoodsGoodsList := checkCreateRealGoods(stub, t)

	fmt.Println(fmt.Sprintf("获取商品信息\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealGoodsList"),
		}).Payload)))
	//先发起
	fmt.Println(fmt.Sprintf("发起捐赠\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("createDonating"),
		[]byte(realGoodsList[0].RealGoodsID),
		[]byte(realGoodsList[0].Proprietor),
		[]byte(realGoodsList[2].Proprietor),
	}).Payload)))

	fmt.Println(fmt.Sprintf("获取商品信息\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealGoodsList"),
		}).Payload)))

	fmt.Println(fmt.Sprintf("1、查询所有\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("queryDonatingList"),
	}).Payload)))
	fmt.Println(fmt.Sprintf("2、查询指定%s\n%s", realGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryDonatingList"),
		[]byte(realGoodsList[2].Proprietor),
	}).Payload)))
	fmt.Println(fmt.Sprintf("3、查询指定受赠%s\n%s", realGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryDonatingListByGrantee"),
		[]byte(realGoodsList[2].Proprietor),
	}).Payload)))

	//fmt.Println(fmt.Sprintf("4、接受受赠%s\n%s", realGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
	//	[]byte("updateDonating"),
	//	[]byte(realGoodsList[0].RealGoodsID),
	//	[]byte(realGoodsList[0].Proprietor),
	//	[]byte(realGoodsList[2].Proprietor),
	//	[]byte("done"),
	//}).Payload)))
	fmt.Println(fmt.Sprintf("4、取消受赠%s\n%s", realGoodsList[0].Proprietor, string(checkInvoke(t, stub, [][]byte{
		[]byte("updateDonating"),
		[]byte(realGoodsList[0].RealGoodsID),
		[]byte(realGoodsList[0].Proprietor),
		[]byte(realGoodsList[2].Proprietor),
		[]byte("cancelled"),
	}).Payload)))

	fmt.Println(fmt.Sprintf("获取商品信息\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealGoodsList"),
		}).Payload)))
}
