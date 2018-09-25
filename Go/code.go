 
package main

import (
	"fmt"
	"strconv"
)

//公用:2. 21.
type ListNode struct {
	Val int
	Next *ListNode
}

func main (){
	
	//1.两数之和 twoSum
	
	
	//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。 
	//你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。 
	//示例: 
	//给定 nums = [2, 7, 11, 15], target = 9 
	//因为 nums[0] + nums[1] = 2 + 7 = 9
	//所以返回 [0, 1]
	//nums := []int{2, 7, 11, 15, 2}
	//result := twoSum(nums, 4)
	//fmt.Println(result)
	
	
	//2. 两数相加 addTwoNumbers
	
	//给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，
	//它们的每个节点只存储单个数字。将两数相加返回一个新的链表。 
	//你可以假设除了数字 0 之外，这两个数字都不会以零开头。 
	//示例： 
	//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	//输出：7 -> 0 -> 8
	//原因：342 + 465 = 807
	//l1 := &ListNode{Val : 2, Next : &ListNode{Val : 4, Next : &ListNode{Val : 3, Next: nil}}}
	//l2 := &ListNode{Val : 5, Next : &ListNode{Val : 6, Next : &ListNode{Val : 4, Next: &ListNode{Val : 3, Next: nil}}}}
	////l1 := &ListNode{Val : 5, Next :  nil}
	////l2 := &ListNode{Val : 5, Next :  nil}
	//l3 := addTwoNumbers(l1, l2) 
	//for l3 != nil{
	//	fmt.Print(l3.Val)
	//	l3 = l3.Next
	//}


	//7.反转整数

	//给定一个 32 位有符号整数，将整数中的数字进行反转。
	//示例 1:
	//输入: 123
	//输出: 321
	//示例 2:
	//输入: -123
	//输出: -321
	//示例 3:
	//输入: 120
	//输出: 21	
	//注意:
	//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
	
    //fmt.Println(reverse(1534236469))
		
	
	
	//9.回文数 isPalindrome
	
	//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	//示例 1:
	//输入: 121
	//输出: true
	//示例 2:
	//输入: -121
	//输出: false
	//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
	//示例 3:
	//输入: 10
	//输出: false
	//解释: 从右向左读, 为 01 。因此它不是一个回文数。
	//进阶:
	//你能不将整数转为字符串来解决这个问题吗？
	//var c int
	//fmt.Println(isPalindrome(c))


	//13. 罗马数字转整数  Roman to Integer
	//罗马数字包含以下七种字符：I， V， X， L，C，D 和 M。 
	//字符          数值
	//I             1
	//V             5
	//X             10
	//L             50
	//C             100
	//D             500
	//M             1000 
	//例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。 
	//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
	//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。 
	//给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。 
	//示例 1: 
	//输入: "III"
	//输出: 3 
	//示例 2: 
	//输入: "IV"
	//输出: 4 
	//示例 3: 
	//输入: "IX"
	//输出: 9 
	//示例 4: 
	//输入: "LVIII"
	//输出: 58
	//解释: C = 100, L = 50, XXX = 30, III = 3. 
	//示例 5: 
	//输入: "MCMXCIV"
	//输出: 1994
	//解释: M = 1000, CM = 900, XC = 90, IV = 4.


	//14. 最长公共前缀  Longest Common Prefix
	//编写一个函数来查找字符串数组中的最长公共前缀。
	//
	//如果不存在公共前缀，返回空字符串 ""。
	//
	//示例 1:
	//
	//输入: ["flower","flow","flight"]
	//输出: "fl"
	//
	//示例 2:
	//
	//输入: ["dog","racecar","car"]
	//输出: ""
	//解释: 输入不存在公共前缀。
	//
	//说明:
	//
	//所有输入只包含小写字母 a-z 。 
	//fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	//fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	//fmt.Println(longestCommonPrefix([]string{"dog"}))
	//fmt.Println(longestCommonPrefix([]string{}))
	//fmt.Println("end")



	//20. 有效的括号 isValid
	
	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。 
	//有效字符串需满足： 
	//左括号必须用相同类型的右括号闭合。
	//左括号必须以正确的顺序闭合。 
	//注意空字符串可被认为是有效字符串。
	//示例 1: 
	//输入: "()"
	//输出: true 
	//示例 2:
	//输入: "()[]{}"
	//输出: true
	//示例 3:
	//输入: "(]"
	//输出: false
	//示例 4:
	//输入: "([)]"
	//输出: false
	//示例 5:
	//输入: "{[]}"
	//输出: true 
	//s := ""
	//fmt.Println(isValid(s))
	//fmt.Println(isValidOther(s))
	
	
	//21. 合同两个有序链表 mergeTwoLists
	//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
	//示例： 
	//输入：1->2->4, 1->3->4
	//输出：1->1->2->3->4->4
	//l1 := &ListNode{Val : 1, Next : &ListNode{Val : 2, Next : &ListNode{Val : 4, Next: nil}}}
	//l2 := &ListNode{Val : 1, Next : &ListNode{Val : 3, Next : &ListNode{Val : 4, Next: &ListNode{Val : 5, Next: nil}}}}
	//var l1 *ListNode
	//var l2 *ListNode
	//l3 := mergeTwoLists(l1, l2)
	//fmt.Println(l3)
	//for l3 != nil{
	//	fmt.Print(l3.Val)
	//	l3 = l3.Next
	//}

	
	//26. 删除排序数组中的重复项
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	
	
	//88.合并两个有序数组 Merge Sorted Array
	//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
	//说明:
	//初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
	//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
	//示例:
	//输入:
	//nums1 = [1,2,3,0,0,0], m = 3
	//nums2 = [2,5,6],       n = 3
	//输出: [1,2,2,3,5,6]
	//nums1 := []int{1,2,3,0,0,0}
	//m := 3
	//nums2 := []int{2,5,6}
	//n := 3
	//merge(nums1, m, nums2, n)
	//fmt.Println(nums1)
	
	
}

