package bmreqs

import (
	"context"
)

// Requirements types
const (
	// ObjectSet is the set-alike requirements
	ObjectRoot = uint8(0) + iota
	ObjectSet
)

// Requirements operations
const (
	OpAdd = uint8(0) + iota
	OpGet
)

type ReqRequest struct {
	Node  string
	Name  string
	T     uint8
	Op    uint8
	Value string
}

type ReqResponse struct {
	Value string
	Error error
}

// ReqGroup is the entry point data structure for all the bmreqs operations
type ReqGroup struct {
	ask    chan ReqRequest
	answer chan ReqResponse
	ctx    context.Context
	cancel context.CancelFunc
	bmReqMap
}

// NewReqGroup creates a new *ReqGroup object
func NewReqGroup() *ReqGroup {
	rg := new(ReqGroup)
	rg.ask = make(chan ReqRequest)
	rg.answer = make(chan ReqResponse)
	rg.bmReqMap = make(map[string]bmReqSet)
	ctx, cancel := context.WithCancel(context.Background())
	rg.ctx = ctx
	rg.cancel = cancel
	go rg.run()
	return rg
}

func (rg *ReqGroup) Requirement(req ReqRequest) ReqResponse {
	rg.ask <- req
	return <-rg.answer

}

func (rg *ReqGroup) Close() {
	rg.cancel()
}
