package slicebigdigits

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var bigDigits = [][]string{
	{
		"   000   ",
		"  0   0  ",
		" 0     0 ",
		" 0     0 ",
		" 0     0 ",
		"  0   0  ",
		"   000   ",
	},
	{
		"   1   ",
		"  11   ",
		" 111   ",
		"   1   ",
		"   1   ",
		"   1   ",
		" 11111 ",
	},
	{
		"  222  ",
		" 2   2 ",
		"    2  ",
		"   2   ",
		"  2    ",
		" 2     ",
		" 22222 ",
	},
	{
		"  333  ",
		" 3   3 ",
		"     3 ",
		"   33  ",
		"     3 ",
		"     3 ",
		"  333  ",
	},
	{
		"    4   ",
		"   44   ",
		"  4 4   ",
		" 4  4   ",
		" 444444 ",
		"    4   ",
		"    4   ",
	},
	{
		" 55555 ",
		" 5     ",
		" 5     ",
		" 555   ",
		"     5 ",
		" 5   5 ",
		"  555  ",
	},
	{
		"  666   ",
		" 6      ",
		" 6      ",
		" 66666  ",
		" 6    6 ",
		" 6    6 ",
		"  6666  ",
	},
	{
		" 777777 ",
		"      7 ",
		"     7  ",
		"    7   ",
		"   7    ",
		"  7     ",
		" 7      ",
	},
	{
		"  8888  ",
		" 8    8 ",
		" 8    8 ",
		"  8888  ",
		" 8    8 ",
		" 8    8 ",
		"  8888  ",
	},
	{
		"  9999 ",
		" 9   9 ",
		" 9   9 ",
		"  9999 ",
		"     9 ",
		"     9 ",
		"  999  ",
	},
}

//0~6years
var yearsStr = [][]string{
	{
		"                             ",
		"                             ",
		"                             ",
		" 111111111111                ",
		"                             ",
		"                             ",
		"                             ",
	},
	{
		"                             ",
		"                             ",
		"                             ",
		"  2222222222                 ",
		"                             ",
		"                             ",
		"22222222222222               ",
	},
	{
		"    3333333                  ",
		"                             ",
		"                             ",
		"  3333333333                 ",
		"                             ",
		"                             ",
		"33333333333333               ",
	},
	{
		"                             ",
		"                             ",
		"  444444444444               ",
		" 4  4      4  4              ",
		" 4   4    4   4              ",
		"  4   4  4   4               ",
		"   4444444444                ",
	},
	{
		"  55555555555                ",
		"        5                    ",
		"       5                     ",
		"   5555555                   ",
		"     5   5                   ",
		"    5   5                    ",
		"5555555555555                ",
	},
	{
		"        66                   ",
		"                             ",
		"   666666666666              ",
		"      6    6                 ",
		"     6      6                ",
		"    6        6               ",
		"                             ",
	},
	{
		" 3   3   3                   ",
		" 3   3   3                   ",
		"  3333333                    ",
		"   33333                     ",
		"  3 3  3                     ",
		"     33                      ",
		"    3                        ",
	},
	{
		"  3333333  ",
		"       3   ",
		"      3    ",
		"      3    ",
		"      3    ",
		"   33 3    ",
		"    333    ",
	},
}

var happyBirthday = [][]string{
	{
		" H     H ",
		" H     H ",
		" H     H ",
		" HHHHHHH ",
		" H     H ",
		" H     H ",
		" H     H ",
	},
	{
		"         ",
		"         ",
		"   aaa   ",
		" a    aa ",
		"a     aa ",
		" a    aa ",
		"   aaa  a",
	},
	{
		"       ",
		"       ",
		"ppppp  ",
		"p    p ",
		"ppppp  ",
		"p      ",
		"p      ",
	},
	{
		"       ",
		"       ",
		"ppppp  ",
		"p    p ",
		"ppppp  ",
		"p      ",
		"p      ",
	},
	{
		"           ",
		"           ",
		"y     y    ",
		"y     y    ",
		" yyyyyy    ",
		"      y    ",
		" yyyyy     ",
	},
	{
		" BBBBB   ",
		" B     B ",
		" B    B  ",
		" BBBBB   ",
		" B    B  ",
		" B     B ",
		" BBBBB   ",
	},
	{
		"  ",
		"  ",
		"i ",
		"  ",
		"i ",
		"i ",
		"i ",
	},
	{
		"       ",
		"       ",
		"       ",
		"r rrrrr",
		"r      ",
		"r      ",
		"r      ",
	},
	{
		"      ",
		"      ",
		" t    ",
		" t    ",
		"ttttt ",
		" t    ",
		" tttt ",
	},
	{
		"     ",
		"     ",
		"h    ",
		"h    ",
		"hhhh ",
		"h  h ",
		"h  h ",
	},
	{
		"d d      ",
		"d    d   ",
		"d     d  ",
		"d      d ",
		"d     d  ",
		"d    d   ",
		"d d      ",
	},
	{
		"         ",
		"         ",
		"   aaa   ",
		" a    aa ",
		"a     aa ",
		" a    aa ",
		"   aaa  a",
	},
	{
		"           ",
		"           ",
		"y     y    ",
		"y     y    ",
		" yyyyyy    ",
		"      y    ",
		" yyyyy     ",
	},
	{
		"TTTTTTTTT",
		"    T    ",
		"    T   ",
		"    T   ",
		"    T   ",
		"    T   ",
		"    T   ",
	},
	{
		"          ",
		"          ",
		"   ooo     ",
		" o     o   ",
		"o       o  ",
		" o     o   ",
		"   ooo     ",
	},
	{
		"Y       Y",
		" Y     Y ",
		"  Y   Y  ",
		"   YYY ",
		"    Y  ",
		"    Y  ",
		"    Y  ",
	},
	{
		"         ",
		"         ",
		" ooo   ",
		" o     o ",
		"o       o",
		" o     o ",
		"   ooo   ",
	},
	{
		"          ",
		"          ",
		"u       u ",
		"u       u ",
		"u       u ",
		" u     u  ",
		"  uuuuu   ",
	},
}

