package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateblock -address --创建区块命令")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -- 转账")
	fmt.Println("\tprintchain -- 打印区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(txs []*Transaction) {

	if DbExists() == false{
		fmt.Println("请先创建创世区块.....")
		os.Exit(1)
	}

	blockchain := BlockchainObject()
	defer blockchain.DB.Close()


	blockchain.AddBlockToBlockchain(txs)
}

func (cli *CLI) printchain() {
	if DbExists() == false{
		fmt.Println("请先创建创世区块.....")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.Printchain()
}

func (cli *CLI) Run() {
	isValidArgs()

	sendCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createblock", flag.ExitOnError)

	flagSendFrom := sendCmd.String("from", "", "源地址")
	flagSendTo := sendCmd.String("to", "", "目的地地址")
	flagAmount := sendCmd.String("amount", "", "转账的金额")
	flagCreateBlockChainAddress := createBlockChainCmd.String("address", "", "创建创世区块的地址")

	switch os.Args[1] {
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblock":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if sendCmd.Parsed() {
		if *flagSendFrom == "" || *flagSendTo == "" || *flagAmount == ""{
			printUsage()
			os.Exit(1)
		}
		//cli.addBlock([]*Transaction{})
		fmt.Println(*flagSendFrom)
		fmt.Println(*flagSendTo)
		fmt.Println(*flagAmount)
		fmt.Println("--------------------------------")
		fmt.Println(JsonToArray(*flagSendFrom))
		fmt.Println(JsonToArray(*flagSendTo))
		fmt.Println(JsonToArray(*flagAmount))
	}

	if printChainCmd.Parsed() {
		cli.printchain()
	}

	if createBlockChainCmd.Parsed(){
		if *flagCreateBlockChainAddress == ""{
			fmt.Println("地址不能为空......")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockChainAddress)
	}
}

func (cli *CLI) createGenesisBlockchain(address string) {
	CreateBlockchainWithGenesisBlock(address)
}
