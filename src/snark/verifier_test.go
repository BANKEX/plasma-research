package main

import (
	"math/big"
	"testing"
)

/*

A = Pairing.G1Point(0x716af8745e1c56bd34ec221c41a7c8280da7289cdd5e3da097cd1ec0c37b824, 0x1b8af56c8a8e336e332039f51f69198f96a840a9f759e47928ee2005da0c5e75);
A_p = Pairing.G1Point(0xb817e60a38d9db1e33c328c52298e6585ebd3b2814bf79c646c12f617e8bb13, 0x268e77e26ca86cd4bef5c586dc2d7529bc2a3b52faf7df620d6f6df3e9b3978a);
B = Pairing.G2Point([0x3f9b971ca8e049d867175d1d099330f13f56eb14b1dcf5c7e65d72b0a61e946, 0x1e891f7806aac50bcb39627bbe4ae1b7674fa1247b8170542e58d37b20500457], [0x869f049ffd8648cf1cfe393119cfc1d2fab79c1d607ca03a35c769394864351, 0x9b6a9e259c151fe9327e68c3803b6c81322357f36a9fc62ced911525e966b26]);
B_p = Pairing.G1Point(0x2ab11c8e091ee2d5284465c22c3cc645c774d63f202f6dd60d7a1d9a930ace62, 0x2ff2ead04ccd73e86f643ffc4189a73b69179d90e6b2401d0aa493114f64c199);
C = Pairing.G1Point(0x2a5ef5a4b5a86c359f25957e2279d6d0eb6b97faf2342fc06e3bb93129254a8d, 0x259ea5957e70f5c35a2d3f4c9a4a84a74c35f10c0e0e2f0dea2b169b092c0fdb);
C_p = Pairing.G1Point(0x2bf222911543023f5932afb40b85818b1c316fa10b47aea8b591163fce350204, 0x1f2e95df3a6d50611931d90061c633122f8109f668541b3d100a5c6d457a069f);
H = Pairing.G1Point(0x268e91f0cf52a1162b60a3f923691a25c7d690b6915966d0266ea134c1f5e7f7, 0xb47e1bf8177da29fc0c7bb551a1d56bab5c9e942fe52fa749e9e6d5f8321e7a);
K = Pairing.G1Point(0x1cb3bd3cecf247723d17060e3ff538050609eaf57d240aee0caa289ce4735f84, 0x15f3cab12799c117a746d60a351980055fd1cf21257042d13966159169dbf966);



*/

func Hex2BI(s string) *big.Int {
	res, _ := new(big.Int).SetString(s[2:], 16)
	return res
}

func TestVerifier(t *testing.T) {
	// TODO: need to wrap this test
	// // privkey for seed 1 0x69b39aa2fb86c7172d77d4b87b459ed7643c1e4b052536561e08d7d25592b373
	// client, err := ethclient.Dial("http://127.0.0.1:8545/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// privateKey, err := crypto.HexToECDSA("69b39aa2fb86c7172d77d4b87b459ed7643c1e4b052536561e08d7d25592b373")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error casting public key to ECDSA")
	// }
	//
	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// auth := bind.NewKeyedTransactor(privateKey)
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)      // in wei
	// auth.GasLimit = uint64(3000000) // in units
	// auth.GasPrice = gasPrice
	//
	// // key, _ := crypto.GenerateKey()
	// // auth := bind.NewKeyedTransactor(key)
	//
	// // alloc := make(core.GenesisAlloc)
	// // alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	// // sim := backends.NewSimulatedBackend(alloc, 6000000)
	// addr, _, contract, err := verifier.DeployVerifier(auth, client)
	//
	// if err != nil {
	// 	log.Fatalf("could not deploy contract: %v", err)
	// }
	//
	// A := [...]*big.Int{Hex2BI("0x2f1db7ed3094b2962c9ae0e37a389d35ddf13b3db562e0cd062e502201f58e9f"), Hex2BI("0x2916ada7e690230422ff96abb3285b9c037c6dcc04fd36b3742484b979e44ea1")}
	// A_p := [...]*big.Int{Hex2BI("0x215413bd49dab0b7b48bdfa162e27d4544455e3c4f4dba5c877065d328e6a4b4"), Hex2BI("0x60a7af4f7b566535ca1be7aab1d9344544aa4fc7e4436104a93b6d6e5a04ce9")}
	// B := [2][2]*big.Int{{Hex2BI("0x2fdce36547c70eefdd025c19ed8b4333eb8da402d2a05601ed18b4439d422053"), Hex2BI("0x1abf8e4b45db855a19e3198c7e0b913cd91ba00581da5924ad675f956118fb4a")}, {Hex2BI("0x2487c070c6d86aa5a5568f51a2416062b79404a4d78299c9b615db4a86393715"), Hex2BI("0x19a7ed34fdf4b7c9d0dc48ac6762009736a44e366b1399d4e8d982dbad713d96")}}
	// B_p := [...]*big.Int{Hex2BI("0x25c601e2085cc0b42fa8be22933cbc61b5aa41914b00be6761111f548f05ae70"), Hex2BI("0x6b08e7b8f765bf895cc4cfd4dd906c23f322629be4f3fa7caf84062d3fab850")}
	// C := [...]*big.Int{Hex2BI("0x1199620115c33e607d064839f75187fd75654e37db2d85bd7abfd6185f1d580a"), Hex2BI("0x2a580cd37820657b3fb5608de812c5872c3534304ab93fec605a4e9cbc37f75d")}
	// C_p := [...]*big.Int{Hex2BI("0x15a25be848f85dcf2e5aef7f82fe0a0a3b4920ab70cb3562e2abd3d88d0c7746"), Hex2BI("0x12cbcd1006d443d26ef63bd1b07fdc72ef61f829ed345276ba6f8af2bd741145")}
	// H := [...]*big.Int{Hex2BI("0x21b63c44cd103725875aa34efc928318191a3a51dd35aee6b6a40c7c932c476c"), Hex2BI("0x270eeb8c586b6cd2b55ccab34e7bdc3688022bf2d0416375e6189e76a1ec3e9f")}
	// K := [...]*big.Int{Hex2BI("0x114f3c647f4e888b92e20d10150e5b0652aa442286becf9fa38984baa35527d5"), Hex2BI("0x79270f00ad8e237330020e4b16cefb3f95f256b59db2744e2093801097a6e6e")}
	//
	// input := [...]*big.Int{big.NewInt(2533788), big.NewInt(2533888), Hex2BI("0x187b247a7cfc9e02ceca0a063d086a216425e0c585a410ba8a5042e7d49440d5")}
	//
	// result, _ := contract.VerifyTx(nil, A, A_p, B, B_p, C, C_p, H, K, input)
	// _ = addr
	// _ = contract
	// fmt.Println(result)
}
