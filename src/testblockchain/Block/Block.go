package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	preHash   string
	hashCode  string
	timeStamp string
	diff      int
	data      string
	index     int
	nonce     int
}

func GenertFirstBlock(data string) Block {
	var firstblock Block
	firstblock.preHash = "0"
	firstblock.timeStamp = time.Now().String()
	firstblock.diff = 4
	firstblock.data = data
	firstblock.index = 1
	firstblock.nonce = 0
	firstblock.hashCode = gengerhash(firstblock)
	return firstblock
}

func gengerhash(block2 Block) string {
	var hashdata = strconv.Itoa(block2.index) + strconv.Itoa(block2.nonce) + strconv.Itoa(block2.diff) + block2.timeStamp
	var hash = sha256.New()
	hash.Write([]byte(hashdata))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

func GengertNewBlock(data string, oldblock Block) Block {
	var newblock Block
	newblock.preHash = oldblock.hashCode
	newblock.timeStamp = time.Now().String()
	//前到0为4
	newblock.diff = 10
	newblock.data = data
	newblock.index = 2
	//由旷工来调整
	newblock.nonce = 0
	//利用pow挖矿
	newblock.hashCode = pow(newblock.diff, &newblock)
	return newblock
}

//pow工作量证明算法进行hash 碰撞
//传指针保证是同一个对象
func pow(diff int, block1 *Block) string {
	for {
		s := gengerhash(*block1)
		fmt.Println(s)
		if strings.HasPrefix(s, strings.Repeat("0", diff)) {
			fmt.Println("挖矿成功")
			return s
		} else {
			block1.nonce++
		}
	}

}
