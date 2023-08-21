package order

import "gorm.io/gorm"

type Order struct {
	ID               uint    `json:"id" gorm:"primary_key;auto_increment;not null;comment:'主键 - 自增id'"`
	Type             int     `json:"type" gorm:"column:type;comment:'1.空投2.寄售'"`
	CreatedAt        string  `json:"createdAt" gorm:"column:createdAt;type:datetime(0);comment:'创建时间'"`
	UpdatedAt        string  `json:"updatedAt" gorm:"column:updatedAt:type:datetime(0);comment:'更新时间'"`
	OrderNo          string  `json:"orderNo" gorm:"column:orderNo;comment:'订单编号'"`
	OrderName        string  `json:"orderName" gorm:"column:orderName;comment:'订单名称'"`
	OrderNumber      int     `json:"orderNumber" gorm:"orderNumber;comment:'订单数量'"`
	OrderImg         string  `json:"orderImg" gorm:"orderImg;comment:'图片'"`
	OrderPrice       float32 `json:"orderPrice" gorm:"orderPrice;comment:'订单价格'"`
	OrderCouponId    int32   `json:"orderCouponId" gorm:"orderCouponId;comment:'优惠券ID'"`
	OrderCouponPrice string  `json:"orderCouponPrice" gorm:"orderCouponPrice;comment:'优惠券价格'"`
	OrderCouponType  int     `json:"orderCouponType" gorm:"orderCouponType;comment:'优惠券类型1.折扣2.免单'"`
	OrderCouponDesc  string  `json:"orderCouponDesc" gorm:"orderCouponDesc;comment:'优惠券说明'"`
	OrderStatus      int     `json:"orderStatus" gorm:"orderStatus;comment:'订单状态0.取消1.支付成功2.支付中3.挂售中4.卖方已取消5.买方已取消'"`
	WalletAddress    string  `json:"walletAddress" gorm:"walletAddress;comment:'钱包地址'"`
	PayType          string  `json:"payType" gorm:"payType;comment:'支付类型1.'"`
	PayName          string  `json:"payName" gorm:"payName;comment:'支付类型名称'"`
	PayDesc          string  `json:"payDesc" gorm:"payDesc;comment:'支付备注'"`
	SendId           string  `json:"sendId" gorm:"sendId;comment:'支付方id'"`
	SendName         string  `json:"sendName" gorm:"sendName;comment:'支付方名称'"`
	SendAddress      string  `json:"sendAddress" gorm:"sendAddress;comment:'支付方钱包地址'"`
	ReceiveId        string  `json:"receiveId" gorm:"receiveId;comment:'接收方地址'"`
	ReceiveName      string  `json:"receiveName" gorm:"receiveName;comment:'接收方名称'"`
	ReceiveAddress   string  `json:"receiveAddress" gorm:"receiveAddress;comment:'接收方地址'"`
}

func OrderDbinit(db *gorm.DB) {
	err := db.AutoMigrate(Order{})
	if err != nil {
		return
	}
}
