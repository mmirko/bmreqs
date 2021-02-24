package bmreqs

import (
	"fmt"
)

type objectSet struct {
	name string
	t    uint8
	bmReqMap
}

func (o *objectSet) init() {
	o.bmReqMap = make(map[string]bmReqSet)
}

//

func (o *objectSet) getName() string {
	return o.name
}

func (o *objectSet) getType() uint8 {
	return o.t
}

func (o *objectSet) setName(name string) error {
	o.name = name
	return nil
}

func (o *objectSet) setType(t uint8) error {
	o.t = t
	return nil
}

//

func (o *objectSet) insertReq(req string) error {
	if o.bmReqMap == nil {
		o.bmReqMap = make(bmReqMap)
	}
	//	if req != "" {
	o.bmReqMap[req] = nil
	//	}
	return nil
}

func (o *objectSet) removeReq(req string) error {
	if o.bmReqMap != nil {
		if _, ok := o.bmReqMap[req]; ok {
			delete(o.bmReqMap, req)
		}
	}
	return nil
}

func (o *objectSet) matchReq(req string) bool {
	if o.bmReqMap == nil {
		return false
	}
	if _, ok := o.bmReqMap[req]; ok {
		return true
	}
	return false
}

//

func (o *objectSet) getReqs() string {
	if o.bmReqMap == nil {
		return ""
	}
	return fmt.Sprint(o.bmReqMap)
}

//

func (o *objectSet) supportSubReq() bool {
	return true
}

func (o *objectSet) getSubReqMap() bmReqMap {
	return o.bmReqMap
}

func (o *objectSet) getSubReq(req string) bmReqSet {
	if o.bmReqMap == nil {
		return nil
	}
	if node, ok := o.bmReqMap[req]; ok {
		return node
	}
	return nil
}

func (o *objectSet) setSubReq(req string, node string, ob bmReqSet) {
	fmt.Println(req)
	//	o.bmReqMap[req] = ob
}
