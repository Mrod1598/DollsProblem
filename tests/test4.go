package tests

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Mrod1598/dolls/dolldelivery"
)

//Test4 a direct path
func Test4() {
	const start = "Kruthika's abode"
	const end = "Matt's pad"
	distance, path := dolldelivery.GoFindShortestPath(start, end, dolldelivery.Neighborhood)
	if len(path) > 0 {
		fmt.Println("Minimum distance: ", distance)
		fmt.Printf("%s ", path[0])
		for i := 1; i < len(path); i++ {
			fmt.Printf("===> %s ", path[i])
		}
	}
	fmt.Print("\nPress 'Enter' to close...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