//1.
func twoSum(nums []int, target int) []int {
	var result []int
	maps := make(map[int]int)
	
	for k, v := range nums{
		if _, ok := maps[target - v]; ok {
			result = append(result, maps[target - v])
			result = append(result, k)
			return result
		}
		maps[v] = k
	} 
	return result
	
}


//2.
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/ 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	tempP := l3
	tempN := 0 
	temp := 0
	for { 
		if l1 != nil{
			temp += l1.Val
			l1 = l1.Next
		} 
		if l2 != nil{
			temp += l2.Val
			l2 = l2.Next
		}
		tempP.Val = (temp + tempN) % 10
		tempN = (temp + tempN) / 10
		temp = 0
		if l1 == nil && l2 == nil && tempN == 0{
			break
		}
		tempP.Next = &ListNode{}
		tempP = tempP.Next
	}
	return l3
}

//7  2147483647
func reverse(x int) int { 
	re := 0
	for x != 0{
		t := x % 10
		re = re * 10 + t
		x = x / 10
	}
	if re < -2147483648 || re > 2147483647 {
		return 0
	}
	return re
}
func reverseOther(x int) int {
	s :=""
	if x<0 {
		x = -x
		s+="-"
	}


	xStr := strconv.Itoa(x)
	runs :=[]rune(xStr)
	xLen :=len(runs)
	for i:=xLen-1; i>=0;i--  {
		s +=string(runs[i])
	}

	tt,_ :=strconv.Atoi(s)
	if tt> (2<<30)-1 || tt <(-2<<30) {
		tt= 0
	}
	return tt

}


//9.
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for len(s) > 1{
		if s[:1] == s[len(s)-1:]{
			s = s[1:]
			s = s[:len(s)-1]
		}else{
			return false
		}
	}
	return true
}


//13. 罗马数字转整数
func romanToInt(s string) int {
	maps := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		
	}
	return len(maps)
}


