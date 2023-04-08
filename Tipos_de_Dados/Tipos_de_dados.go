package main

import (
	"errors"
	"fmt"
)

func main() {
	// int tem varios int8,int16,int32,int64
	numero_int := 562363636
	fmt.Println(numero_int)

	// uint: uint8,uint16,uint32,uint64 
	var numero_unit uint = 5623
	fmt.Println(numero_unit)

	//alias
	//int32 == rune
	var numero_3 rune = 3525
	fmt.Println(numero_3)

	// byte = uint8
	var numero_4 byte = 32
	fmt.Println(numero_4)

	//float32, float64
	var numero_double float32 = 3421.63
	var numero_double2 float64 = 452562.36
	numero_double3 := 525525.345
	fmt.Println(numero_double,"\n",numero_double2,"\n",numero_double3)

	//String
	var string_1 string = "jose"
	string_2 := "joão"
	fmt.Println(string_1)
	fmt.Println(string_2)
	char := 'B'
	fmt.Println(char)

	//bool
	var booleano1 bool
	fmt.Println(booleano1)

	//error
	var erro error = errors.New("Erro Go")
	fmt.Println(erro)
}