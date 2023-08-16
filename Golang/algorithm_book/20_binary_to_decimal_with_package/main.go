package main

import (
	"fmt"
	"strconv"
)

var num, ori, des int

طfunc main() {
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)

	fmt.Print("Enter the base of origin: ")
	fmt.Scanln(&ori)

	fmt.Print("Enter the base of destination: ")
	fmt.Scanln(&des)

	//فانکشن strconv.Itoa(int) string
	//اینت برمبنای 10 میگیرد و استرینگ بر مبنای 10 میدهد تا ورودی/خوراک فانکشن بعدی را تامین کند
	//فاکشن strconv.ParseInt(string, int , int) (int64 , error)
	// استرینگ بر مبنای 10 و اینت (مبنای دلخواه) و اینت (تعداد بیت های عدد) را میگیرد و اینت64 بر مبنای دلخواه و یک ارور میدهد
	//خروجی اش را ذخیره میکند convnumدر
	convnum, err := strconv.ParseInt(strconv.Itoa(num), ori, 64)
	if err != nil {
		fmt.Printf("Error converting number: %v\n", err)
		return
	}
	//فانکشن strconv.FormatInt(int64 , int) string
	//یک اینت64 و اینت (مبنای دلخواه) میگیرد و یک استرینگ بر اساس همان مبنا میدهد
	result := strconv.FormatInt(convnum, des)
	//نتیجه: یک اینت64 با مبنای دلخواه درست کردیم و دادیم به فانکشن خط بالا و یک استرینگ با مبنای دلخواه تولید کردیم
	fmt.Printf("%v based on %v = %v based on %v.\n", num, ori, result, des)
}