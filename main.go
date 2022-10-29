package main

import (
	"fmt"
	"gopractice/test"
	"gopractice/local/pkctest"
	"gopractice/local/under_score"
)

func main(){
	// var str string;
	// fmt.Scan(&str);
	// if len(str) <= 5{
	// 	fmt.Println("入力された文字数は 5 文字以下です")
	// }else{
	// 	fmt.Println("入力された文字数は 6 文字以上です")
	// }
	fmt.Println("main関数から呼び出された");
	
	test.Test1()
	//test.test2() test.goのtest2関数は先頭が小文字なので参照できない
	pkctest.Test3()
	a_score.Under()
}