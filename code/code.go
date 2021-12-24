package code 

import (
	"encoding/binary"
	"fmt"
)

type Instructions [] byte //字节码集合

type Opcode byte  //操作码

const (
	OpConstant Opcode = iota
	OpPop  //弹出堆栈
	OpAdd  //将栈顶两个数弹出并相加，把结果压入堆栈
	OpSub
	OpMul 
	OpDiv  
	//后面还有更多操作码需要定义
)

type Definition struct {
	Name string //操作码对应的字符串
	OperandWidths []int //操作数对应的字节数
}

var definitions = map[Opcode]*Definition {
	OpConstant: {"OpConstant", []int{2}} ,
}

func Lookup(op byte) (*Definition, error ) {
	//给定操作码，返回它对应的信息定义
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func Make(op Opcode, operands ...int) []byte {
	//给定操作码，创建字节码指令
	def , ok := definitions[op]
	if !ok {
		return []byte{}
	}

	//一条指令的字节长度包括操作码对应的长度加上操作数对应的长度
	instructionLen :=1  //操作码长度始终为1
	for _, w := range def.OperandWidths {
		instructionLen += w 
	}
    //一条指令由一系列字节组成,第一个字节就是操作码，接下来的字节对应操作数
	instructions := make([]byte, instructionLen) 
	instructions[0] = byte(op) //设置操作码对应的字节
	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			//把一个16比特数,也就是uint16类型的数值拆解成2个byte放到数组中
			binary.BigEndian.PutUint16(instructions[offset:], uint16(o))
		}
		offset += width 
	}

	return instructions
}
