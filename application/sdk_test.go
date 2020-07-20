/**
 * @Author: 网红电商组
 * @Email: 233_666@qq.com
 * @Date: 2020/7/18 5:14 下午
 * @Description: sdk的测试
 */
package main_test

import (
	"fmt"
	"github.com/Stupid-K/blockchain-real-goods/application/blockchain"
	"testing"
)

func TestInvoke_QueryAccountList(t *testing.T) {
	blockchain.Init()
	response, e := blockchain.ChannelQuery("queryAccountList", [][]byte{})
	if e != nil {
		fmt.Println(e.Error())
		t.FailNow()
	}
	fmt.Println(string(response.Payload))
}
