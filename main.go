package main

import "fmt"

func main() {
	// var waitSortLen int
	// fmt.Println("请输入待排序数组的长度:")
	// fmt.Scanf("%d", &waitSortLen)
	// // fmt.Println("您输入的值是:%d", waitSortLen)
	// waitSort := make([]int, waitSortLen)

	// fmt.Println("请输入需要排序的数组:")
	// for i := 0; i < waitSortLen; i++ {
	// 	fmt.Scan(&waitSort[i])
	// }
	waitSort := []int{5, 2, 3, 1, 4}
	fmt.Printf("输入的数组是:[")
	for i := 0; i < len(waitSort); i++ {
		fmt.Printf("%d,", waitSort[i])
	}
	// fmt.Println("]")
	// fmt.Print("冒泡排序结果:")
	// afterSort := DubbleSort(waitSort[:])
	// for i := 0; i < len(afterSort); i++ {
	// 	fmt.Printf("%d ", afterSort[i])
	// }
	// fmt.Println("")

	// fmt.Print("选择排序结果:")
	// afterSelectSort := SelectSort(waitSort, false)
	// for i := 0; i < len(afterSelectSort); i++ {
	// 	fmt.Printf("%d ", afterSelectSort[i])
	// }
	// fmt.Println("")

	// fmt.Print("插入排序结果:")
	// insertSort := InsertSort(waitSort)
	// PrintArrInt(insertSort)

	// fmt.Print("希尔排序结果:")
	// shellSort := ShellSort(waitSort)
	// PrintArrInt(shellSort)

	// fmt.Print("归并排序结果:")
	// mergeSort := MergeSort(waitSort)
	// PrintArrInt(mergeSort)

	fmt.Print("快速排序结果:")
	fastSort := FastSort(waitSort)
	PrintArrInt(fastSort)
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

//希尔排序
func ShellSort(waitSort []int) []int {
	length := len(waitSort)
	gap := 5
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			fmt.Print("a ")
			temp := waitSort[i]
			j := i - gap
			for j >= 0 && waitSort[j] > temp {
				fmt.Print("b ")
				waitSort[j+gap] = waitSort[j]
				j -= gap
			}
			waitSort[j+gap] = temp
		}
		gap = gap / 3
	}

	return waitSort
}

//归并排序
func MergeSort(waitSort []int) []int {
	length := len(waitSort)
	if length < 2 {
		return waitSort
	}
	middleIndex := length / 2
	left := waitSort[0:middleIndex]
	right := waitSort[middleIndex:]
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

//快速排序
// 算法步骤
// 从数列中挑出一个元素，称为 "基准"（pivot）;

// 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
//在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；

// 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
func FastSort(waitSort []int) []int {
	return _fastSort(waitSort, 0, len(waitSort)-1)
}

//  5,2,3,1,4      | 5,2,3,1,4  |3 2 1 8 5 6 9 4 7
//  0 1 2 3 4      | 0 1 2 3 4  |0 1 2 3 4 5 6 7 8

func _fastSort(arr []int, left, right int) []int { // left=0,right=4

	if left < right {
		partitionIndex := partition(arr, left, right)
		_fastSort(arr, left, partitionIndex)
		_fastSort(arr, partitionIndex+1, right)

	}
	return arr
}

func partition(arr []int, leftInex, rightIndex int) int {
	pivot := leftInex
	index := pivot + 1
	for i := index; i <= rightIndex; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			PrintArrInt(arr)
			index += 1
		}
	}

	swap(arr, pivot, index-1)
	return index - 1
}

//较换2个数组元素
func swap(arr []int, i, j int) {
	fmt.Println("变化前:")
	PrintArrInt(arr)
	arr[i], arr[j] = arr[j], arr[i]
	fmt.Println("变化后:")
	PrintArrInt(arr)
}
