package bmreqs

import (
	"errors"
	"fmt"
	"strings"
)

func (rg *ReqRoot) decodeNode(node string) (*bmReqObj, error) {
	if node == "/" {
		return rg.obj, nil
	}

	seqlist := strings.Split(node, "/")
	var currNode *bmReqObj = rg.obj
	var lName string
	var lReq string
	for _, leaf := range seqlist {
		currMap := currNode.getMap()
		if leaf != "" {
			decLeaf := strings.Split(leaf, ":")
			if len(decLeaf) != 2 {
				return nil, errors.New("Malformed node")
			}
			lReq = decLeaf[0]
			lName = decLeaf[1]

			if reqset, ok := currMap[lReq]; ok {
				if reqset.supportSub() {
					if newNode, err := reqset.getSub(lName); err == nil {
						if newNode != nil {
							currNode = newNode
						} else {
							return nil, errors.New("Node empty")
						}
					} else {
						return nil, errors.New("Node not exists")
					}
				} else {
					return nil, errors.New("Requirement Set does not support sub requirements")
				}
			} else {
				return nil, errors.New("Unknown Requirement Set")
			}
		}
	}

	return currNode, nil
}

func (rg *ReqRoot) run() {
	ctx := rg.ctx
	for {
		select {
		case <-ctx.Done():
			return
		case req := <-rg.ask:
			resp := ReqResponse{Value: "", Error: nil}
			switch req.Op {
			case OpAdd:
				if node, err := rg.decodeNode(req.Node); err == nil {
					currMap := node.getMap()
					if reqSet, ok := currMap[req.Name]; ok {
						if reqSet.getType() != req.T {
							resp.Error = errors.New("Insert failed: Mismatch types")
						} else {
							if err := reqSet.insertReq(req.Value); err != nil {
								resp.Error = errors.New("Insert failed: " + fmt.Sprint(err))
							}
						}
					} else {
						switch req.T {
						case ObjectSet:
							newReq := new(objectSet)
							newReq.init()
							newReq.setName(req.Name)
							newReq.setType(ObjectSet)
							if err := newReq.insertReq(req.Value); err == nil {
								node.getMap()[req.Name] = newReq
							} else {
								resp.Error = errors.New("Insert failed: " + fmt.Sprint(err))
							}
						default:
							resp.Error = errors.New("Unknown Type")
						}
					}
				} else {
					resp.Error = errors.New("Node decoding failed: " + fmt.Sprint(err))
				}
			case OpGet:
				if node, err := rg.decodeNode(req.Node); err == nil {
					if reqSet, ok := node.getMap()[req.Name]; ok {
						resp.Value = reqSet.getReqs()
					} else {
						resp.Error = errors.New("Set of requirements not found")
					}
				} else {
					resp.Error = errors.New("Node decoding failed: " + fmt.Sprint(err))
				}
			default:
				resp.Error = errors.New("Unknown Operation")
			}

			rg.answer <- resp
		}
	}
}
