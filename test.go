package main

//This test cannot be run because of the gopath settings
//It is only an example, if you wanna try it copy paste this code in your environment
//And import the graph package correctly

import (
	"fmt"
)

func main() {
	g := new(Graph)
	g.Add("test")
	g.Add("cool")
	g.Add("graph")
	g.Connect("test", "cool", 3)
	g.Connect("graph", "cool", 2)
	g.Connect("graph", "test", 5)
	fmt.Println(g.InGraph("test"))
	fmt.Println(g.InGraph("cool"))
	fmt.Println(g.InGraph("cool"))
	fmt.Println(g.InGraph("test"))
	g.SetValue("test", 2)
	fmt.Println(g.GetEdgeValue("test", "cool"))
	fmt.Println(g.Neighbors("cool")[0].Value)
	g.Disconnect("test", "cool")
	fmt.Println(g.InGraph("test"))
	g.Remove("test")
	fmt.Println(g.InGraph("test"))
}
