package bmreqs

type bmReqObj struct {
	bmReqMap
}

type bmReqSet interface {

	// Initialization
	init()

	// Handling of the single requirement node
	getName() string
	getType() uint8
	setName(string) error
	setType(uint8) error

	// Insert and remove requirements from/to the current node
	insertReq(string) error
	removeReq(string) error
	matchReq(string) bool

	// Exporting requirements for the current node
	getReqs() string

	// SubRequirements
	supportSubReq() bool
	getSubReqMap() bmReqMap
	getSubReq(string) bmReqSet
	setSubReq(string, string, bmReqSet)
}

type bmReqMap map[string]bmReqSet
