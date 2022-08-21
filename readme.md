# 

## Introduction

This is an experimental repo. My goal is to parse the VE.Direct protocol from Victron with [TinyGo](https://tinygo.org/) on a [Raspberry Pico](https://www.raspberrypi.com/products/raspberry-pi-pico/).

See specifications of the protocol [here](/doc/VE.Direct-Protocol-3.32.pdf)

I want to have a display in my van to see :
- Battery voltage and current
- Solar panel Voltage and Current
- Current State of the MPTT Controller
- Maximum power
- Minimum and maximum voltage

### Flash the pico 

```
tinygo flash -target=pico main.go
```

### Testing the VE.Direct parsing lib

```
go test -v ./vedirect
```

### Buil and see alloc

```
tinygo build -o firmware.uf2 -target=pico -print-allocs=. main.go
```