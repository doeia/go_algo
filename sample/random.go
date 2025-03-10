package sample

import (
	"../pb"
	"math/rand"
)

func randowKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 0:
		return pb.Keyboard_QWERTY
	case 1:
		return pb.Keyboard_QWERTZ
	case 2:
		return pb.Keyboard_AZERTY
	}
	return pb.Keyboard_QWERTY
}

func randomCPUBrand() string {
	return randomStringFromSet("AMD", "Intel")
}

func randomCPUName(brand string) string {
	if brand == "AMD" {
		return randomStringFromSet("Ryzen 7", "Ryzen 5", "Ryzen 3")
	}
	return randomStringFromSet("Core i9", "Core i7", "Core i5", "Core i3")
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet("RTX 2080 Ti", "RTX 2080", "GTX 1660 Ti")
	}
	return randomStringFromSet("RX 590", "RX 580", "RX 570")
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
