package bmreqs

import (
	"errors"
	"fmt"
)

type objectSet struct {
	name string
	t    uint8
	set  map[string]*bmReqObj
}

func (o *objectSet) init() {
	o.set = make(map[string]*bmReqObj)
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
	if o.set == nil {
		return fmt.Errorf("Uninitialized Set")
	}
	newObj := new(bmReqObj)
	newObj.init()
	o.set[req] = newObj
	return nil
}

func (o *objectSet) removeReq(req string) error {
	if o.set != nil {
		if _, ok := o.set[req]; ok {
			delete(o.set, req)
		}
	} else {
		return errors.New("Uninitialized Set")
	}
	return nil
}

//

func (o *objectSet) getReqs() string {
	if o.set == nil {
		return ""
	}
	return fmt.Sprint(o.set)
}

//

func (o *objectSet) supportSub() bool {
	return true
}

func (o *objectSet) getSub(req string) (*bmReqObj, error) {
	if o.set == nil {
		return nil, errors.New("Uninitialized Set")
	}
	if node, ok := o.set[req]; ok {
		return node, nil
	}
	return nil, nil
}
