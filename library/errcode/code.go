package errcode

/*
错误码设计
*/

// 系统错误码
var (
	Success       = &Errno{Code: 0, Message: "Success"}
	ApiSuccess    = &Errno{Code: 0, Message: "Success"}
	SysSuccess    = &Errno{Code: 0, Message: "Success"}
	ApiFailure    = &Errno{Code: 1, Message: "FAILED"}
	SystemFailure = &Errno{Code: 1, Message: "FAILED"}
	Failed        = &Errno{Code: 1, Message: "Request failed"}
	SystemError   = &Errno{Code: 1, Message: "System internal error,please try again"}
	Unknown       = &Errno{Code: 1, Message: "unknown error"}
)

// HTTP客户端错误码
var (
	HttpLackUser        = &Errno{Code: 1000, Message: "Lack user"}                                  //缺少用户名
	HttpLackPassword    = &Errno{Code: 1001, Message: "Lack password"}                              //缺少密码
	HttpNoUser          = &Errno{Code: 1002, Message: "No such user"}                               //用户名不存在
	HttpInvalidUser     = &Errno{Code: 1003, Message: "User forbidden or deleted"}                  //该用户被禁用或者被删除
	HttpWrongPassword   = &Errno{Code: 1004, Message: "Wrong password"}                             //密码错误
	WrongUserOrPassword = &Errno{Code: 1004, Message: "Username is not exist or password is wrong"} //密码错误
	HttpInvalidEmail    = &Errno{Code: 1005, Message: "Invalid email"}                              //邮箱格式不正确
	HttpInvalidPassword = &Errno{Code: 1006, Message: "Invalid password"}                           //密码格式错误
	HttpUserExist       = &Errno{Code: 1007, Message: "User already exist"}                         //已经注册
	HttpNoEmail         = &Errno{Code: 1008, Message: "Email unregistered"}                         //当前邮箱未注册
	HttpSendEmailFail   = &Errno{Code: 1009, Message: "Send email failed"}                          //发送邮件失败
	HttpEmptyCode       = &Errno{Code: 1010, Message: "Empty code"}                                 //验证码不能为空
	HttpWrongTicket     = &Errno{Code: 1011, Message: "Wrong ticket"}                               //ticket错误
	HttpCodeOverdue     = &Errno{Code: 1012, Message: "Code overdue"}                               //验证码过期
	HttpWrongCode       = &Errno{Code: 1013, Message: "Wrong code"}                                 //验证码错误
	HttpSamePassword    = &Errno{Code: 1014, Message: "Password duplicate"}                         //新密码与原密码相同
	TokenExpire         = &Errno{Code: 1019, Message: "Token expire"}                               //token过期
	HttpWrongParameter  = &Errno{Code: 1021, Message: "Wrong parameter"}                            //接口参数有错
	InvalidParameter    = &Errno{Code: 1021, Message: "Invalid parameter"}                          //接口参数有错
	HttpNotLogin        = &Errno{Code: 1022, Message: "No login"}                                   //没有登录态
	HttpKeyCheckFail    = &Errno{Code: 1023, Message: "Sign check failed"}                          //加密校验错误
	TokenExist          = &Errno{Code: 1024, Message: "Token exist"}                                //加密校验错误
	HttpNoData          = &Errno{Code: 1025, Message: "There is no matched data"}                   //特定查询的时候出现数据并不存在的问题
	ValueOverFlow       = &Errno{Code: 1026, Message: "Reaching max value allowed, cannot proceed"}
	HttpInvalidSign     = &Errno{Code: 1027, Message: "invalid Http signiture using Joyent Scheme"} //Joyent规范的http签名验证不通过
	HttpTooFrequent     = &Errno{Code: 1028, Message: "Too frequent operation"}                     //操作过于频繁
)

// HTTP服务端错误码
var (
	HttpDbError               = &Errno{Code: 1200, Message: "Server database error"}         //服务器数据库错误
	HttpUnableGenToken        = &Errno{Code: 1201, Message: "Server unable gen token"}       //无法生成TOKEN
	HttpClientTimeout         = &Errno{Code: 1202, Message: "Http request timeout"}          //客户端超时
	HttpClientFailed          = &Errno{Code: 1203, Message: "Http request failed"}           //客户端错误
	HttpClientConnectionError = &Errno{Code: 1204, Message: "Http request is not reachable"} //客户端错误
	HttpRequestTimeout        = &Errno{Code: 1205, Message: "Http request is timeout"}       //客户端请求超时
	HttpServerTimeout         = &Errno{Code: 1206, Message: "Http server is timeout"}        //服务端超时
	InvalidAuth               = &Errno{Code: 1210, Message: "Operation is not permitted"}
	Timeout                   = &Errno{Code: 1211, Message: "Operation is not permitted"}
	InvalidMethod             = &Errno{Code: 1212, Message: "Invalid method called"}
	InvalidCode               = &Errno{Code: 1213, Message: "Invalid code"}
	SessionStart              = &Errno{Code: 1214, Message: "Session create failed"}
	DeviceExist               = &Errno{Code: 1215, Message: "Device is not exist or not own it"}
	CaptchaMatch              = &Errno{Code: 1226, Message: "Captcha is not correct"}
	RefreshTokenError         = &Errno{Code: 1227, Message: "Refresh token is not exist"}
	AccessTokenError          = &Errno{Code: 1228, Message: "Access token is not exist"}
	DeviceOffline             = &Errno{Code: 1229, Message: "Device is offline"}
	MalformedData             = &Errno{Code: 1230, Message: "Server data is malformed, task cannot be done"}
	NotSupported              = &Errno{Code: 1231, Message: "Not support yet"}
	MaxRetry                  = &Errno{Code: 1232, Message: ""}
	NotExpect                 = &Errno{Code: 1233, Message: ""}
	HttpGoogleSyntaxFail      = &Errno{Code: 1240, Message: "Bad request sent by the client due to invalid syntax"}
	HttpGoogleNotFound        = &Errno{Code: 1241, Message: "The requested resource could not be found"}
	HttpGoogleThreshold       = &Errno{Code: 1242, Message: "too many sync requests in a given amount of time"}
	HttpAlexaNoEntity         = &Errno{Code: 1243, Message: "alexa does not contain such entity"}
	HttpAlexaTokenFail        = &Errno{Code: 1244, Message: "alexa update token failed"}
)

