package main
import "fmt"

func main() {
    var arr [5]int
    arr[0] = 10
    arr[1] = 20
    arr[2] = 30
    arr[3] = 40
    arr[4] = 50
    fmt.Println("Array elements are:", arr)



    for i := 0; i < len(arr); i++ {
        fmt.Println("Element at index", i, "is", arr[i])
    }


	arr2 := [5]int{10, 20, 30, 40, 50}
    fmt.Println("Array elements are:", arr2)

    for i := 0; i < len(arr2); i++ {
        fmt.Println("Element at index", i, "is", arr2[i])	
	}



	//数组非常高效但是很死板。很多时候，我们在事前并不知道数组的长度是多少。针对这个情况，slices （切片） 出来了。







}