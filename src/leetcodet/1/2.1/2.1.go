package main

import (
				"fmt"
				"os"
)
func removeDuplicate(a []int) []int{
	n:=len(a)
	if n <= 1{
		return a
	}
	var i int = 0
	for j:=1;j<n;j++ {
		if a[j]!=a[i]{
			a[i+1]=a[j]
			i++
		}
	}
	return a[:i+1]
}

func main(){
	f,_ := os.Open("test.in")
	defer f.Close()
	//f:=os.Stdin
	var n int
	fmt.Fscan(f,&n)
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Fscan(f,&a[i])
	}
	a=removeDuplicate(a);
	fmt.Print(len(a),":",a)
}