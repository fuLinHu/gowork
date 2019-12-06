package BlockChain

import (
	"fmt"
	"testblockchain/Block"
)

type Node struct {
	//指针
	NextNode *Node
	//数据区域
	Data *Block.Block
}

//创建头节点

func CreateHeaderNode(data *Block.Block) *Node {
	//创建头节点
	headerNode := new(Node)
	//头节点的指针与  指向nil
	headerNode.NextNode = nil
	//数据区域保存传入的data
	headerNode.Data = data
	return headerNode

}

//添加节点
//当挖矿成功添加区块
func AddNode(data *Block.Block, node *Node) *Node {
	//创建新节点
	var newNode *Node = new(Node)
	newNode.NextNode = nil
	newNode.Data = data
	node.NextNode = newNode
	return newNode

}

//查看列表数据
func ShowNode(node *Node) {
	n := node
	for {
		//如果有就打印
		if n.NextNode == nil {
			//没有下个节点 打印本节点
			fmt.Println("本节点：", n.Data)
			break
		} else {
			fmt.Println("本节点：", n.Data)
			n = n.NextNode
		}

	}
}
