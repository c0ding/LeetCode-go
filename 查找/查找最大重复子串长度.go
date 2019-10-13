package main

import (
	"fmt"
	"math"
)

//https://blog.csdn.net/heart66_A/article/details/83714168

func main() {

	leng := lengthOfLongestSubstring("abccwcdsf")
	fmt.Println(leng)
}


//利用滑动窗口和容器
func lengthOfLongestSubstring(s string)int{
	//借助一个容器，来判断，子串中是否有重复 妙
	m := make(map[byte]byte)
	sLen := len(s)
	start,end := 0,0
	var repeatCount int = 0
	//start 和 end 双条件判断，只有end一个也可以，可能这样更严谨一些吧
	for  start <sLen && end <sLen{
		temp := s[end]
		if _,ok := m[temp];!ok{
			//不存在说明该key是唯一的
			m[s[end]] = s[end]
			end++ //移动滑动窗口
			repeatCount = int(math.Max(float64(repeatCount),float64(end-start)))
		}else{
			//说明了有重复的，滑动窗口移动，则start+1,
			delete(m,s[start])
			start++
		}
	}
	return repeatCount

}