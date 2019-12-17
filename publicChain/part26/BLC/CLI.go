package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	Blockchain *Blockchain
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -data DATA -- 交易数据")
	fmt.Println("\taddblock -data DATA -- 交易数据")
	fmt.Println("\tprintchian -- 输出区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.Blockchain.AddBlockToBlockchain(data)
}

func (cli *CLI) printchain() {
	cli.Blockchain.Printchain()
}

func (cli *CLI) Run() {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

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
	case "createblockchain":
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
		cli.addBlock(*flagAddBlockData)
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
		cli.createGenesisBlockchain(*flagCreateBlockChainData)
	}
}

func (cli *CLI) createGenesisBlockchain(s string) {
	CreateBlockchainWithGenesisBlock(s)
}
