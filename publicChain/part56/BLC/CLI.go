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
	fmt.Println("\tcreatewallet --创建钱包")
	fmt.Println("\taddresslists --展示所有钱包地址")
	fmt.Println("\tcreateblock -address --创建区块命令")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -- 转账")
	fmt.Println("\tprintchain -- 打印区块信息")
	fmt.Println("\tgetbalance -address -- 打印区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}


func (cli *CLI) Run() {
	isValidArgs()

	sendCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createblock", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	addressListsCmd := flag.NewFlagSet("addresslists", flag.ExitOnError)

	flagSendFrom := sendCmd.String("from", "", "源地址")
	flagSendTo := sendCmd.String("to", "", "目的地地址")
	flagAmount := sendCmd.String("amount", "", "转账的金额")
	flagCreateBlockChainAddress := createBlockChainCmd.String("address", "", "创建创世区块的地址")
	flagGetBalance := getBalanceCmd.String("address","","")

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
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addresslists":
		err := addressListsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if sendCmd.Parsed() {
		if *flagSendFrom == "" || *flagSendTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}


		from := JsonToArray(*flagSendFrom)
		to := JsonToArray(*flagSendTo)

		for index,forAddress := range from{
			if IsValidForAddress([]byte(forAddress)) == false || IsValidForAddress([]byte(to[index])) == false{
				fmt.Println("地址无效......")
				os.Exit(1)
			}
		}
		amount := JsonToArray(*flagAmount)

		cli.send(from, to, amount)
	}

	if printChainCmd.Parsed() {
		cli.printchain()
	}

	if createBlockChainCmd.Parsed() {
		if IsValidForAddress([]byte(*flagCreateBlockChainAddress)) == false {
			fmt.Println("地址无效......")
			printUsage()
			os.Exit(1)
		}



		cli.createGenesisBlockchain(*flagCreateBlockChainAddress)
	}
	if getBalanceCmd.Parsed() {
		if IsValidForAddress([]byte(*flagGetBalance)) == false  {
			fmt.Println("地址无效......")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalance)
	}

	//创建钱包
	if createWalletCmd.Parsed(){
		cli.createWallet()
	}

	if addressListsCmd.Parsed(){
		cli.addressLists()
	}
}





