package main

import (
	"fmt"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/skitt/viper"
	_ "github.com/skitt/viper/remote"
)

func main() {
	viper.SetDefault("Foo", "bar")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.viper-demo")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Couldn't read configuration: %w\n", err)
	}
	fmt.Printf("Foo is %s\n", viper.Get("Foo"))
	fmt.Printf("%v\n", difflib.SplitLines("hello\nworld"))
}
