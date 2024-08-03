package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	options := "options: abcdefghijklmnopqrstuvwxyz"
	if len(args) == 0 {
		fmt.Println(options)
		return
	}
	str := toStr(args)
	// fmt.Println(str)

	// check if the arguments are valid
	if !isValid(str){
		fmt.Println("Invalid Option")
		return
	}

	//check if all arguments if h is active
	if isHactive(str){
		fmt.Println(options)
		return 
	}

	chars := []rune{}
	seen := make(map[rune]bool)
	for _, char := range str {
		if char >= 'a' && char <= 'z' && !seen[char]{
			chars = append(chars,char)
			seen[char]=true
		}
	}
	bits := make([]int,32)
	for _, char := range chars {
		pos :=  int(char-'a')+1 
		bits[pos-1] = 1
	}
	result := ""
	count:=0
	for i := len(bits)-1; i>=0;i--{
		if count == 8 && i != 0{
			result += " "
			count=0
		}
		if bits[i] == 1 {
			result+="1"
		}else{
			result+="0"
		}
		count++
	}
	fmt.Println(result)
}

func isValid(s string) bool {
	for _, char := range s{
		if !(char >= 'a' && char <= 'z' || char == '-'){
			return false
		}
	}
	return true 
}

func toStr(arr []string) string {
	res := ""
	for _, s := range arr {
		res += s 
	}
	return res 
}

func isHactive(s string) bool {
	for i := 0; i < len(s)-1; i++{
		if rune(s[i])=='-' && rune(s[i+1])=='h'{
			return true
		}
	}
	return false
}