//14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	re := ""
	index := 0
	for len(strs) > 0{
		flag := false
		if len(strs[0]) >= index{
			re = strs[0][0:index]
		}else {
			break
		}
		for _, v := range strs{
			if len(v) >= index{
				if v[0:index] != re{
					flag = true
					break
				}
			}else {
				flag = true
				break
			}
		}
		if flag{
			break
		}
		index += 1
	}
	if len(strs) < 1{
		return ""
	}
	return re[:index-1]
}


//20.
func isValid(s string) bool {
	str1 := ""
	if s == ""{
		return true
	}
	if len(s) < 2 || s[:1] == "}" || s[:1] == ")" || s[:1] == "]" {
		return false
	}
	for s !=""{
		if s[:1] == "{" || s[:1] == "(" || s[:1] == "["{
			str1 += s[:1]
		}else{
			switch s[:1] {
			case "}":
				if str1 != "" && str1[len(str1)-1:] == "{" {
					str1 = str1[:len(str1)-1]
				}else {
					return false
				}

			case ")":
				if str1 != "" && str1[len(str1)-1:] == "("{
					str1 = str1[:len(str1)-1]
				}else {
					return false
				}
			case "]":
				if str1 != "" && str1[len(str1)-1:] == "["{
					str1 = str1[:len(str1)-1]
				}else {
					return false
				}
			}
		}
		s = s[1:]
		if s =="" && str1 == ""{
			return true
		}
		if s =="" && str1 != ""{
			break
		}
	}
	return false 
}
func isValidOther(s string) bool{
	brackets := map[byte]byte{'(':')','[':']','{':'}'}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if _, ok := brackets[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if 0 == len(stack) {
				return false
			}
			ch := stack[len(stack) - 1]
			if s[i] == brackets[ch] {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}

	if 0 < len(stack) {
		return false
	}
	return true
}


//21.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	tempP := l3
	for l1 != nil || l2 != nil{
		if l1 != nil && l2 != nil{
			if l1.Val <= l2.Val {
				tempP.Next = &ListNode{Val : l1.Val}
				l1 =l1.Next
			}else{
				tempP.Next = &ListNode{Val : l2.Val}
				l2 =l2.Next
			}
		}else{
			if l1 == nil && l2 != nil {
				tempP.Next = &ListNode{Val : l2.Val}
				l2 = l2.Next
			}else{
				tempP.Next = &ListNode{Val : l1.Val}
				l1 =l1.Next
			}
		}
		//if l1 == nil && l2 == nil{
		//	break
		//}
		tempP = tempP.Next
	}
	return l3.Next
}


//26
func removeDuplicates(nums []int) int {
	index := len(nums)
	for i:=1; i<index; i++ {
		if nums[i-1] == nums[i] {
			index --
			for j := i; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
		}
	}
	return index
}

//88.
func merge(nums1 []int, m int, nums2 []int, n int)  {
	var nums3 []int
	i := 0
	j := 0
	k := 0
	if len(nums2) < 1 {
		return 
	}
	for m > -0 || n > 0 {
		if m > -0 && n > 0{
			if nums1[j] > nums2[k]{
				nums3 = append(nums3, nums2[k])
				k++
				n--
			}else{
				nums3 = append(nums3, nums1[j])
				j++
				m--
			}
		}else{
			if m > 0{
				nums3 = append(nums3, nums1[j])
				j++
				m--
			}else{
				nums3 = append(nums3, nums2[k])
				k++
				n--
			}
		}
		i++
	} 
	for k, v := range nums3{
		nums1[k] = v
	}
}
func mergeOther(nums1 []int, m int, nums2 []int, n int)  {
	for m > 0 && n > 0 {
		a, b := nums1[m-1], nums2[n-1]
		if a > b {
			nums1[m+n-1] = a
			m--
		} else {
			nums1[m+n-1] = b
			n--
		}
	}
	if n > 0 {
		copy(nums1[:n], nums2[:n])
	}
}