// hub
var HubDeviceNotSupport = &Errno{Code: 1250, Message: "sub device type is not supported"}

// mrc
var OverNumberLimit = &Errno{Code: 1255, Message: "The number of remote control boards exceeded the limit"}

// rate limit
var (
	HttpBeyondRateLimit   = &Errno{Code: 1300, Message: "Your request rate is beyond limit, access is forbidden for a while"}
	HttpTooManyLoginToken = &Errno{Code: 1301, Message: ""}
	ServiceMainTain       = &Errno{Code: 1310, Message: ""}
)

// garage
var (
	HavingCompatibleModel    = &Errno{Code: 1256, Message: "compatible mode is having"}
	NotHavingCompatibleModel = &Errno{Code: 1257, Message: "compatible mode is not having"}
)

// mysql错误码
var (
	MysqlDupEntry        = &Errno{Code: 1400, Message: "Request failed by the duplicate major key"}
	MysqlRowIsReferenced = &Errno{Code: 1401, Message: "Request failed by the referred row"}
	MysqlTimeout         = &Errno{Code: 1402, Message: "Mysql execution timeout"}
	MysqlUpdateNoAffect  = &Errno{Code: 1403, Message: "Data is already the newest in db"}
	MysqlTransDeadLock   = &Errno{Code: 1404, Message: "Mysql deadlock detected"}
	MysqlHasGoneAway     = &Errno{Code: 1405, Message: "Mysql has gone away, try it again"}
	MysqlInProgress      = &Errno{Code: 1406, Message: "Operation now in progress, try it again"}
)

// EMQX错误码
var (
	EMQXHttpMNGMTFail      = &Errno{Code: 1550, Message: ""}
	EMQXCriticalAlarm      = &Errno{Code: 1551, Message: ""}
	EMQXMonitorRestarting  = &Errno{Code: 1800, Message: ""}
	EMQXMonitorFailListen  = &Errno{Code: 1801, Message: ""}
	EMQXMonitorFailRestart = &Errno{Code: 1802, Message: ""}
)

// redis错误码
var (
	RedisTimeout     = &Errno{Code: 1601, Message: "Redis excution timeout,and close connection"}
	RedisWaitTimeout = &Errno{Code: 1602, Message: "Redis queue is too long,and wait timeout"}
	RedisClose       = &Errno{Code: 1603, Message: "Redis disconnect"}
	RedisScriptLoad  = &Errno{Code: 1604, Message: "Redis lua script load failed, check the syntax"}
	RedisScriptExec  = &Errno{Code: 1605, Message: "Redis lua script exec failed, check the syntax or parameters"}
)

var MailTimeout = &Errno{Code: 1700, Message: "Send mail timeout"}

// DynamoDB错误码
var (
	DynamoDBInternal                 = &Errno{Code: 2000, Message: ""}
	DynamoDBNoAccess                 = &Errno{Code: 2001, Message: ""}
	DynamoDBConditionFail            = &Errno{Code: 2002, Message: ""}
	DynamoDBSignFail                 = &Errno{Code: 2003, Message: ""}
	DynamoDBCollectionSizeLimit      = &Errno{Code: 2004, Message: ""}
	DynamoDBActionLimit              = &Errno{Code: 2005, Message: ""}
	DynamoDBNoAuthToken              = &Errno{Code: 2006, Message: ""}
	DynamoDBOutProvisionedThroughput = &Errno{Code: 2007, Message: ""}
	DynamoDBRequestLimit             = &Errno{Code: 2008, Message: ""}
	DynamoDBResourceInUse            = &Errno{Code: 2009, Message: ""}
	DynamoDBResourceNotFound         = &Errno{Code: 2010, Message: ""}
	DynamoDBThrottling               = &Errno{Code: 2011, Message: ""}
	DynamoDBClientUnrecognized       = &Errno{Code: 2012, Message: ""}
	DynamoDBUnavailable              = &Errno{Code: 2013, Message: ""}
	DynamoDBServerError              = &Errno{Code: 2014, Message: ""}
)

// MQTT/固件返回错误码
var (
	MQTTProtocolError        = &Errno{Code: 1500, Message: "Packet is not specified by mqtt protocol"}
	MQTTError                = &Errno{Code: 1501, Message: "Mqtt connection error"}
	MQTTClose                = &Errno{Code: 1502, Message: "Mqtt close event"}
	MQTTTimeout              = &Errno{Code: 1504, Message: "Mqtt execution timeout"}
	MQTTBadClient            = &Errno{Code: 1505, Message: ""}
	MQTTBadMessage           = &Errno{Code: 1506, Message: ""}
	DeviceValueError         = &Errno{Code: 5000, Message: ""}
	DeviceSignError          = &Errno{Code: 5000, Message: ""}
	DeviceCommandUnsupport   = &Errno{Code: 5000, Message: ""}
	DeviceCommandInProgress  = &Errno{Code: 5000, Message: ""}
	MQTTHubRangeOut          = &Errno{Code: 5060, Message: ""}
	MQTTSubDeviceUnreachable = &Errno{Code: 5061, Message: ""}
	MQTTSubDeviceNoExist     = &Errno{Code: 5062, Message: ""}
)
