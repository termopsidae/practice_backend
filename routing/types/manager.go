package types

// LoginReq 登录注册参数
type LoginReq struct {
	ManagerName string `json:"manager_name"` //管理员用户名
	Password    string `json:"password"`     //密码
}
type LoginResp struct {
	Token string `json:"token"` // token
	Class string `json:"class"` // 管理员等级
}

// SelectManagerListReq   查询资源包列表参数
type SelectManagerListReq struct {
}
type SelectManagerListResp struct {
	ManagerList []ManagerInfo `json:"Manager_list"`
}
type ManagerInfo struct {
	ManagerId   uint   `json:"manager_id"`   //管理员id
	ManagerName string `json:"manager_name"` //管理员用户名
	Class       string `json:"class"`        //等级 1-主要 2-普通
	Flag        string `json:"flag"`         // 激活标志("2"- 空投用"1"-已生效 "0"-未生效)
}

// CreateNewManagerReq 创建新管理员
type CreateNewManagerReq struct {
	ManagerName string `json:"manager_name"` //管理员用户名
	Password    string `json:"password"`     //密码
}
type CreateNewManagerResp struct {
	Token string `json:"token"` //token
}

// UpdateManagerReq 调整管理员账户
type UpdateManagerReq struct {
	ManagerId   uint   `json:"manager_id"`   //管理员id
	ManagerName string `json:"manager_name"` //管理员用户名
	Password    string `json:"password"`     //密码
	Class       string `json:"class"`        //等级 1-主要 2-普通
	Flag        string `json:"flag"`         //标识 1-启用 0-停用
}
type UpdateManagerResp struct {
}

// AirdropPackageReq 空投资源包
type AirdropPackageReq struct {
	PackageId   uint     `json:"package_id"`   //资源包id
	AddressList []string `json:"address_list"` //地址数组
}
type AirdropPackageResp struct {
}

// PlacePackageReq  配置资源包
type PlacePackageReq struct {
	PackageName string `json:"package_name"` //资源包名字
	Price       string `json:"price"`        //价格
	Value       string `json:"value"`        //价值
	Limit       int64  `json:"limit"`        //限量
	Flag        string `json:"flag"`         // 激活标志("2"- 空投用"1"-已生效 "0"-未生效)
}

type PlacePackageResp struct {
}

// UpdatePackageInfoReq  挂单参数
type UpdatePackageInfoReq struct {
	PackageId   uint   `json:"package_id"`   //资源包id
	PackageName string `json:"package_name"` //资源包名字
	Price       string `json:"price"`        //价格
	Limit       int64  `json:"limit"`        //限量
	Value       string `json:"value"`        //价值
	Flag        string `json:"flag"`         // 激活标志("2"- 空投用"1"-已生效 "0"-未生效)=
}
type UpdatePackageInfoResp struct {
}

// SelectUserReq 展示用户信息
type SelectUserReq struct {
	WalletAddress      string `json:"wallet_address"`      //地址
	RecommenderAddress string `json:"recommender_address"` //推荐人地址
	Level              int64  `json:"level"`               //用户等级
	Flag               string `json:"flag"`                // 激活标志(1"-已激活 "0"-未激活)
	OrderBy            string `json:"order_by"`            // 排序 "usdt_balance" "a_balance" "b_balance" "locked_balance" "" - 不排序
	Page               int    `json:"page"`                //分页
}
type SelectUserResp struct {
	UserInfoList []UserInfo `json:"user_info_list"`
	PageSize     int        `json:"page_size"` //分
}

// UpdateUserInfoReq  更新用户
type UpdateUserInfoReq struct {
	WalletAddress string `json:"wallet_address"`
	Level         int64  `json:"level"`
	Signal        string `json:"signal"` //1-普通 2-白名单
	Frozen        string `json:"frozen"` //1-冻结 0-正常
}
type UpdateUserInfoResp struct {
}

// ShowSignLogReq 签到日志
type ShowSignLogReq struct {
	WalletAddress string `json:"wallet_address"`
}
type ShowSignLogResp struct {
	Dates []string `json:"dates"`
}

// SelectCurrenciesReq 展示币价信息
type SelectCurrenciesReq struct {
}
type SelectCurrenciesResp struct {
	CurrencyInfoList []CurrencyInfo `json:"currency_info_list"`
}
type CurrencyInfo struct {
	CurrencyId uint   `json:"currency_id"`
	CoinType   string `json:"coin_type"`
	Value      string `json:"value"`
}

