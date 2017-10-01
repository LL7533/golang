package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var baseData = [][]string{
	{
		" 000 ",
		"0   0",
		"0   0",
		"0   0",
		"0   0",
		"0   0",
		" 000 ",
	},
	{
		"  1  ",
		"1 1  ",
		"  1  ",
		"  1  ",
		"  1  ",
		"  1  ",
		"11111",
	},
	{
		" 222 ",
		"2   2",
		"   2 ",
		"  2  ",
		" 2   ",
		"2    ",
		"22222",
	},
	{
		" 333 ",
		"3   3",
		"   3 ",
		"  3  ",
		"   3 ",
		"3   3",
		" 333 ",
	},
	{
		"   4 ",
		"  44 ",
		" 4 4 ",
		"4  4 ",
		"44444",
		"   4 ",
		"   4 ",
	},
}

func main() {
	//判断取值看是否有参数
	if len(os.Args) == 1 {
		fmt.Printf("userArgs:%s<whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	stringInsert := os.Args[1]
	for row := range baseData[0] {
		line := ""
		for column := range stringInsert[0] {
			digit = stringInsert[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += baseData[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
		}
	}
	fmt.Println(line)
}
