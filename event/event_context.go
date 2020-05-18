package event

import (
	"go-ratel/common"
	"log"
)

// 应用上下文（应用层）
type Context struct {
	clientChan *chan common.ClientTransferDataProtoc
	serverChan *chan common.ServerTransferDataProtoc
	UserId     int
}

func NewEventContext(clientChan *chan common.ClientTransferDataProtoc, serverChan *chan common.ServerTransferDataProtoc) *Context {
	return &Context{
		clientChan: clientChan,
		serverChan: serverChan,
		UserId:     0,
	}
}

func (ctx *Context) DoListen() {
	go func() {
		for {
			transferData := <-*ctx.clientChan
			ctx.call(transferData.Code, transferData.Data)
		}
	}()
}

func (ctx *Context) call(code string, data string) {
	switch code {
	case CODE_CLIENT_CONNECT:
		ListenerClientConnect(ctx, data)
	case CODE_CLIENT_EXIT:
		ListenerClientExit(ctx, data)
	case CODE_CLIENT_KICK:
		ListenerClientKick(ctx, data)
	case CODE_CLIENT_NICKNAME_SET:
		ListenerClientNicknameSet(ctx, data)
	case CODE_GAME_LANDLORD_CONFIRM:
		ListenerGameLandlordConfirm(ctx, data)
	case CODE_GAME_LANDLORD_CYCLE:
		ListenerGameLandlordCycle(ctx, data)
	case CODE_GAME_LANDLORD_ELECT:
		ListenerGameLandlordElect(ctx, data)
	case CODE_GAME_OVER:
		ListenerGameOver(ctx, data)
	case CODE_GAME_POKER_PLAY:
		ListenerGamePokerPlay(ctx, data)
	case CODE_GAME_POKER_PLAY_CANT_PASS:
		ListenerGamePokerPlayCantPass(ctx, data)
	case CODE_GAME_POKER_PLAY_INVALID:
		ListenerGamePokerPlayInvalid(ctx, data)
	case CODE_GAME_POKER_PLAY_LESS:
		ListenerGamePokerPlayLess(ctx, data)
	case CODE_GAME_POKER_PLAY_MISMATCH:
		ListenerGamePokerPlayMismatch(ctx, data)
	case CODE_GAME_POKER_PLAY_ORDER_ERROR:
		ListenerGamePokerPlayOrderError(ctx, data)
	case CODE_GAME_POKER_PLAY_PASS:
		ListenerGamePokerPlayPass(ctx, data)
	case CODE_GAME_POKER_PLAY_REDIRECT:
		ListenerGamePokerPlayRedirect(ctx, data)
	case CODE_GAME_STARTING:
		ListenerGameStarting(ctx, data)
	case CODE_PVE_DIFFICULTY_NOT_SUPPORT:
		ListenerPVEDifficultyNotSupport(ctx, data)
	case CODE_ROOM_CREATE_SUCCESS:
		ListenerRoomCreateSuccess(ctx, data)
	case CODE_ROOM_JOIN_FAIL_BY_FULL:
		ListenerRoomJoinFailByFull(ctx, data)
	case CODE_ROOM_JOIN_FAIL_BY_INEXIST:
		ListenerRoomJoinFailByInExist(ctx, data)
	case CODE_ROOM_JOIN_SUCCESS:
		ListenerRoomJoinSuccess(ctx, data)
	case CODE_SHOW_OPTIONS:
		ListenerShowOptions(ctx, data)
	case CODE_SHOW_OPTIONS_PVE:
		ListenerShowOptionsPVE(ctx, data)
	case CODE_SHOW_OPTIONS_PVP:
		ListenerShowOptionsPVP(ctx, data)
	case CODE_SHOW_OPTIONS_SETTING:
		ListenerShowOptionsSettings(ctx, data)
	case CODE_SHOW_POKERS:
		ListenerShowPokers(ctx, data)
	case CODE_SHOW_ROOMS:
		ListenerShowRooms(ctx, data)
	default:
		log.Println("Event code invalid")
	}
}

func (ctx *Context) pushToServer(serverCode string, data string) {
	transferData := common.ServerTransferDataProtoc{
		Code: serverCode,
		Data: data,
	}
	*ctx.serverChan <- transferData
}