var smileStrings = [][]string{
	{
		"                         ooooooooo                         ",
		"                    o                 o                    ",
		"                o                         o                ",
		"             o                               o             ",
		"          o                                     o          ",
		"        o     o o                         o o     o        ",
		"      o     o     o                     o     o     o      ",
		"    o     o         o                 o         o     o    ",
		"   o                                                   o   ",
		"  o                                                     o  ",
		" o                                                       o ",
		"o                                                         o",
		"o                                                         o",
		"o                                                         o",
		"o                                                         o",
		"o                                                         o",
		"o                                                         o",
		" o                                                       o ",
		"  o                                                     o  ",
		"   o            u                         u            o   ",
		"    o             u                     u             o    ",
		"      o             u                 u             o      ",
		"        o               uuuuuuuuuuu               o        ",
		"          o                                     o          ",
		"             o                               o             ",
		"                o                         o                ",
		"                    o                 o                    ",
		"                         ooooooooo                         ",
	},
	{
		"                                                           ",
		"                                                           ",
		"                                                           ",
		"                                                           ",
		"                                                           ",
		"                           ooooo                           ",
		"                       o     o     o                       ",
		"                    o   o    o    o    o                   ",
		"                  o      o   o   o       o                 ",
		"                o                         o                ",
		"               o    o o             o o    o               ",
		"              o   o     o         o     o   o              ",
		"             o      o o             o o      o             ",
		"             o                               o             ",
		"             o                               o             ",
		"              o                             o              ",
		"               o     ooooooooooooooooo     o               ",
		"                o     o             o                      ",
		"                  o      o       o      o                  ",
		"                    o        o        o                    ",
		"                       o           o                       ",
		"                            ooo                            ",
		"                                                           ",
		"                                                           ",
		"                                                           ",
		"                                                           ",
		"                                                           ",
		"                                                           ",
	},
	{
		"                                                       ",
		"                                                       ",
		"                     ooooooooo                         ",
		"                o                 o                    ",
		"            o                         o                ",
		"          o                             o              ",
		"        o                                 o            ",
		"      o                                     o          ",
		"    o      o o                       o o      o        ",
		"   o    o       o                 o       o    o       ",
		"  o                                             o      ",
		" o                                               o     ",
		"o                                                 o    ",
		"o                                                 o    ",
		"o                                                 o    ",
		" o                                               o     ",
		"  o           o                     o           o      ",
		"   o           o                   o           o       ",
		"    o            o               o            o        ",
		"     o             o           o             o         ",
		"       o              o  o  o              o           ",
		"         o                               o             ",
		"            o                         o                ",
		"               o                   o                   ",
		"                     ooooooooo                         ",
		"                                                       ",
		"                                                       ",
		"                                                       ",
	},
}

func SayBigDigits() {
	if len(os.Args) == 1 {
		defaultStrings := "20150515"
		printString(defaultStrings, bigDigits)
		SayTabLine(100)
		SayYears("3")

	} else {
		stringsOfDigits := os.Args[1]
		printString(stringsOfDigits, bigDigits)
	}
	if len(os.Args) == 3 {
		SayTabLine(100)
		SayYears(os.Args[2])
	}

}

func printString(stringsOfDigits string, bigDigits [][]string) {
	for row := range bigDigits[0] {
		line := ""
		for column := range stringsOfDigits {
			digit := stringsOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		time.Sleep(300 * time.Millisecond)
		fmt.Println(line)
	}
}

func SayTabLine(length int) {
	fmt.Println("                                                                                                ")
	for i := 0; i < length; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Print("#")
	}
	fmt.Println("                                                                                                ")
}

func SayHappyBirthDay() {

	for middle := range happyBirthday[0] {
		line := ""
		for i := 0; i < len(happyBirthday); i++ {
			line += happyBirthday[i][middle] + " "
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Println(line)
	}

}

func SayYears(yearStr string) {
	if yearInt, err := strconv.Atoi(yearStr); err != nil {
		log.Fatal("invalid years number")
	} else {
		for middle := range yearsStr[0] {
			line := ""
			if 0 <= yearInt && yearInt <= 9 {
				line += yearsStr[yearInt-1][middle] + " "
			} else {
				log.Fatal("invalid years number")
			}

			for i := len(yearsStr) - 2; i < len(yearsStr); i++ {
				line += yearsStr[i][middle] + " "
			}
			time.Sleep(50 * time.Millisecond)
			fmt.Println(line)
		}
	}
}

func DisplaySmileFace() {

	for middle := range smileStrings[0] {
		line := ""
		for i := 0; i < len(smileStrings); i++ {
			line += smileStrings[i][middle] + " "
		}
		time.Sleep(300 * time.Millisecond)
		fmt.Println(line)
	}

}

func DisplayAll() {
	SayTabLine(80)
	SayBigDigits()
	SayTabLine(150)
	SayHappyBirthDay()
	SayTabLine(200)
	DisplaySmileFace()
}
