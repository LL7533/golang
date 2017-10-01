// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var stringB = flag.String("-bar", "-b", "input ur bar")

func main() {

	flag.Parse()
	fmt.Println("name=", *stringB)

	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if len(os.Args) == 3 {
		stringOfB := os.Args[1]
		if stringOfB == "-b" || stringOfB == "-bar" {
			stringOfDigits := os.Args[2]
			for row := range bigDigits[0] {
				line := ""
				xingTest := ""
				for column := range stringOfDigits {
					digit := stringOfDigits[column] - '0'
					if 0 <= digit && digit <= 9 {
						if row == 0 || row == len(bigDigits[0])-1 {

							for i := 1; i <= len([]rune(bigDigits[digit][row]+"  ")); i++ {
								xingTest += xing
							}

						}
						line += bigDigits[digit][row] + "  "
					} else {
						log.Fatal("invalid whole number")
					}
				}
				if row == 0 {
					fmt.Println(xingTest)
				}
				fmt.Println(line)
				if row == len(bigDigits[0])-1 {
					fmt.Println(xingTest)
				}
			}
		} else {
			fmt.Println("invalid whole  -b ")
		}

	}
	//	stringOfB  := os.Args[1];

}

var xing string = "&"

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
		"   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
