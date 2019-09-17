package msg

import "hnzzmj-server/game/mahjong"

type C2S_CreateHNZZRoom struct {
	mahjong.HNZZRule
}

const (
	S2C_CreateRoom_OK              = 0
	S2C_CreateRoom_InnerError      = 1 // 创建房间出错，请稍后重试
	S2C_CreateRoom_CreateRepeated  = 2 // "房间: " + S2C_CreateRoom.RoomNumber + " 已存在"
	S2C_CreateRoom_InOtherRoom     = 3 // 正在其他房间对局，是否回去？
	S2C_CreateRoom_LackOfRoomCards = 4 // 房卡不足，需要 + S2C_S2C_CreateRoom.RoomCards + 张房卡才能游戏
	S2C_CreateRoom_RuleError       = 5 // 规则错误，请稍后重试
	S2C_CreateRoom_LocationError   = 6 // 定位参数错误，请检查GPS
)

type S2C_CreateRoom struct {
	Error     int
	RoomCards int // 需要的房卡数
}

type C2S_EnterRoom struct {
	RoomNumber string
	GPS        bool // 是否开启GPS
	Location   []float64
}

const (
	S2C_EnterRoom_OK              = 0
	S2C_EnterRoom_NotCreated      = 1 // "房间: " + S2C_EnterRoom.RoomNumber + " 未创建"
	S2C_EnterRoom_Full            = 2 // "房间: " + S2C_EnterRoom.RoomNumber + " 玩家人数已满"
	S2C_EnterRoom_Unknown         = 4 // 进入房间出错，请稍后重试
	S2C_EnterRoom_LackOfRoomCards = 5 // 房卡不足，需要 + S2C_EnterRoom.RoomCards + 张房卡才能进入
	S2C_EnterRoom_IPConflict      = 6 // IP重复，无法进入
	S2C_EnterRoom_GPSNotOpen      = 7 // 定位失败，请检查GPS是否开启
	S2C_EnterRoom_LocationError   = 8 // 定位参数错误，请检查GPS
	S2C_EnterRoom_NotRightNow     = 9 // 比赛暂未开始，请到时再来
)

type S2C_EnterRoom struct {
	Error         int
	RoomType      int // 房间类型， 0 练习，1 房卡匹配，2 私人
	RedPacketType int // 红包种类(元): 1、10、100、999
	RoomNumber    string
	Position      int
	RoomDesc      string
	MaxPlayers    int // 最大玩家数
	MaxRounds     int // 总局数
	RoomCards     int // 进入房间需要的房卡数
	GamePlaying   bool
}

type C2S_GetAllPlayers struct{}

type S2C_SitDown struct {
	Position   int
	AccountID  int
	LoginIP    string
	Nickname   string
	Headimgurl string
	Sex        int
	Owner      bool
	Ready      bool
	Location   []float64
}

type S2C_StandUp struct {
	Position int
}

type C2S_ExitOrDisbandRoom struct{}

const (
	S2C_ExitRoom_OK          = 0
	S2C_ExitRoom_GamePlaying = 1 // 游戏进行中，不能退出房间
)

type S2C_ExitRoom struct {
	Error    int
	Position int
}

const (
	S2C_DisbandRoom_OK           = 0
	S2C_DisbandRoom_PlayerRefuse = 1 // 玩家拒绝
)

type S2C_DisbandRoom struct {
	Error            int
	RoomNumber       string
	OwnerNickName    string // 房主
	RejecterNickName string // 拒绝者
}
