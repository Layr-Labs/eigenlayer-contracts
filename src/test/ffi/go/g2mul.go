package main

import (
	"fmt"
	"os"
	"strings"
	"math/big"
	"github.com/consensys/gnark-crypto/ecc/bn254"
)

func main() {
	//parse args
	arg1 := os.Args[1]
	n := new(big.Int)
	n, _ = n.SetString(arg1, 10)

	//g2 mul
	pubkey := new(bn254.G2Affine).ScalarMultiplication(GetG2Generator(), n)
	px := pubkey.X.String()
	py := pubkey.Y.String()

	//parse out point coords to big ints
	pxs := strings.Split(px, "+")	
	pxss := strings.Split(pxs[1], "*")

	pys := strings.Split(py, "+")	
	pyss := strings.Split(pys[1], "*")

	pxsInt := new(big.Int)
	pxsInt, _ = pxsInt.SetString(pxs[0], 10)

	pxssInt := new(big.Int)
	pxssInt, _ = pxssInt.SetString(pxss[0], 10)

	pysInt := new(big.Int)
	pysInt, _ = pysInt.SetString(pys[0], 10)

	pyssInt := new(big.Int)
	pyssInt, _ = pyssInt.SetString(pyss[0], 10)

	//swtich to print coord requested
	switch os.Args[2] {
	case "1":
		fmt.Printf("0x%064X", pxsInt)
	case "2":
		fmt.Printf("0x%064X", pxssInt)
	case "3":
		fmt.Printf("0x%064X", pysInt)
	case "4":
		fmt.Printf("0x%064X", pyssInt)
	}

}

func GetG2Generator() *bn254.G2Affine {
    g2Gen := new(bn254.G2Affine)
    g2Gen.X.SetString("10857046999023057135944570762232829481370756359578518086990519993285655852781",
        "11559732032986387107991004021392285783925812861821192530917403151452391805634")
    g2Gen.Y.SetString("8495653923123431417604973247489272438418190587263600148770280649306958101930",
        "4082367875863433681332203403145435568316851327593401208105741076214120093531")
    return g2Gen
}