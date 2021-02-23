package bmreqs

type bmReq interface {
	getName() string
	getType() uint8
	setName(string) error
	setType(uint8) error

	getReqMap() bmReqMap

	insertReq(string) error
	removeReq(string) error
	matchReq(string) bool

	getReqs() string

	supportSubReq() bool
	getSubReq(string) bmReq
}

type bmReqMap map[string]bmReq
