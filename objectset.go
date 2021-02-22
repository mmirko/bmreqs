package bmreqs

type objectSet struct {
	name     string
	t        uint8
	elements map[string]*bmReq
}

type tbmReq interface {
	getReqName() string
	getReqType() uint8
	setReqName(string) error
	setReqType(uint8) error
	insertReq(string)
	removeReq(string)
	createSubReq(string, uint8) error
	removeSubReq(string) error
}
