package StructTools

/*
	定义一个共享的结构体
*/
type MessageInfo struct {
	Status      bool   `json:"status"`       // 群聊属性
	AtMe        bool   `json:"atme"`         // 是否@我
	ReplyResult string `json:"reply_result"` // 自动回复
	Reply       bool   `json:"reply"`        // 自动回复状态
	PassResult  string `json:"pass_result"`  // pass 原因
	Pass        bool   `json:"pass"`         // pass 状态
}
