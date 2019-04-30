package main

import (
    "fmt"
    "github.com/Ankr-network/dccn-common/binance"
)

func main() {
    accounts,addresses,err:=binance.CreateAccount(100)
    if err!=nil {
        fmt.Errorf("CreateAccount error",err.Error())
    }
    fmt.Printf("accounts=%v,address=%v",accounts,addresses)
}
