package action

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	files, _ := ioutil.ReadDir(util.RootPath() + "/volumes/storage/action/system_image/")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func Test(t *testing.T) {
	// 假設從資料夾讀出 input1 ~ input3 txt 資料
	inputTxt1 := "3,962,661,39,23\n" +
		"3,964,670,39,22\n" +
		"0,1064,683,37,44\n" +
		"3,1056,968,63,41\n" +
		"0,1227,551,69,35\n" +
		"0,1453,799,51,43\n" +
		"0,1315,614,91,42\n" +
		"0,1330,800,45,39\n" +
		"0,1515,799,50,44\n" +
		"0,1271,796,45,42\n" +
		"0,1394,798,44,38\n" +
		"0,1292,578,84,43\n" +
		"0,1162,686,43,37\n" +
		"1,1004,815,82,95\n" +
		"0,1212,682,45,42\n" +
		"3,1161,535,27,16\n" +
		"2,1345,564,13,20\n" +
		"3,1324,563,8,20"

	inputTxt2 := "3,962,661,39,23\n" +
		"3,964,670,39,22\n" +
		"0,1064,683,37,44\n" +
		"3,1324,563,8,20"

	inputTxt3 := "3,962,661,39,23\n" +
		"3,964,670,39,22\n" +
		"0,1064,683,37,44\n" +
		"3,1324,563,8,20"
	inputs := make([]string, 0)
	inputs = append(inputs, inputTxt1)
	inputs = append(inputs, inputTxt2)
	inputs = append(inputs, inputTxt3)

	// 將 input1 ~ input3 用 handler 函數正規化並將結果放入 dataList 資料集中，方便進行後續分析
	dataList := make([][][]float64, 0)
	for _, input := range inputs{
		outputs := handler(input)
		dataList = append(dataList, outputs)
	}
	// 正規化後的資料集
	fmt.Println(dataList)
}

// 正規化函數
func handler(input string) (outputs [][]float64) {
	// 以斷行符號拆分字串行資料
	rows := strings.Split(input, "\n")
	outputs = make([][]float64, 0)
	for _, row := range rows {
		output := make([]float64, 0)
		// 以逗號拆分字串行欄位資料
		columns := strings.Split(row, ",")
		// 取出各欄位數值
		columns0, _ := strconv.ParseFloat(columns[0], 64)
		columns1, _ := strconv.ParseFloat(columns[1], 64)
		columns2, _ := strconv.ParseFloat(columns[2], 64)
		columns3, _ := strconv.ParseFloat(columns[3], 64)
		columns4, _ := strconv.ParseFloat(columns[4], 64)
		// 第0欄不變
		output = append(output, columns0)
		// 第1欄=(第1欄+第3欄/2)/1920
		output = append(output, (columns1+columns3/2.0)/1920.0)
		// 第2欄=(第2欄+第4欄/2)/1080
		output = append(output, (columns2+columns4/2.0)/1080.0)
		// 第3欄=第3欄/1920
		output = append(output, columns3/1920.0)
		// 第4欄=第4欄/1080
		output = append(output, columns4/1080.0)
		// 將結果放入 outputs array 中
		outputs = append(outputs, output)
	}
	return outputs
}
