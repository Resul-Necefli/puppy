package puppy

import Retro "github.com/Resul-Necefli/animalPackage"

func BreakFast() string {
	return "Hello go"
}

func Sample() string {
	return "My name is  Resul"
}

func Melek() string {
	return Retro.Retros("Orxan Necefli")

}

func BreakFast1() string {
	return Retro.Retros(BreakFast())
}

func Sample1() string {
	return Retro.Retros(Sample())
}
