package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func printUsage(){
	fmt.Print("Usage:")
	fmt.Print("\taddBlock -data DATA -- 交易数据")
	fmt.Print("\tprintchian -- 输出区块信息")
}

func isValidArgs(){
	if len(os.Args) < 2{
		printUsage()
		os.Exit(1)
	}
}


func main() {

	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "hah", "交易数据......")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed(){
		if *flagAddBlockData == ""{
			printUsage()
			os.Exit(1)
		}
		fmt.Println(*flagAddBlockData)
	}

	if printChainCmd.Parsed(){
		fmt.Println("输出所有的数据......")
	}

}
