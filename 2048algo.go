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
	fmt.Printf("----- EMPTY GRID: -----\n")
	printGrid(g)
	g = initGrid(g)
	fmt.Printf("----- INIT GRID: -----\n")
	printGrid(g)
	g = swipe_right(g)
	fmt.Printf("----- SWIPE RIGHT: -----\n")
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

func swipe_right(g map[string]int) map[string]int {
	// by swiping right, we move all numbers to the right, combining them when the number to their right is the same
	var a [4]int
	var r [4]int
	for y:=0; y<4; y++ {
		a[0] = g[strconv.Itoa(0)+","+strconv.Itoa(y)]
		a[1] = g[strconv.Itoa(1)+","+strconv.Itoa(y)]
		a[2] = g[strconv.Itoa(2)+","+strconv.Itoa(y)]
		a[3] = g[strconv.Itoa(3)+","+strconv.Itoa(y)]
		fmt.Printf("On Y access of " + strconv.Itoa(y)+"\n")
		r = determine_swipe_result(a)
		for z := 0; z<4; z++ {
			//fmt.Printf("Returned for X of " + strconv.Itoa(z) + " with value of access of " + strconv.Itoa(r[z])+"\n")
			g[strconv.Itoa(z)+","+strconv.Itoa(y)] = r[z]
		}
	}
	return g
}

func determine_swipe_result(a [4]int) [4]int {
	//always pushes toward D
	ar,br,cr,dr :=0,0,0,0
	//find num same right next to each other non-zero. Add op needed then push right.
	   // if 1==2 & 3==4 case
	//else push right
	
	// number of zeroes in the grid
	zcount :=0
	for x := 0; x<4; x++ {
		if a[x] == 0 {
			zcount++
		}
	}

	if zcount == 4 {
		return [4]int{ar,br,cr,dr} 
	} else if zcount == 3 {
		return [4]int{ar,br,cr,a[0] + a[1] + a[2] + a[3]}
	}

	if (a[0] != a[1]) && (a[1] != a[2]) && (a[2] != a[3]) {
		if a[3] == 0 {
			ar,br,cr,dr := 0,a[0],a[1],a[2]
			return [4]int{ar,br,cr,dr}
		} else if a[2] == 0 {
			ar,br,cr,dr := 0,a[0],a[1],a[3]
			return [4]int{ar,br,cr,dr}
		} else if a[1] == 0 {
			ar,br,cr,dr := 0,a[0],a[2],a[3]
			return [4]int{ar,br,cr,dr}
		}  else {
			ar,br,cr,dr := a[0],a[1],a[2],a[3]
			return [4]int{ar,br,cr,dr}
		}
	}


	
	if a[3] == a[2] {
		dr = a[3] + a[2]
	}
	if (a[1] == a[2]) && (dr == 0) {
		cr = a[1] + a[2]
	}
	if (a[0] == a[1]) && (cr == 0) {
		br=a[0] + a[2]
	}
	if (a[3] != 0) && (dr != 0) {
		dr = a[2]
	}

    return [4]int{ar,br,cr,dr}
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