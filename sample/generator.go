package sample

import (
	"./pb",
	"math/rand"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout: randowKeyboardLayout(),
	}

}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 16)
	minGhz := randomFloat64(2.0, 5.0)
	maxGhz := randomFloat64(minGhz, 5.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   numberCores,
		NumberThreads: numberThreads,
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 2.0)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 12)
	memory := &pb.Memory{
		Value: unit64(randomInt(2, 12)),
		Unit:  pb.Memory_GIGABYTE,
	}
	gpu := &pb.GPU{
		Brand:    brand,
		Name:     name,
		MinGhz:   minGhz,
		MaxGhz:   maxGhz,
		MemoryGB: memGB,
		Memory:   memory,
	}
	return gpu
}

func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

func NewSSD() *pb.Storage {
	ssd: &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

func NewHDD() *pb.Storage {
	hdd: &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewScreen() *pb.Screen {
	width := randomInt(13, 17)
	height := width * 9 / 16
	screen := &pb.Screen{
		Resolution: &pb.Screen_Resolution{
			Width:  uint32(width),
			Height: uint32(height),
		},
		Panel: pb.Screen_IPS,
		MultiTouch: randomBool(),
	}
	return screen
}