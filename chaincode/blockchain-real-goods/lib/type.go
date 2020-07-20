/**
 * @Author: 网红电商组
 * @Email: 233_666@qq.com
 * @Date: 2020/7/20 2:00 下午
 * @Description: 定义的数据结构体、常量
 */
package lib

//账户，虚拟管理员和若干业主账号
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名
	Balance   float64 `json:"balance"`   //余额
}

//商品的状态不对（盗版、伪证。。。）时Encumbrance为true，默认状态false。
//仅当Encumbrance为false时，才可发起出售
//Brand和RealGoodsID一起作为复合键,保证可以通过Brand查询到此品牌的所有产品
type RealGoods struct {
	RealGoodsID string  `json:"realGoodsId"` //房地产ID
	Brand       string  `json:"brand"`       //品牌所有(品牌的AccountId)
	Encumbrance  bool    `json:"encumbrance"`  //处在良好状态
	isBrand     int     `json:"isBrand"`       //是否品牌保证
	isSeller    int     `json:"isSeller"`      //是否卖家认证
}

//销售要约
//需要确定ObjectOfSale是否属于Seller
//买家初始为空
//Seller和ObjectOfSale一起作为复合键,保证可以通过seller查询到名下所有发起的销售
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //销售对象(正在出售的商品RealGoodsID)
	Seller        string  `json:"seller"`        //发起推荐的网红(卖家AccountId)
	Buyer         string  `json:"buyer"`         //参与销售人、买家(买家AccountId)
	Price         float64 `json:"price"`         //价格
	CreateTime    string  `json:"createTime"`    //创建时间
	SalePeriod    int     `json:"salePeriod"`    //智能合约的有效期(单位为天)
	SellingStatus string  `json:"sellingStatus"` //销售状态
}

//销售状态
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "销售中", //正在销售状态,等待买家光顾
		"cancelled": "已取消", //被卖家取消销售或买家退款操作导致取消
		"expired":   "已过期", //销售期限到期
		"delivery":  "交付中", //买家买下并付款,处于等待卖家确认收款状态,如若卖家未能确认收款，买家可以取消并退款
		"done":      "完成",  //卖家确认接收资金，交易完成
	}
}

//买家参与销售
//销售对象不能是买家发起的
//Buyer和CreateTime作为复合键,保证可以通过buyer查询到名下所有参与的销售
type SellingBuy struct {
	Buyer      string  `json:"buyer"`      //参与销售人、买家(买家AccountId)
	CreateTime string  `json:"createTime"` //创建时间
	Selling    Selling `json:"selling"`    //销售对象
}

const (
	AccountKey         = "account-key"
	RealGoodsKey      = "real-Goods-key"
	SellingKey         = "selling-key"
	SellingBuyKey      = "selling-buy-key"
)
