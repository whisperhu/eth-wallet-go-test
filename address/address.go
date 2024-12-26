package address

import "github.com/ethereum/go-ethereum/crypto"

type EthereumAddress struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

func CreatAddressFromPrivateKey() (*EthereumAddress, error) {

	crypto.GenerateKey()
	return nil, nil
}
