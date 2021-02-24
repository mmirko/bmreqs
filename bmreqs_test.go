package bmreqs

import (
	"fmt"
	"testing"
)

func TestBMReq(t *testing.T) {

	rg := NewReqGroup()

	// Some errors
	fmt.Println(rg.Requirement(ReqRequest{Node: "/", T: ObjectSet, Name: "processors", Op: 34}))

	fmt.Println(rg.Requirement(ReqRequest{Node: "/", T: ObjectSet, Name: "processors", Value: "cp0", Op: OpAdd}))
	fmt.Println(rg.Requirement(ReqRequest{Node: "/", T: ObjectSet, Name: "processors", Value: "cp1", Op: OpAdd}))
	fmt.Println(rg.Requirement(ReqRequest{Node: "/", T: ObjectSet, Name: "processors", Value: "cp2", Op: OpAdd}))
	fmt.Println(rg.Requirement(ReqRequest{Node: "/", T: ObjectSet, Name: "processors", Value: "cp3", Op: OpAdd}))
	fmt.Println(rg.Requirement(ReqRequest{Node: "/", T: ObjectSet, Name: "processors", Value: "cp4", Op: OpAdd}))
	fmt.Println(rg.Requirement(ReqRequest{Node: "/", Name: "processors", Op: OpGet}))

	fmt.Println(rg.Requirement(ReqRequest{Node: "/processors:cp0", T: ObjectSet, Name: "opcodes", Value: "rset", Op: OpAdd}))
	fmt.Println(rg.Requirement(ReqRequest{Node: "/processors:cp0", Name: "opcodes", Op: OpGet}))

	rg.Close()

}
