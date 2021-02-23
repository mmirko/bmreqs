package bmreqs

import (
	"context"
	"fmt"
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
	rg.bmReqMap = make(map[string]bmReq)
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

func (rg *ReqGroup) getName() string {
	return "/"
}

func (rg *ReqGroup) getType() uint8 {
	return ObjectRoot
}

func (rg *ReqGroup) setName(name string) error {
	return nil
}

func (rg *ReqGroup) setType(t uint8) error {
	return nil
}

func (rg *ReqGroup) getReqMap() bmReqMap {
	return rg.bmReqMap
}

func (rg *ReqGroup) insertReq(req string) error {
	if rg.bmReqMap == nil {
		rg.bmReqMap = make(map[string]bmReq)
	}
	rg.bmReqMap[req] = nil
	return nil
}

func (rg *ReqGroup) removeReq(req string) error {
	if rg.bmReqMap != nil {
		if _, ok := rg.bmReqMap[req]; ok {
			delete(rg.bmReqMap, req)
		}
	}
	return nil
}

func (rg *ReqGroup) getReqs() string {
	if rg.bmReqMap == nil {
		return ""
	}
	return fmt.Sprint(rg.bmReqMap)
}

func (rg *ReqGroup) supportSubReq() bool {
	return true
}

func (rg *ReqGroup) matchReq(req string) bool {
	if rg.bmReqMap == nil {
		rg.bmReqMap = make(map[string]bmReq)
	}
	if _, ok := rg.bmReqMap[req]; ok {
		return true
	}
	return false
}

func (rg *ReqGroup) getSubReq(req string) bmReq {
	if rg.bmReqMap == nil {
		return nil
	}
	if node, ok := rg.bmReqMap[req]; ok {
		return node
	}
	return nil
}
