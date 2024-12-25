package types

type TeamResp struct {
	List []TeamOne `json:"list"`
}

type TeamOne struct {
	Address              string `json:"address"`                // 地址
	TeamNum              string `json:"team_num"`               // 团队人数
	Performance          string `json:"performance"`            // 总业绩
	SmallZonePerformance string `json:"small_zone_performance"` // 小区业绩
}

type TeamReq struct {
	WalletAddress string `json:"wallet_address"` // 地址
}
type GetTeamListReq struct {
}

type GetTeamListResp struct {
	TeamList []TeamInfo `json:"team_list"`
}
type TeamInfo struct {
	WalletAddress string `json:"wallet_address"` // 团队节点地址
	TeamName      string `json:"team_name"`      // 团队名
	Flag          string `json:"flag"`           // 启用标志(0-不限制 1-限制转账 2-限制提现 3-全限制)
}

type CreateNewTeamReq struct {
	WalletAddress string `json:"wallet_address"` // 团队节点地址
	TeamName      string `json:"team_name"`      // 团队名
	Flag          string `json:"flag"`           // 启用标志(0-不限制 1-限制转账 2-限制提现 3-全限制)
}

type CreateNewTeamResp struct {
}
type TeamInfoReq struct {
	WalletAddress string `json:"wallet_address"` // 团队节点地址
}

type TeamInfoResp struct {
	UserTeamInfoList             []UserTeamInfo `json:"team_list"`
	PerformanceTotal             string         `json:"performance_total"`               // 总业绩
	SmallZonePerformance         string         `json:"small_zone_performance"`          // 小区业绩
	PerformanceIncrease          string         `json:"performance_increase"`            // 日新增总业绩
	SmallZonePerformanceIncrease string         `json:"small_zone_performance_increase"` // 日新增小区业绩
	WithdrawTotal                string         `json:"withdraw_total"`                  // 累计提现
	WithdrawDaily                string         `json:"withdraw_daily"`                  // 日提现
	NewUserPerformance           string         `json:"new_user_performance"`            // 新用户业绩
	RePurchasePerformance        string         `json:"re_purchase_performance"`         // 复投新增业绩
	AIGPTotal                    string         `json:"aigp_total"`                      // 总aigp存量
	USDTTotal                    string         `json:"usdt_total"`                      // 总usdt存量
	DailySoldU                   string         `json:"daily_sold_u"`                    // 今日接受AIGP卖单支出(USDT)
	DailySoldA                   string         `json:"daily_sold_a"`                    // 今日接受AIGP卖单收入(AIGP)
	DailyBuyU                    string         `json:"daily_buy_u"`                     // 今日接受AIGP买单收入(USDT)
	DailyBuyA                    string         `json:"daily_buy_a"`                     // 今日接受AIGP买单支出(AIGP)
}
type UserTeamInfo struct {
	WalletAddress        string `json:"wallet_address"`         // 地址
	RecommendAddress     string `json:"recommend_address"`      // 推荐人地址
	ABalance             string `json:"a_balance"`              // A币余额
	BBalance             string `json:"b_balance"`              // B币余额
	USDTBalance          string `json:"usdt_balance"`           // 充值USDT余额
	ExUSDTBalance        string `json:"ex_usdt_balance"`        // 兑换USDT余额
	Level                int64  `json:"level"`                  // 等级
	Performance          string `json:"performance"`            // 除自己外总业绩
	Consume              string `json:"consume"`                // usdt消费
	PerformanceTotal     string `json:"performance_total"`      // 总业绩
	SmallZonePerformance string `json:"small_zone_performance"` // 小区业绩
	Flag                 string `json:"flag"`                   // 启用标志(0-正常 -1-冻结)
	CreatedAt            string `json:"created_at"`             // 创建日期
}
type DailyTeamDataReq struct {
	WalletAddress string `json:"wallet_address"` // 团队节点地址
	Date          int64  `json:"date"`           // 团队节点地址
}

type DailyTeamDataResp struct {
	PerformanceTotal             string `json:"performance_total"`               // 总业绩
	SmallZonePerformance         string `json:"small_zone_performance"`          // 小区业绩
	PerformanceIncrease          string `json:"performance_increase"`            // 日新增总业绩
	SmallZonePerformanceIncrease string `json:"small_zone_performance_increase"` // 日新增小区业绩
	WithdrawTotal                string `json:"withdraw_total"`                  // 累计提现
	WithdrawDaily                string `json:"withdraw_daily"`                  // 日提现
	NewUserPerformance           string `json:"new_user_performance"`            // 新用户业绩
	RePurchasePerformance        string `json:"re_purchase_performance"`         // 复投新增业绩
	AIGPTotal                    string `json:"aigp_total"`                      // 总aigp存量
	USDTTotal                    string `json:"usdt_total"`                      // 总usdt存量
}
type UpdateTeamLimitFlagReq struct {
	WalletAddress string `json:"wallet_address"` // 团队节点地址
	TeamName      string `json:"team_name"`      // 团队名
	Flag          string `json:"flag"`           // 启用标志(0-不限制 1-限制转账 2-限制提现 3-全限制)
}

type UpdateTeamLimitFlagResp struct {
}
