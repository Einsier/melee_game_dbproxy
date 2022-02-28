## 表结构

### account（账户表）

字段名 | 类型 | 含义
---------|----------|---------
 _id | ObjectID | 
 player_id | int32 | 玩家id
 account_name | string | 账户名称
 phone | string | 手机号
 password | string | 密码
 recent_login | int64 | 最近登录时间
 create_at | int64 | 创建时间
 update_at | int64 | 更新时间
 delete | bool | 是否注销

### player（玩家表）

字段名 | 类型 | 含义
---------|----------|---------
 player_id | int32 | 玩家id
 account_id | string | 账户id
 nickname | string | 玩家昵称
 game_count | int32 | 参与游戏总局数
 kill_num | int32 | 总击杀数
 max_kill | int32 | 最高单局击杀数
 create_at | int64 | 创建时间
 update_at | int64 | 更新时间

 ### single_game（单局表）

 字段名 | 类型 | 含义
---------|----------|---------
_id | ObjectID | 
 players | []int32 | 参与玩家id
 start_time | int64 | 开始时间
 end_time | int64 | 结束时间
 create_at | int64 | 创建时间

## 项目目录
- api：与外模块进行rpc通信
- configs：一些配置信息
- db：与数据库进行交互
- model：对应的一些实体类

## 接口
- IsAccountLegal：根据手机号和密码判断当前账户是否合法，若合法返回对应的账户id，否则返回相应错误
```go
type IsAccountLegalRequest struct {
	Phone    string
	Password string
}

type IsAccountLegalResponse struct {
	Ok       bool
	Error    ErrorType //失败
	PlayerId int32
}
```
- SearchPlayerInfo：根据玩家id查找玩家信息
```go
type PlayerInfo struct {
	PlayerId  int32
	NickName  string
	GameCount int32 //参与游戏数
	KillNum   int32 //总击杀数
	MaxKill   int32 //最高单局击杀数
}
type SearchPlayerInfoRequest struct {
	PlayerId int32
}

type SearchPlayerInfoResponse struct {
	Ok    bool
	Error ErrorType   //失败
	Info  *PlayerInfo //成功
}
```
- UpdatePlayerInfo：根据玩家id和更改信息修改玩家信息
```go
type UpdatePlayerInfoRequest struct {
	Info *PlayerInfo
}

type UpdatePlayerInfoResponse struct {
	Ok    bool
	Error ErrorType
	Info  *PlayerInfo //如果失败,返回原来的info
}
```
- AddSingleGameInfo：添加单局游戏信息
```go
type SingleGameInfo struct {
	Players   []int32 //参加游戏的玩家id
	StartTime int64   //游戏开始时间
	EndTime   int64   //游戏结束时间
}

type AddSingleGameInfoRequest struct {
	Info *SingleGameInfo
}

type AddSingleGameInfoResponse struct {
	Ok    bool
	Error ErrorType
}
```