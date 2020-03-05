/*
   This file is part of go-palletone.
   go-palletone is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-palletone is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * @author PalletOne core developers <dev@pallet.one>
 * @date 2018-2020
 */
package fabricadaptor

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/palletone/adaptor"
)

func TestHashMessage(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	//fabAdapotr.UserName = "User1"
	output,err := fabAdapotr.HashMessage(&adaptor.HashMessageInput{Message:[]byte("123456")})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
	fmt.Println("output.Hash", hex.EncodeToString(output.Hash))
}

func TestSignMessage(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	output,err := fabAdapotr.SignMessage(&adaptor.SignMessageInput{Message:[]byte("123456")})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//30440220742c2cc005742c2d378634e1662f6604690879c4c45c107ed8e9f2c8fb594334022042ef319f4b0f6e2931e8f74de08c5edad2392c32c8269c4b7f3c42b5919507af
	fmt.Println("output.Hash", hex.EncodeToString(output.Signature))
}

func TestVerifySignature(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	signatureHex := "30440220742c2cc005742c2d378634e1662f6604690879c4c45c107ed8e9f2c8fb594334022042ef319f4b0f6e2931e8f74de08c5edad2392c32c8269c4b7f3c42b5919507af"
	signature,_ := hex.DecodeString(signatureHex)
	output,err := fabAdapotr.VerifySignature(&adaptor.VerifySignatureInput{Signature:signature,Message:[]byte("123456")})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//30440220066d5ecd29578ec40d9e4351331e6556a67c88317ee04d6f0b8e8c0ab2650b0f02207f94e337bc8665adfb5077d2a3b4caff384edc00faa32d587189085c4095c1da
	fmt.Println("output.Pass", output.Pass)
}

func TestGetBlockInfo(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	fabAdapotr.ChannelID = "mychannel"//sdk check
	//output,err := fabAdapotr.GetBlockInfo(&adaptor.GetBlockInfoInput{Latest:true})
		output,err := fabAdapotr.GetBlockInfo(&adaptor.GetBlockInfoInput{Height:31})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//30440220066d5ecd29578ec40d9e4351331e6556a67c88317ee04d6f0b8e8c0ab2650b0f02207f94e337bc8665adfb5077d2a3b4caff384edc00faa32d587189085c4095c1da
	fmt.Println("output.Block.BlockHeight", output.Block.BlockHeight)
	fmt.Println("output.Block.BlockID", hex.EncodeToString(output.Block.BlockID))//todo
}

func TestCreateContractInstallTx(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	fabAdapotr.ChannelID = "mychannel"//sdk check
	fabAdapotr.OrgAdmin = "Admin"
	fabAdapotr.OrgName = "org1"
	fabAdapotr.OrgID = "Org1MSP" //不能为空,否则合约实例化后不能调用
	output,err := fabAdapotr.CreateContractInstallTx(&adaptor.CreateContractInstallTxInput{
		Contract:[]byte("fabcar_xl10"),Extra:[]byte("github.com/fabcar_xl")})
	if err != nil {
		fmt.Println("fabAdapotr.CreateContractInstallTx", err.Error())
		return
	}
	_=output
	//fmt.Println(len(output.RawTransaction))
}

func TestCreateContractInitialTx(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	fabAdapotr.ChannelID = "mychannel"//sdk check
	fabAdapotr.OrgAdmin = "Admin"
	fabAdapotr.OrgName = "org1"
	fabAdapotr.OrgID = "Org1MSP" //不能为空,否则合约实例化后不能调用
	output,err := fabAdapotr.CreateContractInitialTx(&adaptor.CreateContractInitialTxInput{
		Contract:[]byte("fabcar_xl10"),Extra:[]byte("github.com/fabcar_xl")})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_=output
	//fmt.Println(string(output.RawTransaction))
}

func TestCreateContractInvokeTx(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	fabAdapotr.ChannelID = "mychannel"//sdk check
	args := [][]byte{[]byte("CAR11"),[]byte("zxl"), []byte("xiangli"), []byte("red"), []byte("zhang")}
	output,err := fabAdapotr.CreateContractInvokeTx(&adaptor.CreateContractInvokeTxInput{ContractAddress:"fabcar",Function:"createCar",Args:args})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_=output
	//fmt.Println(string(output.RawTransaction))
}

func TestGetContractInvokeTx(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	fabAdapotr.ChannelID = "mychannel"//sdk check
	txIDHex := "d606d16a62221a3ac1279208219a50e3fa36610bcaf99ca1ab6f934e4c72a261"
	//txIDHex := "701dbf13b4859d901b771d8e47862fb5a51d97fff2963d804c26587446d79258"
	txID, _ := hex.DecodeString(txIDHex)
	output,err := fabAdapotr.GetContractInvokeTx(&adaptor.GetContractInvokeTxInput{TxID:txID})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(output.TxRawData))
}

func TestQueryContract(t *testing.T) {
	fabAdapotr := NewAdaptorFabric("./connection.yaml")
	fabAdapotr.UserName = "User1"
	fabAdapotr.ChannelID = "mychannel"//sdk check
	output,err := fabAdapotr.QueryContract(&adaptor.QueryContractInput{ContractAddress:"fabcar",Function:"queryAllCars"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(output.QueryResult))
}
