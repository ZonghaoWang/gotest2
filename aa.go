package main
import "fmt"
func firstMissingPositive(nums []int) int {
    var index int
    l := len(nums)
    if l == 0 {
        return 1
    }
    if l == 1 {
        if nums[0] == 1 {
            return 2
        } else {
            return 1
        }
    }
    for index < l {
        var current int
        current = nums[index]
        for current < l && current > 0 {
            currentValue := nums[current]
            nums[current] = 0
            if current < index {
                break
            }
            current = currentValue
        }
        index++
    }
    for index, value := range nums[1:] {
        if value != 0 {
            return index + 1
        }
    }
    return l
}
func main() {
    fmt.Println(firstMissingPositive([]int{3,4,-1,1}))
}