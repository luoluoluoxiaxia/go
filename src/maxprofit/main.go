package main

func main() {
	prices := []int{1,2,3,4,5}
	pricess := 0
	for i:=1;i<len(prices) ; i++ {
		if(prices[i]>prices[i-1]){
			pricess += (prices[i]-prices[i-1])
		}
	}
	print(pricess)
}
