package main

func main() {
	first := 4
	second := 8
	var multiplayer func(x, y int) int
	multiplayer = func(x, y int) int { return x + y }
	println(multiplayer(first, second))
	subtract := func(x, y int) int { return x / y }
	println(subtract(second, first))
	println(calc(first, second, sumFunc))
	println(calc(second, first, substructFunc))
	funcDevide2 := funcMaker(2)
	funcDevide10 := funcMaker(10)
	println(funcDevide2(100))
	println(funcDevide10(100))
}

func calc(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

func sumFunc(x, y int) int {
	return x + y
}

func substructFunc(x, y int) int {
	return x - y
}

func funcMaker(devider int) func(x int) int {
	deviderFunc := func(x int) int {
		return x / devider
	}
	return deviderFunc
}
