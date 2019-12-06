package main

import (
	"testblockchain/Block"
	"testblockchain/BlockChain"
)

func main() {
	first := Block.GenertFirstBlock("创世区块")
	//生成下个区块
	second := Block.GengertNewBlock("第二个区块", first)
	//创建链表
	header := BlockChain.CreateHeaderNode(&first)
	//将第二个区块加入链表

	BlockChain.AddNode(&second, header)
	//查看链表的数据

	BlockChain.ShowNode(header)
}
