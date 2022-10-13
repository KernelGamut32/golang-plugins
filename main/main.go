package main

import (
	"fmt"
	"os"
	"plugin"
	"strconv"

	"github.com/KernelGamut32/golang-plugins/employee"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: run main/main.go <employee_type> <pay_rate>")
		os.Exit(1)
	}

	employee_type := os.Args[1]
	pay_rate, parse_err := strconv.ParseFloat(os.Args[2], 64)
	if parse_err != nil {
		fmt.Println(parse_err)
		os.Exit(1)
	}
	module := fmt.Sprintf("./%s/%s.so", employee_type, employee_type)

	_, err := os.Stat(module)
	if os.IsNotExist(err) {
		fmt.Println("can't find employee type", employee_type)
		os.Exit(1)
	}

	p, err := plugin.Open(module)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symbol, err := p.Lookup("Employee")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	employee, ok := symbol.(employee.Employee)
	if !ok {
		fmt.Println("that's not a valid employee type")
		os.Exit(1)
	}

	fmt.Printf("Weekly pay for %s equates to %.2f\n", employee_type, employee.ProcessPay(pay_rate))
}