// SelectTransactionsReq 查询交易信息
type SelectTransactionsReq struct {
	WalletAddress string  `json:"wallet_address"` //"" - 不限制用户
	TimeDuration  []int64 `json:"time_duration"`  //起始终止时间戳的int64数组
	ChangeType    string  `json:"change_type"`
	Page          int     `json:"page"`
}

// UpdateCurrencyReq  更新币价
type UpdateCurrencyReq struct {
	CurrencyId uint   `json:"currency_id"`
	Value      string `json:"value"`
}
type UpdateCurrencyResp struct {
}

// UpdateKeyValueReq  更新key-value
type UpdateKeyValueReq struct {
	KeyValueInfoArray []KeyValueInfo `json:"key_value_info_array"`
}
type UpdateKeyValueResp struct {
}

// SelectKeyValuesReq 展示币价信息
type SelectKeyValuesReq struct {
}
type SelectKeyValuesResp struct {
	DynamicList      []KeyValueInfo `json:"dynamic_list"`
	PerformanceList  []KeyValueInfo `json:"performance_list"`
	GroupList        []KeyValueInfo `json:"group_list"`
	BonusList        []KeyValueInfo `json:"bonus_list"`
	KeyValueInfoList []KeyValueInfo `json:"key_value_info_list"`
}
type KeyValueInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UpdateMiningReq struct {
	MiningId      uint   `json:"mining_id"`
	Price         string `json:"price"`         // 价格 bi
	Magnification string `json:"magnification"` // 倍率
	BBenefit      string `json:"b_benefit"`     // 空投AIGT数量 bi
	Profit        string `json:"profit"`        // 收益率
	Flag          string `json:"flag"`          // 0-停用 1-可合成 2-可购买 3-可合成可购买
}
type UpdateMiningResp struct {
}

type UpdateMiningOrderReq struct {
	ID   uint   `json:"id"`
	Flag string `json:"flag"` // 0-停用 1-启用 2-不产静态
}

type UpdateMiningOrderResp struct {
}

type UpdateNodeReq struct {
	NodeId   uint   `json:"node_id"`
	Price    string `json:"price"`     // 价格 bi
	Level    int64  `json:"level"`     // 等级
	Last     string `json:"last"`      // 剩余数量
	BBenefit string `json:"b_benefit"` // 空投AIGT数量 bi
	MiningID uint   `json:"mining_id"` // 赠送矿机id
	Flag     string `json:"flag"`      // 0 停用 1 启用
}
type UpdateNodeResp struct {
}
type UpdateLevelReq struct {
	WalletAddress string `json:"wallet_address"`
	Level         int64  `json:"level"` // 等级
}
type UpdateLevelResp struct {
}
type UpdateUserBalanceReq struct {
	WalletAddress string `json:"wallet_address"`
	Amount        string `json:"amount"`
	Hash          string `json:"hash"`
	Flag          string `json:"flag"` // 0 扣减 1 增加
}
type UpdateUserBalanceResp struct {
}

type UpdateUserFlagReq struct {
	WalletAddress string `json:"wallet_address"`
	Flag          string `json:"flag"` // 0 正常 -1 冻结
}
type UpdateUserFlagResp struct {
}

type ReviewWithdrawReq struct {
	WalletAddress string `json:"wallet_address"`
	ID            uint   `json:"id"`
	Flag          string `json:"flag"` // 2-审核通过 -1-审核驳回
}
type ReviewWithdrawResp struct {
}

