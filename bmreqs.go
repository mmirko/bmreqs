package bmreqs

const (
	// ObjectSet is the set-alike requirements
	ObjectSet = uint8(0) + iota
)

type bmReq interface {
	getReqName() string
	getReqType() uint8
	setReqName(string) error
	setReqType(uint8) error
	insertReq(string)
	removeReq(string)
	createSubReq(string, uint8) error
	removeSubReq(string) error
}

// ReqGroup is the entry point data structure for all the bmreqs operations
type ReqGroup struct {
	reqs map[string]*bmReq
}

// NewReqGroup creates a new *ReqGroup object
func NewReqGroup() *ReqGroup {
	return nil
}
