package main

import "fmt"

// // Long way
// func main() {
// 	lines, err := datafile.GetStrings("votes.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var names []string
// 	var counts []int
// 	for _, line := range lines {
// 		matched := false
// 		for i, name := range names {
// 			if name == line {
// 				matched = true
// 				counts[i]++
// 			}
// 		}
// 		if matched == false {
// 			names = append(names, line)
// 			counts = append(counts, 1)
// 		}
// 	}

// 	fmt.Println(lines)
// 	for i, name := range names {
// 		fmt.Printf("%s: %d\n", name, counts[i])
// 	}
// }

// //Short way
// func main() {
// 	lines, err := datafile.GetStrings("votes.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	counts := make(map[string]int)
// 	for _, line := range lines {
// 		counts[line]++
// 	}
// 	fmt.Println(counts)
// 	for key, value := range counts {
// 		fmt.Printf("%s : %d \n", key, value)
// 	}
// }

// magnets on the fridge
func main() {
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d \n", medal, rank)
	}
	fmt.Println("")
	delete(ranks, "bronze")
	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d \n", medal, rank)
	}
}
