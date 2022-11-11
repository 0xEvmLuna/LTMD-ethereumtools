package ethereumsupport

import (
	"context"
	"errors"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
)

var (
	MasterSendAddress common.Address
	MasterNonce       uint64
	MasterPrivatekey  string
)

func (p *Provide) SendFaucet(receiver string) ([]byte, error) {
	gasPrice, err := p.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	if os.Getenv("sendpoolA") == "" &&
		os.Getenv("sendpoolB") == "" &&
		os.Getenv("sendpoolC") == "" {
		return nil, errors.New("something sendPool is nil")
	}

	if os.Getenv("sendPrivatekeyA") == "" &&
		os.Getenv("sendPrivatekeyB") == "" &&
		os.Getenv("sendPrivatekeyC") == "" {
		return nil, errors.New("something sendPool is nil")
	}

	faucet := new(Faucet)
	faucet.SendPool = []common.Address{
		common.HexToAddress(os.Getenv("sendpoolA")),
		common.HexToAddress(os.Getenv("sendpoolB")),
		common.HexToAddress(os.Getenv("sendpoolC")),
	}

	faucet.SendPrivateKey = []string{
		os.Getenv("sendPrivatekeyA"),
		os.Getenv("sendPrivatekeyB"),
		os.Getenv("sendPrivatekeyC"),
	}

	var na, nb, nc uint64
	na, _ = p.Client.PendingNonceAt(context.Background(), common.HexToAddress(os.Getenv("sendpoolA")))
	nb, _ = p.Client.PendingNonceAt(context.Background(), common.HexToAddress(os.Getenv("sendpoolB")))
	nc, _ = p.Client.PendingNonceAt(context.Background(), common.HexToAddress(os.Getenv("sendpoolC")))

	faucet.Nonce = []uint64{na, nb, nc}

	faucet.GasPrice = gasPrice
	faucet.Receiver = common.HexToAddress(receiver)
	// 0.05 eth value.
	faucet.Amount = big.NewInt(50000000000000000)
	faucet.Gaslimit = uint64(21000)

	return sendWithFaucet(p, faucet)
}

func sendWithFaucet(p *Provide, faucet *Faucet) ([]byte, error) {
	for i := 0; i < len(faucet.SendPool); i++ {
		balance, _ := GetBalance(p, faucet.SendPool[i])
		if balance > 0.05 {
			MasterSendAddress = faucet.SendPool[i]
			MasterNonce = faucet.Nonce[i]
			MasterPrivatekey = faucet.SendPrivateKey[i]
		}
	}

	privateKey, err := crypto.HexToECDSA(MasterPrivatekey)
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	tx := types.NewTransaction(
		MasterNonce,
		faucet.Receiver,
		faucet.Amount,
		faucet.Gaslimit,
		faucet.GasPrice,
		data,
	)

	chainID, err := p.Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = p.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	faucet.Tx = signedTx.Hash().Hex()

	bytesFaucet, _ := Marshal(faucet)
	return bytesFaucet, nil
}
