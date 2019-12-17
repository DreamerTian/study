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
	fmt.Println("\tcreateblock -data DATA（交易数据） --创建区块命令")
	fmt.Println("\taddblock -data DATA(交易数据) -- 增加新的区块")
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

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createblock", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "hah", "交易数据......")
	flagCreateBlockChainData := createBlockChainCmd.String("data", "Block data......", "交易数据......")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
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

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		cli.addBlock([]*Transaction{})
	}

	if printChainCmd.Parsed() {
		cli.printchain()
	}

	if createBlockChainCmd.Parsed(){
		if *flagCreateBlockChainData == ""{
			fmt.Println("交易数据不能为空")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain([]*Transaction{})
	}
}

func (cli *CLI) createGenesisBlockchain(txs []*Transaction) {
	CreateBlockchainWithGenesisBlock(txs)
}
