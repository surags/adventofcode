package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//file, err := os.Open("sample.in")
	//file, err := os.Open("sample2.in")
	//file, err := os.Open("sample3.in")
	file, err := os.Open("dec15.in")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var input []string

	for scanner.Scan() {
		v := scanner.Text()
		input = append(input, v)
	}

	file.Close()

	warehouse, wareHouse2, moves, r, r2 := parseInput(input)

	sumGPSCoords := part1(warehouse, moves, r)
	fmt.Println("Part 1: Sum of all GPS Coordinates: " + strconv.Itoa(sumGPSCoords))

	sumGPSCoords2 := part2(wareHouse2, moves, r2)
	fmt.Println("Part 2: Sum of all GPS Coordinates: " + strconv.Itoa(sumGPSCoords2))
}

type robot struct {
	posX, posY int
}

func parseInput(input []string) ([][]string, [][]string, []string, robot, robot) {
	warehouse := [][]string{}
	warehouse2 := [][]string{}
	r := robot{}
	r2 := robot{}
	ind := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			ind = i
			break
		}

		m := []string{}
		m2 := []string{}
		for j, s := range input[i] {
			if string(s) == "@" {
				r.posX = i
				r.posY = j
			}
			m = append(m, string(s))

			// part2
			if string(s) == "@" {
				r2.posX = i
				r2.posY = 2 * j
				m2 = append(m2, string(s))
				m2 = append(m2, ".")
			} else if string(s) == "O" {
				m2 = append(m2, "[", "]")
			} else if string(s) == "#" {
				m2 = append(m2, "#", "#")
			} else {
				m2 = append(m2, string(s), string(s))
			}
		}
		warehouse = append(warehouse, m)
		warehouse2 = append(warehouse2, m2)
	}
	ind++

	moves := []string{}
	for i := ind; i < len(input); i++ {
		for _, m := range input[i] {
			moves = append(moves, string(m))
		}
	}

	return warehouse, warehouse2, moves, r, r2
}

func moveObject(warehouse [][]string, obj string, curX, curY int, move string) bool {
	if move == "^" {
		x, y := curX-1, curY
		// up
		if warehouse[x][y] == "#" {
			// cant move
			return false
		} else if warehouse[x][y] == "." {
			// can move
			return true
		} else if warehouse[x][y] == "@" {
			if obj == "O" {
				// move possible since @ will move
				return true
			} else {
				return false
			}
		} else {
			// move box O
			canPush := moveObject(warehouse, warehouse[x][y], x, y, move)
			if canPush {
				warehouse[x-1][y] = warehouse[x][y]
				warehouse[x][y] = "."
				return true
			} else {
				return false
			}
		}
	} else if move == "v" {
		// down
		x, y := curX+1, curY
		// up
		if warehouse[x][y] == "#" {
			// cant move
			return false
		} else if warehouse[x][y] == "." {
			// can move
			return true
		} else if warehouse[x][y] == "@" {
			if obj == "O" {
				// move possible since @ will move
				return true
			} else {
				return false
			}
		} else {
			// move box O
			canPush := moveObject(warehouse, warehouse[x][y], x, y, move)
			if canPush {
				warehouse[x+1][y] = warehouse[x][y]
				warehouse[x][y] = "."
				return true
			} else {
				return false
			}
		}
	} else if move == ">" {
		// right
		x, y := curX, curY+1
		// up
		if warehouse[x][y] == "#" {
			// cant move
			return false
		} else if warehouse[x][y] == "." {
			// can move
			return true
		} else if warehouse[x][y] == "@" {
			if obj == "O" {
				// move possible since @ will move
				return true
			} else {
				return false
			}
		} else {
			// move box O
			canPush := moveObject(warehouse, warehouse[x][y], x, y, move)
			if canPush {
				warehouse[x][y+1] = warehouse[x][y]
				warehouse[x][y] = "."
				return true
			} else {
				return false
			}
		}
	} else if move == "<" {
		// left
		x, y := curX, curY-1
		// up
		if warehouse[x][y] == "#" {
			// cant move
			return false
		} else if warehouse[x][y] == "." {
			// can move
			return true
		} else if warehouse[x][y] == "@" {
			if obj == "O" {
				// move possible since @ will move
				return true
			} else {
				return false
			}
		} else {
			// move box O
			canPush := moveObject(warehouse, warehouse[x][y], x, y, move)
			if canPush {
				warehouse[x][y-1] = warehouse[x][y]
				warehouse[x][y] = "."
				return true
			} else {
				return false
			}
		}
	} else {
		fmt.Println("Invalid move: " + move)
		return false
	}
}

