## Build shared object (`.so`) files.

```bash
$ go build -buildmode=plugin -o hourly_employee/hourly_employee.so hourly_employee/hourly_employee.go
$ go build -buildmode=plugin -o salaried_employee/salaried_employee.so salaried_employee/salaried_employee.go
```

## Check file types

```bash
$ file hourly_employee/hourly_employee.so
$ file salaried_employee/salaried_employee.so
```

Depending on your OS, you'll see something like:

```bash
$ file hourly_employee/hourly_employee.so
hourly_employee/hourly_employee.so: ELF 64-bit LSB shared object, x86-64, version 1 (SYSV), dynamically linked, BuildID...
```
