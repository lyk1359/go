package main

import(
	"fmt"
	"math/rand"
	"time" 
)
//난수 추출된 수의 소수 판정 프로그램 v0.6

func main(){
	
	seed :=time.Now().Unix()
	rand.Seed(seed)

	isPrime :=true
	number := rand.Intn(150)+2 
	number=21
	fmt.Println("임의로 추출된 수: ",number)
	
	
	for i:=2; i < number; i++{ //1과 number일때 loop돌지 않음
		if number%i == 0{
			isPrime=false
			//break
		}
		fmt.Print(i, " ")
	}
	if isPrime{ //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다!")
	}else{
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}