func part1(warehouse [][]string, moves []string, r robot) int {
	sumGPSCoords := 0
	for _, m := range moves {
		canMove := moveObject(warehouse, warehouse[r.posX][r.posY], r.posX, r.posY, m)
		if canMove {
			warehouse[r.posX][r.posY] = "."
			if m == "^" {
				r.posX = r.posX - 1
			} else if m == "v" {
				r.posX = r.posX + 1
			} else if m == ">" {
				r.posY = r.posY + 1
			} else if m == "<" {
				r.posY = r.posY - 1
			}
			warehouse[r.posX][r.posY] = "@"

		}

		//fmt.Printf("Map %d, moved: %v, move: %s \n", i, canMove, m)
		//for _, m := range warehouse {
		//	for _, s := range m {
		//		fmt.Printf("%s", s)
		//	}
		//	fmt.Printf("\n")
		//}
	}

	for i, s := range warehouse {
		for j, s2 := range s {
			if s2 == "O" {
				sumGPSCoords += 100*i + j
			}
		}
	}

	//fmt.Println("Map")
	//for _, m := range warehouse {
	//	for _, s := range m {
	//		fmt.Printf("%s", s)
	//	}
	//	fmt.Printf("\n")
	//}

	return sumGPSCoords
}

type ops struct {
	dotX, dotY   int // set these to "."
	dot2X, dot2Y int
	obj1         string
	obj1X, obj1Y int
	obj2         string
	obj2X, obj2Y int
}

