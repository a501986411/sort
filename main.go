package main

import "fmt"

func main() {
	var waitSortLen int
	fmt.Println("请输入待排序数组的长度:")
	fmt.Scanf("%d", &waitSortLen)
	// fmt.Println("您输入的值是:%d", waitSortLen)
	waitSort := make([]int, waitSortLen)

	fmt.Println("请输入需要排序的数组:")
	for i := 0; i < waitSortLen; i++ {
		fmt.Scan(&waitSort[i])
	}
	fmt.Printf("输入的数组是:[")
	for i := 0; i < len(waitSort); i++ {
		fmt.Printf("%d,", waitSort[i])
	}
	fmt.Println("]")
	fmt.Print("冒泡排序结果:")
	afterSort := DubbleSort(waitSort[:])
	for i := 0; i < len(afterSort); i++ {
		fmt.Printf("%d ", afterSort[i])
	}
	fmt.Println("")

	fmt.Print("选择排序结果:")
	afterSelectSort := SelectSort(waitSort, false)
	for i := 0; i < len(afterSelectSort); i++ {
		fmt.Printf("%d ", afterSelectSort[i])
	}
	fmt.Println("")

	fmt.Print("插入排序结果:")
	insertSort := InsertSort(waitSort)
	PrintArrInt(insertSort)
}

func PrintArrInt(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println("")
}

//冒泡排序算法
func DubbleSort(waitSort []int) []int {
	for i := 0; i < len(waitSort)-1; i++ {
		for j := i + 1; j < len(waitSort); j++ {
			if waitSort[i] > waitSort[j] {
				waitSort[i], waitSort[j] = waitSort[j], waitSort[i]
			}
		}
	}
	return waitSort
}

//选择排序--升序
//isAsc 是否是升序,true升序，false：降序
func SelectSort(waitSort []int, isAsc bool) []int {
	length := len(waitSort)
	for i := 0; i < length-1; i++ {
		target := i
		for j := i + 1; j < length; j++ {
			if isAsc {
				if waitSort[target] > waitSort[j] {
					target = j
				}
			} else {
				if waitSort[target] < waitSort[j] {
					target = j
				}
			}
		}
		waitSort[i], waitSort[target] = waitSort[target], waitSort[i]
	}
	return waitSort
}

//插入排序
func InsertSort(waitSort []int) []int {
	for i := 1; i < len(waitSort); i++ {
		currentValue := waitSort[i]
		for j := i - 1; j >= 0; j-- {
			if currentValue < waitSort[j] {
				waitSort[j+1], waitSort[j] = waitSort[j], currentValue
			} else {
				waitSort[j+1] = currentValue
			}
		}
	}
	return waitSort
}
