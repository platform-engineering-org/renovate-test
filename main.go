package main

import (
	"fmt"
	_ "github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	_ "github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
)

func main() {
	fmt.Println("Hello")
}
