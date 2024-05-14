package tron

import (
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/protobuf/proto"
	"math/big"
	"time"
)

func BalanceOf(address string) (balance uint64, err error) {
	info, err := GetWalletAccount(address)
	if err != nil {
		return
	}

	balance = info.Balance
	return
}
func BalanceOfTrc20(address string) (balance *big.Int, err error) {
	balance, err = grpcClient.TRC20ContractBalance(address, tokenAddr)
	return
}

func HasEnoughTrc20(from, amount string) (ok bool, balance *big.Int, err error) {
	balance, err = BalanceOfTrc20(from)
	if err != nil {
		return
	}
	amountInt, _ := String2int64(amount)
	ok = balance.Int64() >= amountInt
	return
}

func SendTrx(mnemonic string, idx int, to, amount string) (blockNumber int64, txId string, err error) {
	valInt, err := String2int64(amount)
	if err != nil {
		return
	}
	from, err := AddressFromMnemonic(mnemonic, idx)
	if err != nil {
		return
	}
	tx, err := grpcClient.Transfer(from, to, valInt)

	if err != nil {
		return
	}
	sendingTx, err := signTransaction(mnemonic, idx, tx.Transaction)
	txId = fmt.Sprintf("%x", tx.Txid)
	blockNumber, err = broadcastTransaction(sendingTx, txId)
	if err != nil {
		fmt.Println("Broadcast failed:", err.Error())
		return
	}

	//fmt.Printf("recipt: %x\n", tx.Txid)
	//fmt.Printf("blockNumber: %d\n", blockNumber)
	return
}

func SendTrc20(mnemonic string, idx int, to, amount string) (txID string, err error) {
	feeLimit := int64(1000_000_000)
	amountInt, err := String2int64(amount)
	if err != nil {
		return
	}
	amountBig := big.NewInt(amountInt)

	from, err := AddressFromMnemonic(mnemonic, idx)
	tx, err := grpcClient.TRC20Send(from, to, tokenAddr, amountBig, feeLimit)
	sendingTx, err := signTransaction(mnemonic, idx, tx.Transaction)
	if err != nil {
		return
	}

	_, err = grpcClient.Broadcast(sendingTx)

	if err != nil {
		return
	}
	txID = fmt.Sprintf("%x", tx.Txid)
	return
}

func broadcastTransaction(tx *core.Transaction, txHash string) (blockid int64, err error) {
	_, err = grpcClient.Broadcast(tx)
	if err != nil {
		return
	}
	waitTime := 100 // 100秒
	for {
		time.Sleep(10 * time.Second)
		waitTime = waitTime - 10
		if waitTime < 0 {
			err = fmt.Errorf("timeout")
			return
		}
		txii, errII := grpcClient.GetTransactionInfoByID(txHash)
		if errII != nil {
			fmt.Println("GetTransactionInfoByID failed:", errII.Error())
			continue
		}

		if txii.Result != 0 {
			err = fmt.Errorf("failed:%s", txii.ResMessage)
			return
		}

		blockid = txii.BlockNumber
		return
	}
}

func signTransaction(mnemonic string, idx int, tx *core.Transaction) (*core.Transaction, error) {
	privateKey, err := PrivateFromMnemonic(mnemonic, idx)
	if err != nil {
		return nil, err
	}

	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)
	return tx, nil
}
