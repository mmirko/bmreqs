package bmreqs

import "fmt"

type objectSet struct {
	name string
	t    uint8
	bmReqMap
}

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

func (o *objectSet) getReqMap() bmReqMap {
	return o.bmReqMap
}

func (o *objectSet) insertReq(req string) error {
	if o.bmReqMap == nil {
		o.bmReqMap = make(map[string]bmReq)
	}
	if req != "" {
		o.bmReqMap[req] = nil
	}
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

func (o *objectSet) getReqs() string {
	if o.bmReqMap == nil {
		return ""
	}
	return fmt.Sprint(o.bmReqMap)
}

func (o *objectSet) supportSubReq() bool {
	return true
}

func (o *objectSet) matchReq(req string) bool {
	if o.bmReqMap == nil {
		o.bmReqMap = make(map[string]bmReq)
	}
	if _, ok := o.bmReqMap[req]; ok {
		return true
	}
	return false
}

func (o *objectSet) getSubReq(req string) bmReq {
	if o.bmReqMap == nil {
		return nil
	}
	if node, ok := o.bmReqMap[req]; ok {
		return node
	}
	return nil
}
