package main

import (
	"fmt"
	"github.com/palletone/adaptor"
	"github.com/palletone/fabric-adaptor"
)

func main()  {
	fabAdapotr := fabricadaptor.NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	fabAdapotr.ChannelID = "mychannel"//sdk check

	fabAdapotr.OrgAdmin = "Admin"
	fabAdapotr.OrgName = "org1"

	args := [][]byte{[]byte("A"),[]byte("B"), []byte("100")}
	outputInvokeTx,err := fabAdapotr.CreateContractInvokeTx(
		&adaptor.CreateContractInvokeTxInput{
			ContractAddress:"chaincode_example02",
			Function:"invoke",Args:args})
	if err != nil {
		fmt.Println("CreateContractInvokeTx", err.Error())
		return
	}
	//fmt.Println("len", len(outputInvokeTx.RawTransaction))
	//fmt.Println(hex.EncodeToString(outputInvokeTx.RawTransaction))
	//fmt.Println(string(outputInvokeTx.Extra))
	//sign tx
	outputSign,err := fabAdapotr.SignTransaction(&adaptor.SignTransactionInput{
		PrivateKey:[]byte(fabAdapotr.UserName),
		Transaction:outputInvokeTx.RawTransaction,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//
	//fmt.Println("outputSign.SignedTx", hex.EncodeToString(outputSign.SignedTx))

	outputSend,err := fabAdapotr.SendTransaction(&adaptor.SendTransactionInput{
		Transaction:outputSign.SignedTx,
		Extra:[]byte("invoke"),//Must set
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("outputSend.TxID", string(outputSend.TxID))
}