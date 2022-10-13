.DEFAULT_GOAL := help

help:
	# Commands:
	# make build  - Build the shared object (.so) library files
	# make hourly_employee   - Run the hourly_employee example
	# make salaried_employee   - Run the salaried_employee example

build:
	@go build -buildmode=plugin -o hourly_employee/hourly_employee.so hourly_employee/hourly_employee.go
	@go build -buildmode=plugin -o salaried_employee/salaried_employee.so salaried_employee/salaried_employee.go

hourly_employee: build
	@go run main/main.go hourly_employee 15.00

salaried_employee: build
	@go run main/main.go salaried_employee 30.00
