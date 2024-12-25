package config

import "gopkg.in/go-playground/validator.v9"

const LOCAL_USERID_UINT = 0
const LOCAL_USERID_STRUCT = "LOCAL_USERID_STRUCT"
const LOCAL_MANAGERID_UINT = 0
const LOCAL_ADDRESS_STRING = "user_address_string"

// const LOCAL_USERID_INT64 = 0
const LOCAL_MANAGERID_INT64 = 0
const LOCAL_TOKEN = "token"

//const R_URL = "https://3620-1-162-143-119.ngrok-free.app/task"

const MESSAGE_FAIL = -1
const MESSAGE_SUCCESS = 0
const TOKEN_FAIL = -2
const STOP_SERVER = -3

const MESSAGE_PARSER_ERROR = "传入参数格式错误"
const MESSAGE_ADDRESS_ERROR = "钱包地址格式错误"
const MESSAGE_GET_USER_ERROR = "查询用户失败"
const MESSAGE_GET_KEY_VALUE_ERROR = "查询键值对错误."
const MESSAGE_TRANSACTION_ERROR = "事务失败"
const MESSAGE_GET_TRANSACTION_ERROR = "查询业务信息失败"
const MESSAGE_GET_PACKAGE_ERROR = "查询资源包失败"
const MESSAGE_GET_ORDER_ERROR = "查询订单失败"
const MESSAGE_GET_TEAM_ERROR = "查询团队失败"
const MESSAGE_GET_TRADE_ORDER_ERROR = "查询交易失败"
const MESSAGE_GET_REPORT_ERROR = "查询报表失败"
const MESSAGE_GET_PERFORMANCE_ERROR = "查询业绩失败"
const USER_FROZEN_STATUS = "查询报表失败"
const USER_NOT_ACTIVE_STATUS = "请先激活用户"
const TX_ERR = "交易失败"
const NOT_CURRENTLG_SUPPORTED = "暂时不支持"
const MESSAGE_ALREADY_HAS_NODE = "已有节点"

// 9758 test 环境
// 检测结构体
var Validate = validator.New()
