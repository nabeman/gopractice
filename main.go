package main

import "fmt"

func main(){
	var str string;
	fmt.Scan(&str);
	if len(str) <= 5{
		fmt.Println("入力された文字数は 5 文字以下です")
	}else{
		fmt.Println("入力された文字数は 6 文字以上です")
	}
}