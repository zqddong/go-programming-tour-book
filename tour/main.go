package main

import (
	"github.com/zqddong/go-programming-tour-book/tour/cmd"
	"log"
)

//var name string

//type Name string
//
//func (n *Name) String() string {
//	return fmt.Sprint(*n)
//}
//
//func (n *Name) Set(value string) error {
//	if len(*n) > 0 {
//		return errors.New("name flag already set")
//	}
//
//	*n = Name("name:" + value)
//	return nil
//}

// go run main.go go -name=golang
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

	//var name Name
	//flag.Var(&name, "name", "帮助信息")
	//flag.Parse()
	//log.Printf("name: %s", name)

	//
	//flag.Parse()
	//args := flag.Args()
	//if len(args) <= 0 {
	//	return
	//}
	//
	//switch args[0] {
	//case "go":
	//	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	//	goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
	//	_ = goCmd.Parse(args[1:])
	//case "php":
	//	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	//	phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
	//	_ = phpCmd.Parse(args[1:])
	//}
	//
	//log.Printf("name: %s", name)
}
