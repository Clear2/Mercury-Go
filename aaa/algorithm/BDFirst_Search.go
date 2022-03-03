package main

import "fmt"

// Flood fill

func main() {
	//image := [][]int{
	//	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	//}
	//var sr, sc, newColor = 1, 1, 2
	//fmt.Println(maxAreaOfIsland(image))
	recursion(5, 10)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	helper(&image, sr, sc, image[sr][sc], newColor)
	return image
}

func helper(image *[][]int, i, j, oldColor, newColor int) {
	if i < 0 || i >= len(*image) || j < 0 || j >= len((*image)[0]) || (*image)[i][j] != oldColor {
		return
	}
	(*image)[i][j] = newColor
	helper(image, i+1, j, oldColor, newColor)
	helper(image, i-1, j, oldColor, newColor)
	helper(image, i, j+1, oldColor, newColor)
	helper(image, i, j-1, oldColor, newColor)
}
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = MaxInt(maxArea, AreaOfIsland(grid, i, j))
			}
		}
	}
	return maxArea
}
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func AreaOfIsland(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
		grid[i][j] = 0
		return 1 + AreaOfIsland(grid, i+1, j) + AreaOfIsland(grid, i-1, j) + AreaOfIsland(grid, i, j-1) + AreaOfIsland(grid, i, j+1)

	}
	return 0
}

func recursion(i, j int) {
	fmt.Println("x----->>", i)
	if i <= 0 {
		return
	}
	recursion(i-1, j)
	fmt.Println("y----->>", i)
	fmt.Println("z----->>>", j)
}

//  	1
// 	 2      3
// 4  5      6  7