// DailyReportListReq 每日新增报表
type DailyReportListReq struct {
}
type DailyReportListResp struct {
	USDTRecord USDTRecord `json:"usdt_record"` //usdt相关统计
	AIGPRecord AIGPRecord `json:"aigp_record"` //aigp相关统计
	AIGTRecord AIGTRecord `json:"aigt_record"` //aigt相关统计
}
type NodeRecord struct {
	NodeId     uint        `json:"node_id"`     //节点id
	NodePrice  string      `json:"node_price"`  //节点价格
	NodeRecord GoodsRecord `json:"node_record"` //节点记录
}
type MiningRecord struct {
	MiningId     uint        `json:"mining_id"`     //矿机id
	MiningPrice  string      `json:"mining_price"`  //矿机价格
	MiningRecord GoodsRecord `json:"mining_record"` //矿机记录
}
type GoodsRecord struct {
	Count  string `json:"count"`  //数量
	Amount string `json:"amount"` //金额
}
type AddressRecord struct {
	AvailableAddressCount string `json:"available_address_count"`
	NoConsumeAddressCount string `json:"no_consume_address_count"`
}
type USDTRecord struct {
	AutoWithdrawAmount             string `json:"auto_withdraw_amount"`              // 自动提现数额
	AutoWithdrawCount              string `json:"auto_withdraw_count"`               // 自动提现单数
	ApprovedWithdrawAmount         string `json:"approved_withdraw_amount"`          // 审核通过提现数额
	ApprovedWithdrawCount          string `json:"approved_withdraw_count"`           // 审核通过提现单数
	RefusedWithdrawAmount          string `json:"refused_withdraw_amount"`           // 审核拒绝提现数额
	RefusedWithdrawCount           string `json:"refused_withdraw_count"`            // 审核拒绝提现单数
	PurchaseMiningChargeConsume    string `json:"purchase_mining_charge_consume"`    // 购买矿机消耗充值U数额
	PurchaseMiningExchangeConsume  string `json:"purchase_mining_exchange_consume"`  // 购买矿机消耗兑换U数额
	SynthesisMiningChargeConsume   string `json:"synthesis_mining_charge_consume"`   // 合成矿机消耗充值U数额
	SynthesisMiningExchangeConsume string `json:"synthesis_mining_exchange_consume"` // 合成矿机消耗兑换U数额
	ExchangeAmount                 string `json:"exchange_amount"`                   // 闪兑数额
	ExchangeCount                  string `json:"exchange_count"`                    // 闪兑单数
	ChargeAmount                   string `json:"charge_amount"`                     // 充值数额
	ChargeCount                    string `json:"charge_count"`                      // 充值单数
	ExchangeBenefits               string `json:"exchange_benefits"`                 // usdt手续费分红
}
type AIGPRecord struct {
	AIGPSynthesisConsume     string        `json:"aigp_synthesis_consume"`       // 合成矿机消耗
	AIGPExchangeConsume      string        `json:"aigp_exchange_consume"`        // 闪兑消耗
	AIGPDynamicsBenefits     string        `json:"aigp_dynamics_benefits"`       // 动态奖励
	AIGPTeamBenefits         string        `json:"aigp_team_benefits"`           // 团队奖励
	AIGPParallelBenefits     string        `json:"aigp_parallel_benefits"`       // 平级奖励
	AIGPGlobalBenefits       string        `json:"aigp_global_benefits"`         // 全球分红
	AIGPOutcome              string        `json:"aigp_outcome"`                 // 挖矿产出
	AIGPOutcomeSortedByLevel []GoodsRecord `json:"aigp_outcome_sorted_by_level"` // 挖矿产出
	AIGPAllBenefits          string        `json:"aigp_all_benefits"`            // 总产出 （静态 + 动态）
}
type MiningOutcome struct {
	MiningId uint   `json:"mining_id"`
	Price    string `json:"price"`
	Outcome  string `json:"outcome"`
}
type AIGTRecord struct {
	//AIGTAirdropList []AirdropRecord `json:"aigt_airdrop_list"`
	AIGTAirdropAmount string `json:"aigt_airdrop_amount"` // 空投总量
	AIGTAirdropCount  string `json:"aigt_airdrop_count"`  // 空投单数
}
type AirdropRecord struct {
	WalletAddress string `json:"wallet_address"`
	Amount        string `json:"amount"`
	CreatedAt     string `json:"created_at"` // 创建时间
}

// GoodsReportResp 入金报表
type GoodsReportResp struct {
	MiningIncreaseToday []MiningRecord `json:"mining_increase_today"` //今日新增矿机
	MiningAll           []MiningRecord `json:"mining_all"`            //矿机总量
	NodeIncreaseToday   []NodeRecord   `json:"node_increase_today"`   //今日新增节点
	NodeAll             []NodeRecord   `json:"node_all"`              //节点总量
}

// LastReportResp 存量报表
type LastReportResp struct {
	ChargeUsdt   string `json:"charge_usdt"`   //充值usdt存量
	ExchangeUsdt string `json:"exchange_usdt"` //兑换usdt存量
	AIGP         string `json:"aigp"`          //aigp存量
	AIGT         string `json:"aigt"`          //aigt存量
}

