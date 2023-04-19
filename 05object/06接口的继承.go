package main

/*
接口的继承可以分为显式继承和隐式继承
*/
type Fighter interface {
	Attack() (bloodloss int)
	Defend()
}
type Beast interface {
	// Animal 显式地继承Animal
	Animal
	// Attack 隐式地继承Fighter接口
	Attack() (bloodloss int)
	Defend()
}

func main061() {

}
