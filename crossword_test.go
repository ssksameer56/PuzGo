package puzgo

import (
	"fmt"
	"testing"

	"github.com/ssksameer56/PuzGo/utils"
)

func TestIsAcrossClueNumber(t *testing.T) {

}

func TestIsDownClueNumber(t *testing.T) {

}

func TestDistanceAcross(t *testing.T) {

}

func TestDistanceDown(t *testing.T) {

}

func TestSplitByString(t *testing.T) {
	data := []byte{65, 66, 67, 45, 0, 45, 65, 66}
	delim := byte(0)
	stringd, err := utils.SplitByteString(data, delim)
	if err != nil {
		t.Errorf("couldnt split %s", err)
	}
	if stringd[0] != "ABC-" {
		t.Errorf("string not split properly %s", stringd[0])
	}
	if stringd[1] != "-AB" {
		t.Errorf("string not split properly %s", stringd[1])
	}
}

func TestStringConvert(t *testing.T) {
	data := []byte("1.3")
	data = append(data, 0x00)
	string, _ := utils.ConvertToString(data)
	fmt.Println(string)
}

func TestCrossword(t *testing.T) {
	fp := "./test.puz"
	cw := Crossword{
		filePath: fp,
	}
	flag, err := cw.ParseCrossword()
	if err != nil || !flag {
		t.Errorf("couldnt parse crossword %s", err)
	}
}