func moveObject2(warehouse [][]string, obj string, curX, curY int, move string, o []ops) (bool, []ops) {
	if move == "^" {
		x, y := curX-1, curY
		// up
		if warehouse[x][y] == "#" {
			// cant move
			return false, o
		} else if warehouse[x][y] == "." {
			// can move
			return true, o
		} else if warehouse[x][y] == "@" {
			if obj == "O" {
				// move possible since @ will move
				return true, o
			} else {
				return false, o
			}
		} else {
			if warehouse[x][y] == "]" {
				canPushR, newO := moveObject2(warehouse, warehouse[x][y], x, y, move, o)      // ]
				canPushL, newO2 := moveObject2(warehouse, warehouse[x][y-1], x, y-1, move, o) // [
				if canPushR && canPushL {
					//warehouse[x-1][y] = warehouse[x][y]
					//warehouse[x-1][y-1] = warehouse[x][y-1]
					//warehouse[x][y] = "."
					//warehouse[x][y-1] = "."
					obj1X, obj1Y := x-1, y
					obj1 := warehouse[x][y]
					obj2X, obj2Y := x-1, y-1
					obj2 := warehouse[x][y-1]
					dotX, dotY := x, y
					dot2X, dot2Y := x, y-1
					sw := ops{
						dotX:  dotX,
						dotY:  dotY,
						dot2X: dot2X,
						dot2Y: dot2Y,
						obj1:  obj1,
						obj1X: obj1X,
						obj1Y: obj1Y,
						obj2:  obj2,
						obj2X: obj2X,
						obj2Y: obj2Y,
					}
					retO := []ops{}
					for _, op := range o {
						retO = append(retO, op)
					}
					for _, op := range newO {
						retO = append(retO, op)
					}
					for _, op := range newO2 {
						retO = append(retO, op)
					}
					retO = append(retO, sw)
					return true, retO
				} else {
					return false, o
				}
			} else if warehouse[x][y] == "[" {
				canPushR, newO := moveObject2(warehouse, warehouse[x][y], x, y, move, o)      // [
				canPushL, newO2 := moveObject2(warehouse, warehouse[x][y+1], x, y+1, move, o) // ]
				if canPushR && canPushL {
					//warehouse[x-1][y] = warehouse[x][y]
					//warehouse[x-1][y+1] = warehouse[x][y+1]
					//warehouse[x][y] = "."
					//warehouse[x][y+1] = "."
					obj1X, obj1Y := x-1, y
					obj1 := warehouse[x][y]
					obj2X, obj2Y := x-1, y+1
					obj2 := warehouse[x][y+1]
					dotX, dotY := x, y
					dot2X, dot2Y := x, y+1
					sw := ops{
						dotX:  dotX,
						dotY:  dotY,
						dot2X: dot2X,
						dot2Y: dot2Y,
						obj1:  obj1,
						obj1X: obj1X,
						obj1Y: obj1Y,
						obj2:  obj2,
						obj2X: obj2X,
						obj2Y: obj2Y,
					}
					retO := []ops{}
					for _, op := range o {
						retO = append(retO, op)
					}
					for _, op := range newO {
						retO = append(retO, op)
					}
					for _, op := range newO2 {
						retO = append(retO, op)
					}
					retO = append(retO, sw)
					return true, retO
				} else {
					return false, o
				}
			}
		}
	} else if move == "v" {
		// down
		x, y := curX+1, curY
		// up
		if warehouse[x][y] == "#" {
			// cant move
			return false, o
		} else if warehouse[x][y] == "." {
			// can move
			return true, o
		} else if warehouse[x][y] == "@" {
			if obj == "]" || obj == "[" {
				// move possible since @ will move
				return true, o
			} else {
				return false, o
			}
		} else {
			if warehouse[x][y] == "]" {
				canPushR, newO := moveObject2(warehouse, warehouse[x][y], x, y, move, o)      // ]
				canPushL, newO2 := moveObject2(warehouse, warehouse[x][y-1], x, y-1, move, o) // [
				if canPushR && canPushL {
					//warehouse[x+1][y] = warehouse[x][y]
					//warehouse[x+1][y-1] = warehouse[x][y-1]
					//warehouse[x][y] = "."
					//warehouse[x][y-1] = "."
					obj1X, obj1Y := x+1, y
					obj1 := warehouse[x][y]
					obj2X, obj2Y := x+1, y-1
					obj2 := warehouse[x][y-1]
					dotX, dotY := x, y
					dot2X, dot2Y := x, y-1
					sw := ops{
						dotX:  dotX,
						dotY:  dotY,
						dot2X: dot2X,
						dot2Y: dot2Y,
						obj1:  obj1,
						obj1X: obj1X,
						obj1Y: obj1Y,
						obj2:  obj2,
						obj2X: obj2X,
						obj2Y: obj2Y,
					}
					retO := []ops{}
					for _, op := range o {
						retO = append(retO, op)
					}
					for _, op := range newO {
						retO = append(retO, op)
					}
					for _, op := range newO2 {
						retO = append(retO, op)
					}
					retO = append(retO, sw)
					//o = append(o, sw)
					return true, retO
				} else {
					return false, o
				}
			} else if warehouse[x][y] == "[" {
				canPushR, newO := moveObject2(warehouse, warehouse[x][y], x, y, move, o)      // [
				canPushL, newO2 := moveObject2(warehouse, warehouse[x][y+1], x, y+1, move, o) // ]
				if canPushR && canPushL {
					//warehouse[x+1][y] = warehouse[x][y]
					//warehouse[x+1][y+1] = warehouse[x][y+1]
					//warehouse[x][y] = "."
					//warehouse[x][y+1] = "."
					obj1X, obj1Y := x+1, y
					obj1 := warehouse[x][y]
					obj2X, obj2Y := x+1, y+1
					obj2 := warehouse[x][y+1]
					dotX, dotY := x, y
					dot2X, dot2Y := x, y+1
					sw := ops{
						dotX:  dotX,
						dotY:  dotY,
						dot2X: dot2X,
						dot2Y: dot2Y,
						obj1:  obj1,
						obj1X: obj1X,
						obj1Y: obj1Y,
						obj2:  obj2,
						obj2X: obj2X,
						obj2Y: obj2Y,
					}
					retO := []ops{}
					for _, op := range o {
						retO = append(retO, op)
					}
					for _, op := range newO {
						retO = append(retO, op)
					}
					for _, op := range newO2 {
						retO = append(retO, op)
					}
					retO = append(retO, sw)
					//o = append(o, sw)
					return true, retO
				} else {
					return false, o
				}
			}
		}
	} else if move == ">" {
		// right
		x, y := curX, curY+1
		// up
		if warehouse[x][y] == "#" {
			// cant move
			return false, o
		} else if warehouse[x][y] == "." {
			// can move
			return true, o
		} else if warehouse[x][y] == "@" {
			if obj == "O" {
				// move possible since @ will move
				return true, o
			} else {
				return false, o
			}
		} else {
			// move box []
			canPush, o := moveObject2(warehouse, warehouse[x][y], x, y, move, o)
			if canPush {
				warehouse[x][y+1] = warehouse[x][y]
				warehouse[x][y] = "."
				//obj1X, obj1Y := x, y+1
				//obj1 := warehouse[x][y]
				//dotX, dotY := x, y
				//sw := ops{
				//	dotX:  dotX,
				//	dotY:  dotY,
				//	obj1:  obj1,
				//	obj1X: obj1X,
				//	obj1Y: obj1Y,
				//	obj2:  "-",
				//}
				//o = append(o, sw)
				return true, o
			} else {
				return false, o
			}
		}
	} else if move == "<" {
		// left
		x, y := curX, curY-1
		// up
		if warehouse[x][y] == "#" {
			// cant move
			return false, o
		} else if warehouse[x][y] == "." {
			// can move
			return true, o
		} else if warehouse[x][y] == "@" {
			if obj == "O" {
				// move possible since @ will move
				return true, o
			} else {
				return false, o
			}
		} else {
			// move box []
			canPush, o := moveObject2(warehouse, warehouse[x][y], x, y, move, o)
			if canPush {
				warehouse[x][y-1] = warehouse[x][y]
				warehouse[x][y] = "."
				//obj1X, obj1Y := x, y-1
				//obj1 := warehouse[x][y]
				//dotX, dotY := x, y
				//sw := ops{
				//	dotX:  dotX,
				//	dotY:  dotY,
				//	obj1:  obj1,
				//	obj1X: obj1X,
				//	obj1Y: obj1Y,
				//	obj2:  "-",
				//}
				//o = append(o, sw)
				return true, o
			} else {
				return false, o
			}
		}
	} else {
		fmt.Println("Invalid move: " + move)
		return false, o
	}
	return false, o
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func part2(warehouse [][]string, moves []string, r robot) int {
	sumGPSCoords := 0
	for i, m := range moves {
		// push operations should only be done if everything can be moved
		_ = i
		o := []ops{}
		canMove, retO := moveObject2(warehouse, warehouse[r.posX][r.posY], r.posX, r.posY, m, o)
		if canMove {
			// bug fix ops
			dup := map[string]interface{}{}
			for _, sw := range retO {
				warehouse[sw.obj1X][sw.obj1Y] = sw.obj1
				dup[getKey(sw.obj1X, sw.obj1Y)] = struct{}{}
				if sw.obj2 != "-" {
					warehouse[sw.obj2X][sw.obj2Y] = sw.obj2
					dup[getKey(sw.obj2X, sw.obj2Y)] = struct{}{}
				}
				if _, f := dup[getKey(sw.dotX, sw.dotY)]; !f {
					warehouse[sw.dotX][sw.dotY] = "."
				}
				if sw.obj2 != "-" {
					if _, f := dup[getKey(sw.dot2X, sw.dot2Y)]; !f {
						warehouse[sw.dot2X][sw.dot2Y] = "."
					}
					//warehouse[sw.dot2X][sw.dot2Y] = "."
				}
			}
			warehouse[r.posX][r.posY] = "."
			if m == "^" {
				r.posX = r.posX - 1
			} else if m == "v" {
				r.posX = r.posX + 1
			} else if m == ">" {
				r.posY = r.posY + 1
			} else if m == "<" {
				r.posY = r.posY - 1
			}
			warehouse[r.posX][r.posY] = "@"

		}

		// Code to print in place
		//	fmt.Printf("\033[0;0H")
		//	fmt.Printf("Map %d, moved: %v, move: %s \n", i, canMove, m)
		//	for _, m := range warehouse {
		//		for _, s := range m {
		//			fmt.Printf("%s", s)
		//		}
		//		fmt.Printf("\n")
		//	}
	}

	for i, s := range warehouse {
		for j, s2 := range s {
			if s2 == "[" {
				sumGPSCoords += 100*i + j
			}
		}
	}

	fmt.Println("Map")
	for _, m := range warehouse {
		for _, s := range m {
			fmt.Printf("%s", s)
		}
		fmt.Printf("\n")
	}

	return sumGPSCoords
}
