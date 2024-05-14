package tron

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/weibi/tron-protocol/common/base58"
	"math/big"
	"strconv"
	"strings"
)

// 处理合约获取余额
func processBalanceOfData(trc20 []byte) (amount *big.Int) {
	if len(trc20) >= 32 {
		amount = new(big.Int).SetBytes(trc20[0:32])
	}

	return
}

// 处理合约获取余额参数
func processBalanceOfParameter(addr string) (data []byte) {
	methodID, _ := HexDecode("70a08231")
	add, _ := base58.DecodeCheck(addr)
	paddedAddress := common.LeftPadBytes(add[1:], 32)
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	return
}

func trc20transferData(to string, amount *big.Int) (data []byte) {
	methodID, _ := HexDecode("a9059cbb")
	addr, _ := base58.DecodeCheck(to)
	paddedAddress := common.LeftPadBytes(addr[1:], 32)

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	return
}

// SignTransaction 签名交易
//func SignTransaction(transaction *core.Transaction, key *ecdsa.PrivateKey) (hash []byte, err error) {
//	transaction.GetRawData().Timestamp = time.Now().UnixNano() / 1000000
//	rawData, err := proto.Marshal(transaction.GetRawData())
//
//	if err != nil {
//		return
//	}
//
//	h256h := sha256.New()
//	h256h.Write(rawData)
//	hash = h256h.Sum(nil)
//
//	contractList := transaction.GetRawData().GetContract()
//
//	for range contractList {
//		signature, err := crypto.Sign(hash, key)
//		if err != nil {
//			return nil, err
//		}
//		transaction.Signature = append(transaction.Signature, signature)
//	}
//
//	return
//}

func String2int64(val string) (num int64, err error) {
	val1, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, errors.New("转账金额不对")
	}

	num = int64(val1 * 1000000)
	return
}

func BigToken2Usdt(token *big.Int) (usdt string) {
	s := token.Text(10)
	l := len(s)
	if l > 6 {
		return s[:l-6] + "." + s[l-6:]
	} else {
		padding := 6 - l
		rtn := "0."
		if padding > 0 {
			rtn = rtn + strings.Repeat("0", padding)

		}

		s = rtn + s
		//s = strings.TrimRight(s, "0")
		//s = strings.TrimSuffix(s, ".")
		return s
	}
}

func Balance2Trx(bln uint64) string {
	str := fmt.Sprintf("%06d", bln)

	strlen := len(str)
	if strlen > 6 {
		return str[0:strlen-6] + "." + str[strlen-6:]
	}

	str = "0." + str

	return str
}

// 1,073.091562 , 123 , 123.123
func Usdt2token(usdt string) (token int64, err error) {
	usdt = strings.ReplaceAll(usdt, ",", "")
	n, err := strconv.ParseFloat(usdt, 64)
	if err != nil {
		return
	}
	// 10^6 ,usdt 六位有效小数
	n = n * 1000000

	token = int64(n)
	return
}

func IsNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
