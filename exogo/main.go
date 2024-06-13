package main

import "fmt"

func main (){
  thisCost := 3
  numness := 7421 
  fmt.Print(calculateFinalBill(float64(thisCost),numness))
}
func calculateFinalBill(costpermessage float64, nummessages int) float64 {
  baseBill := calculateBaseBill(costpermessage,nummessages)   
  amountDeducted := baseBill * calculatediscount(nummessages)
  fmt.Print(amountDeducted)
  return baseBill - amountDeducted 
}

func calculatediscount(messagessent int) float64 {
  if messagessent >1000 && messagessent<5000{
    return .1
  }
  if messagessent >5000{
    return .2
  }else{
    return 0.0
  }
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
  return costPerMessage * float64(messagesSent)
}

