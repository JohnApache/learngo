package main

import (
	"mypkg/Fibonacci/mapTree"
	"strconv"

	"gopkg.in/ffmt.v1"
)

var TestString = "11122233355588814432330"
var mt = mapTree.NewMT("0")

//11121131152283385588881
func main() {
	Split(TestString)
	res := GetPosibilityMT(mt)
	for _, data := range res {
		res := IsFibonacci(data[1:])
		if res == true {
			ffmt.Mark(data[1:])
		}
	}
}

func Split(str string) {
	for i := 1; i < len(str); i++ { //第一层
		if len(str[:i]) > len(str[i:]) {
			break
		}
		mt.AddNode(str[:i])
		chmt := mt.GetNode()
		for j := 1; j < len(str[i:]); j++ { //第二层
			if len(str[:i]) > len(str[i+j:]) || len(str[i:j+i]) > len(str[i+j:]) {
				break
			}
			chmt[i-1].AddNode(str[i : j+i])
			cchmt := chmt[i-1].GetNode()
			UnitSplit(str[j+i:], str[i:j+i], cchmt[j-1])
		}
	}
}

func UnitSplit(unit, prev string, n *mapTree.Node) {
	l := len(unit)  //剩余长度
	lp := len(prev) //前一个长度
	if l < lp {
		return
	}
	n.AddNode(unit[:lp])
	chmt := n.GetNode()
	UnitSplit(unit[lp:], unit[:lp], chmt[0])
	if lp+1 <= l {
		n.AddNode(unit[:lp+1])
		chmt = n.GetNode()
		UnitSplit(unit[lp+1:], unit[:lp+1], chmt[1])
	}
}

func GetPosibilityMT(mt *mapTree.MT) [][]interface{} {
	res := make([][]interface{}, 0)
	GetPosibilityNode(mt.Root, nil, &res)
	return res
}

func GetPosibilityNode(n *mapTree.Node, res []interface{}, result *[][]interface{}) {
	if res == nil {
		res = []interface{}{}
	}
	res = append(res, n.Content)
	chmt := n.GetNode()
	if len(chmt) == 0 {
		*result = append(*result, res)
		return
	}
	for _, d := range chmt {
		r := make([]interface{}, len(res))
		copy(r, res)
		GetPosibilityNode(d, r, result)
	}
}

/*
	追踪删除
*/
//func SpiderDelete(parent *Node) {

//}

func IsFibonacci(fib []interface{}) bool {
	if len(fib) < 3 {
		return false
	}
	res := true
	for i := 2; i < len(fib); i++ {
		if AssertInt(fib[i]) != AssertInt(fib[i-1])+AssertInt(fib[i-2]) {
			res = false
			break
		}
	}
	return res
}

func AssertInt(t interface{}) int64 {
	str, _ := t.(string)
	res, _ := strconv.ParseInt(str, 10, 0)
	return res
}
