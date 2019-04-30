package main

import (
    "fmt"
    "github.com/binance-chain/go-sdk/client"
    "github.com/binance-chain/go-sdk/keys"
    "github.com/binance-chain/go-sdk/types"
    "github.com/binance-chain/go-sdk/types/msg"
)

func main() {
    keyManager, err := keys.NewPrivateKeyManager("c0ff5f1df186408fff51689bff2f067b28de5073448975ff9cf2ab9d2e948c84")
    if err!=nil{
        fmt.Errorf("NewPrivateKeyManager error",err.Error())
    }
    toaddress := "bnb1np3tdknj2stkdlv9r8aj7pfj6nlza2pdlcltrh"
    addr, _ := types.AccAddressFromBech32(toaddress)
    cli, _ := client.NewDexClient("testnet-dex.binance.org", types.TestNetwork, keyManager)
    send, err := cli.SendToken([]msg.Transfer{{addr, []types.Coin{{"BNB", 10000}}}}, true)
    if err != nil {
        fmt.Errorf("SendToken error", err.Error())
    }
    if send != nil {
        fmt.Printf("send token successful")
    }else {
        fmt.Printf("send token failure")
    }
}
