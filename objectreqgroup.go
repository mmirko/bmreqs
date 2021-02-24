package bmreqs

import (
	"fmt"
)

func (rg *ReqGroup) init() {
}

//

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

//

func (rg *ReqGroup) insertReq(req string) error {
	return nil
}

func (rg *ReqGroup) removeReq(req string) error {
	return nil
}

func (rg *ReqGroup) matchReq(req string) bool {
	if rg.bmReqMap == nil {
		rg.bmReqMap = make(map[string]bmReqSet)
	}
	if _, ok := rg.bmReqMap[req]; ok {
		return true
	}
	return false
}

//

func (rg *ReqGroup) getReqs() string {
	if rg.bmReqMap == nil {
		return ""
	}
	return fmt.Sprint(rg.bmReqMap)
}

//

func (rg *ReqGroup) supportSubReq() bool {
	return true
}

func (rg *ReqGroup) getSubReqMap() bmReqMap {
	return rg.bmReqMap
}

func (rg *ReqGroup) getSubReq(req string) bmReqSet {
	if rg.bmReqMap == nil {
		return nil
	}
	if node, ok := rg.bmReqMap[req]; ok {
		return node
	}
	return nil
}

func (rg *ReqGroup) setSubReq(req string, node string, ob bmReqSet) {
	if rg.bmReqMap != nil {
		if reqset, ok := rg.bmReqMap[req]; ok {
			reqset.setSubReq(node, "", ob)
		}
	}
}
