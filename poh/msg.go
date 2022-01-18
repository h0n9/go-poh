package poh

import (
	"encoding/json"

	"github.com/h0n9/go-poh/types"
	"github.com/h0n9/go-poh/util"
)

type Msg struct {
	Hash    types.Hash `json:"hash"`
	History types.Hash `json:"history"`
	Body    string     `json:"body"`
}

type MsgToHash struct {
	Hash    types.Hash `json:"-"`
	History types.Hash `json:"history"`
	Body    string     `json:"body"`
}

func NewMsg(history types.Hash, body string) (*Msg, error) {
	msg := Msg{
		History: history,
		Body:    body,
	}
	err := msg.hash()
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (msg *Msg) hash() error {
	msgToHash := MsgToHash(*msg)
	data, err := json.Marshal(&msgToHash)
	if err != nil {
		return err
	}
	msg.Hash = util.Sha256(data)
	return nil
}

func (msg *Msg) String() string {
	data, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}
