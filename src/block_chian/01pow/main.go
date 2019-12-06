package main

type block struct {
	//上一个区块的hash值
	preHash string
	//当前hash的值
	hashCode string
	//时间戳
	timeStamp string
	//当前网络难度系数
	//控制hash值前面有几个前导0
	diff int
	//交易信息
	data string
	//区块高度
	index int
	//区块高度
	nonce int
}

func genertFirstBlock(data string) block {

}
