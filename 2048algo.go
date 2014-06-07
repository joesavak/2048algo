package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)


func main() {
	//prop = get_props()
	
	var g= map[string]int {
		"0,0": 0,
		"1,0": 0,
		"2,0": 0,
		"3,0": 0,
		"0,1": 0,
		"1,1": 0,
		"2,1": 0,
		"3,1": 0,
		"0,2": 0,
		"1,2": 0,
		"2,2": 0,
		"3,2": 0,
		"0,3": 0,
		"1,3": 0,
		"2,3": 0,
		"3,3": 0,
	}
	printGrid(g)
	fmt.Printf("----- EMPTY GRID: -----\n")
	g = initGrid(g)
	fmt.Printf("----- INIT GRID: -----\n")
	printGrid(g)
}

func randomNum() int {
	rand.Seed(time.Now().UTC().UnixNano())
	digit := []int {
		2,
		4,
	}
	return digit[rand.Intn(len(digit))]
}


func getRandomOpenPos(g  map[string]int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	a := []string{}
	for k  := range g {
		if g[k] == 0 {
			a=append(a,k)
		}
	}
	return a[rand.Intn(len(a))]

	
}

func putInGrid(g map[string]int, v int) map[string]int {
	pos := getRandomOpenPos(g)
	g[pos] = v
	return g
}

func initGrid(g  map[string]int) map[string]int {
	num1:= randomNum()
	time.Sleep(10 * time.Millisecond)
	num2:= randomNum()
	fmt.Printf(strconv.Itoa(num1)+"\n")
	g = putInGrid(g, num1)
	g = putInGrid(g, num2)
	return g
}

func printGrid(g map[string]int) {
	val:= " "
	for y:=0; y<4; y++ {
		for x:=0; x<4; x++ {
			key:=strconv.Itoa(x)+","+strconv.Itoa(y)
			if g[key] == 0 {
				val= " "
			} else {
				val= strconv.Itoa(g[key])
			}
			if x == 3 {
					fmt.Printf("|"+val+"|\n")
				} else {
					fmt.Printf("|"+val+"|")
			}
			
		}
	}
}

func get_props() {
	
}

func swipe_right(g *map[string]int) {
	fmt.Printf("Going Right\n")
}

func swipe_left(g *map[string]int) {
	fmt.Printf("Going Left\n")
}

func swipe_down(g *map[string]int) {
	fmt.Printf("Going Down\n")
}

func swipe_up(g *map[string]int) {
	fmt.Printf("Going Up\n")
}