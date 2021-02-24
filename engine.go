package bmreqs

import (
	"errors"
	"fmt"
	"strings"
)

func (rg *ReqGroup) decodeNode(node string) (bmReqSet, bmReqSet, string, string, error) {
	if node == "/" {
		return rg, nil, "", "", nil
	}

	seqlist := strings.Split(node, "/")
	var currNode bmReqSet = rg
	var prevNode bmReqSet = nil
	var lName string
	var lReq string
	for _, leaf := range seqlist {
		currMap := currNode.getSubReqMap()
		if leaf != "" {
			decLeaf := strings.Split(leaf, ":")
			if len(decLeaf) != 2 {
				return nil, nil, "", "", errors.New("Malformed node")
			}
			lName = decLeaf[0]
			lReq = decLeaf[1]

			if reqset, ok := currMap[lName]; ok {
				if reqset.supportSubReq() {
					if reqset.matchReq(lReq) {
						prevNode = currNode
						currNode = reqset.getSubReq(lReq)
					} else {
						return nil, nil, "", "", errors.New("Requirement not exists")
					}
				} else {
					return nil, nil, "", "", errors.New("Requirement Set does not support sub requirements")
				}
			} else {
				return nil, nil, "", "", errors.New("Unknown Requirement Set")
			}
		}
	}

	return currNode, prevNode, lName, lReq, nil
}

func (rg *ReqGroup) run() {
	ctx := rg.ctx
	for {
		select {
		case <-ctx.Done():
			return
		case req := <-rg.ask:
			resp := ReqResponse{Value: "", Error: nil}
			switch req.Op {
			case OpAdd:
				if node, prev, lreqSet, lNode, err := rg.decodeNode(req.Node); err == nil {
					if node == nil {
						switch req.T {
						case ObjectSet:
							newReq := new(objectSet)
							newReq.setName(req.Name)
							newReq.setType(ObjectSet)
							if err := newReq.insertReq(req.Value); err == nil {
								prev.setSubReq(lreqSet, lNode, newReq)
							} else {
								resp.Error = errors.New("Insert failed: " + fmt.Sprint(err))
							}
						default:
							resp.Error = errors.New("Unknown Type")
						}
					} else {
						currMap := node.getSubReqMap()
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
								newReq.setName(req.Name)
								newReq.setType(ObjectSet)
								if err := newReq.insertReq(req.Value); err == nil {
									node.getSubReqMap()[req.Name] = newReq
								} else {
									resp.Error = errors.New("Insert failed: " + fmt.Sprint(err))
								}
							default:
								resp.Error = errors.New("Unknown Type")
							}
						}
					}
				} else {
					resp.Error = errors.New("Node decoding failed: " + fmt.Sprint(err))
				}
			case OpGet:
				if node, _, _, _, err := rg.decodeNode(req.Node); err == nil {
					if reqSet, ok := node.getSubReqMap()[req.Name]; ok {
						resp.Value = reqSet.getReqs()
					} else {
						resp.Error = errors.New("Set of requirements not found" + node.getName())
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