// AirdropNodeReq 空投节点
type AirdropNodeReq struct {
	WalletAddress  string `json:"wallet_address"`  // 钱包地址
	NodeID         uint   `json:"node_id"`         // 节点id
	Hash           string `json:"hash"`            // 交易哈希或说明
	HasMiner       string `json:"has_miner"`       // 是否赠送矿机 0-否 1-是
	HasAIGT        string `json:"has_aigt"`        // 是否空投aigt 0-否 1-是
	HasPerformance string `json:"has_performance"` // 是否有业绩 0-否 1-是
}
type AirdropNodeResp struct {
}

// AirdropMinerReq 空投矿机
type AirdropMinerReq struct {
	WalletAddress  string `json:"wallet_address"`  // 钱包地址
	MinerID        uint   `json:"miner_id"`        // 矿机id
	Hash           string `json:"hash"`            // 交易哈希或说明
	HasAIGT        string `json:"has_aigt"`        // 是否空投aigt 0-否 1-是
	HasPerformance string `json:"has_performance"` // 是否有业绩 0-否 1-是
}
type AirdropMinerResp struct {
}

// BalanceReportReq 提现兑换排序
type BalanceReportReq struct {
	SortedBy string `json:"sorted_by"` // 0 - 不排序 1 - usdt升序 2 - usdt降序 3 - aigp升序 4 - aigp降序 5 - aigt升序 6 - aigt降序
}

// BalanceReportResp 提现兑换报表
type BalanceReportResp struct {
	USDTDailyWithdraw      string                   `json:"usdt_daily_withdraw"`       //usdt日提现
	AIGPDailyExchange      string                   `json:"aigp_daily_exchange"`       //aigp日兑换
	AIGTDailyWithdraw      string                   `json:"aigt_daily_withdraw"`       //aigt日提现
	BalanceSortedByAddress []BalanceSortedByAddress `json:"balance_sorted_by_address"` //地址提现兑换总额统计
}

type BalanceSortedByAddress struct {
	WalletAddress string `json:"wallet_address"`      // 钱包地址
	USDTWithdraw  string `json:"usdt_daily_withdraw"` //usdt总提现
	AIGPExchange  string `json:"aigp_daily_exchange"` //aigp总提现
	AIGTWithdraw  string `json:"aigt_daily_withdraw"` //aigt总提现
}

type SetYesterdayAllAIGPCostReq struct {
	Value string `json:"value"`
}

// NodeReportResp 节点报表
type NodeReportResp struct {
	NodeRecords []NodeRecord `json:"node_records"` //各节点记录
}

// MiningReportResp 矿机报表
type MiningReportResp struct {
	MiningRecords []MiningRecord `json:"mining_records"` //各矿机记录
}

// SelectManagerTradeOrderListReq 查询挂单
type SelectManagerTradeOrderListReq struct {
	WalletAddress  string   `json:"wallet_address"`  // 地址筛选
	TimeDuration   []int64  `json:"time_duration"`   // 时间筛选
	AmountDuration []string `json:"amount_duration"` // 剩余数量筛选
	PriceDuration  []string `json:"price_duration"`  // 单价筛选
	TradeType      string   `json:"trade_type"`      // 交易类型 0-卖单 1-买单 不传就所有
	Flag           string   `json:"flag"`            // 0-已取消 1-挂单中 2-已完成
	Top            string   `json:"top"`             // 置顶 0-未置顶 1-置顶 不传查全部
	Page           int      `json:"page"`            // 分页
}
type SelectManagerTradeOrderListResp struct {
	TradeOrderList []TradeOrderInfo `json:"trade_order_list"` // 单价
	PageSize       int              `json:"page_size"`
}

// TopTradeOrderReq 置顶挂单
type TopTradeOrderReq struct {
	TradeOrderId uint  `json:"trade_order_id"`
	Top          int64 `json:"top"` // 0-取消置顶 1-置顶
}
type TopTradeOrderResp struct {
}

// TradeOrderReportResp 挂单兑换报表
type TradeOrderReportResp struct {
	DailySoldU     string `json:"daily_sold_u"`     // 今日接受AIGP卖单支出(USDT)
	DailySoldA     string `json:"daily_sold_a"`     // 今日接受AIGP卖单收入(AIGP)
	DailyBuyU      string `json:"daily_buy_u"`      // 今日接受AIGP买单收入(USDT)
	DailyBuyA      string `json:"daily_buy_a"`      // 今日接受AIGP买单支出(AIGP)
	TradeOrderSold string `json:"trade_order_sold"` // 卖单总量(AIGP)
	TradeOrderBuy  string `json:"trade_order_buy"`  // 买单总量(USDT)
}
