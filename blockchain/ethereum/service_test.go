package ethereum

import (
	"math/big"
	"testing"
)

func TestStart(t *testing.T) {

	fromKey := "{\"version\":3,\"id\":\"a8bf6362-d44f-4571-aa04-355a919ee38b\",\"address\":\"ef9a561256d282898e6b5297c7d3ff1e0e5493ab\",\"crypto\":{\"ciphertext\":\"e850617ce23d3955cb2db819b8ad18665429782b5716de7da5e088c33041337e\",\"cipherparams\":{\"iv\":\"2129210878639c3a3c318101a01e7932\"},\"cipher\":\"aes-128-ctr\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"salt\":\"d4476d0440a5ba446c7bb2d0eea769d608f88fa3a5d5c9a0c84b1e24acf45151\",\"n\":262144,\"r\":8,\"p\":1},\"mac\":\"163e82e8811712b49a846a16913c05eacfe8278eb39cf48767b8a071cebfc56b\"}}"
		
	fromPassword := "@nkr1234"//Your from account password
	ethS, err := NewEthService()
	if ethS == nil || err != nil {
		t.Error("Can't create eth service")
	}
	bf1,  _:= new(big.Float).SetString("1e-18")
	bf2, _ := new(big.Float).SetString("999999999999999")
	newF :=new(big.Float).Mul(bf2, bf1)

	err = ethS.TokenTransfer("ANKR", fromKey, fromPassword, "0xbbb092e9d4ddaf4e6a793c83eb4fa1a6bcd06389", newF)
	if err != nil {
		t.Errorf("ethS.TokenTransfer err: %s", err.Error())
	}

}
