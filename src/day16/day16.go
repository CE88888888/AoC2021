package main

import (
	"encoding/hex"
	"fmt"
)

type Header struct {
	version int
	id      int
}

func NewHeader(data []int) Header {
	h := Header{
		version: getVersion(data),
		id:      getID(data),
	}

	return h
}

type Packet struct {
	h          Header
	isLiteral  bool
	versionsum int
	length     int
	subpackets []Packet
	value      int
	groups     [][]int
}

type Literal struct {
}

func (p *Packet) Sum() int {
	i := 0
	for _, v := range p.subpackets {
		i += v.value
	}
	return i
}

func (p *Packet) Product() int {
	sum := p.subpackets[0].value
	for i := 1; i < len(p.subpackets); i++ {
		sum = sum * p.subpackets[i].value
	}
	return sum
}

func (p *Packet) Min() int {
	i := p.subpackets[0].value
	for _, v := range p.subpackets {
		if i > v.value {
			i = v.value
		}
	}
	return i
}

func (p *Packet) Max() int {
	i := 0
	for _, v := range p.subpackets {
		if i < v.value {
			i = v.value
		}
	}
	return i
}

func (p *Packet) Greater() int {
	if p.subpackets[0].value > p.subpackets[1].value {
		return 1
	} else {
		return 0
	}
}

func (p *Packet) Less() int {
	if p.subpackets[0].value < p.subpackets[1].value {
		return 1
	} else {
		return 0
	}
}
func (p *Packet) Equal() int {
	if p.subpackets[0].value == p.subpackets[1].value {
		return 1
	} else {
		return 0
	}
}

func main() {
	rawHex := "220D790065B2745FF004672D99A34E5B33439D96CEC80373C0068663101A98C406A5E7395DC1804678BF25A4093BFBDB886CA6E11FDE6D93D16A100325E5597A118F6640600ACF7274E6A5829B00526C167F9C089F15973C4002AA4B22E800FDCFD72B9351359601300424B8C9A00BCBC8EE069802D2D0B945002AB2D7D583E3F00016B05E0E9802BA00B4F29CD4E961491CCB44C6008E80273C393C333F92020134B003530004221347F83A200D47F89913A66FB6620016E24A007853BE5E944297AB64E66D6669FCEA0112AE06009CAA57006A0200EC258FB0440010A8A716A321009DE200D44C8E31F00010887B146188803317A3FC5F30056C0150004321244E88C000874468A91D2291802B25EB875802B28D13550030056C0169FB5B7ECE2C6B2EF3296D6FD5F54858015B8D730BB24E32569049009BF801980803B05A3B41F1007625C1C821256D7C848025DE0040E5016717247E18001BAC37930E9FA6AE3B358B5D4A7A6EA200D4E463EA364EDE9F852FF1B9C8731869300BE684649F6446E584E61DE61CD4021998DB4C334E72B78BA49C126722B4E009C6295F879002093EF32A64C018ECDFAF605989D4BA7B396D9B0C200C9F0017C98C72FD2C8932B7EE0EA6ADB0F1006C8010E89B15A2A90021713610C202004263E46D82AC06498017C6E007901542C04F9A0128880449A8014403AA38014C030B08012C0269A8018E007A801620058003C64009810010722EC8010ECFFF9AAC32373F6583007A48CA587E55367227A40118C2AC004AE79FE77E28C007F4E42500D10096779D728EB1066B57F698C802139708B004A5C5E5C44C01698D490E800B584F09C8049593A6C66C017100721647E8E0200CC6985F11E634EA6008CB207002593785497652008065992443E7872714"
	input := hexToIntArray(rawHex)
	packet := readMessage(input, 0)
	fmt.Println("Version Sum:", packet.versionsum, "Length:", packet.length, "Value:", packet.value)
}

func readMessage(input []int, index int) (p Packet) {
	id := getID(input[index:])
	if id == 4 {
		p = getGroups((input)[index:])
	} else {
		p = getSubpackets((input)[index:])
		switch id {
		case 0:
			p.value = p.Sum()
		case 1:
			p.value = p.Product()
		case 2:
			p.value = p.Min()
		case 3:
			p.value = p.Max()
		case 5:
			p.value = p.Greater()
		case 6:
			p.value = p.Less()
		case 7:
			p.value = p.Equal()
		}
	}
	return p
}

func getSubpackets(input []int) (p Packet) {

	p.isLiteral = false
	p.h = NewHeader(input)
	p.versionsum = p.h.version

	lengthType := input[6]
	noSP, totalLength := 0, 0

	if lengthType == 0 {
		lengthBits := input[7:22]
		totalLength = int(bitsToByte(lengthBits))
	} else {
		lengthBits := input[7:18]
		noSP = int(bitsToByte(lengthBits))
	}
	if noSP < 1 {
		for p.length < totalLength {
			sp := readMessage(input[22+p.length:22+totalLength], 0)
			p.versionsum += sp.versionsum
			p.length += sp.length
			p.subpackets = append(p.subpackets, sp)
		}
		p.length += 15

	} else {
		for i := 0; i < noSP; i++ {
			sp := readMessage(input[18+p.length:], 0)
			p.versionsum += sp.versionsum
			p.length += sp.length
			p.subpackets = append(p.subpackets, sp)
		}
		p.length += 11
	}
	p.length += 7
	return p
}

func getGroups(input []int) (p Packet) {
	p.h = NewHeader(input)
	p.isLiteral = true

	searching := true
	literals := []int{}

	for i := 6; searching; i = i + 5 {
		literals = append(literals, input[(i+1):(i+5)]...)
		p.groups = append(p.groups, input[(i+1):(i+5)])
		if input[i] == 0 {
			searching = false
		}
	}
	p.versionsum = p.h.version
	p.length = (6 + (len(p.groups) * (len(p.groups[0]) + 1)))
	p.value = int(bitsToByte(literals))
	return
}

func hexToIntArray(rawHex string) []int {
	inData, err := hex.DecodeString(rawHex)
	if err != nil {
		panic("shit")
	}
	input := []int{}
	for _, v := range inData {
		input = append(input, byteToBitArray(v)...)
	}
	return input
}

func getID(input []int) int {
	idBits := input[3:6]
	id := bitsToByte(idBits)
	return int(id)
}

func getVersion(i []int) int {
	versionBits := i[0:3]
	version := bitsToByte(versionBits)
	return int(version)

}

func bitsToByte(b []int) int64 {
	var theByte int64
	theByte = 0
	size := len(b)
	for bit, mask := size-1, 1; bit >= 0; bit-- {
		if b[bit] != 0 {
			theByte |= int64(mask)
		}
		mask <<= 1
	}
	return theByte
}

func byteToBitArray(data byte) []int {
	bitArray := [8]int{}
	for bit, mask := 7, 1; bit >= 0; bit-- {
		if int(data)&mask != 0 {
			bitArray[bit] = 1
		}
		mask <<= 1
	}
	return bitArray[:]
}
