package main
 import ("fmt")
 func searchInsert(nums[]int, target int) int {
	  for i,num := range nums{
		if num== target{
		return i 

		}
	  }
	  return -1
 }
    func main() {
		   nums:= []int{1,2,5,4,6}
		   target:= 5
		   result:= searchInsert(nums,target)
		   fmt.Println(result)
	}