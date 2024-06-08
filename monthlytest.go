package main

import "fmt"

func main (){
  subscription := "enterprise"
  fmt.Println(getMonthlyPrice(subscription))
}

func getMonthlyPrice (tier string) int {
  if tier == "basic"{
    return 10000
  } else if tier == "premium"{
    return 15000
  } else if tier=="enterprise"{
    return 20000
  }else {
    return 0
  }
}

