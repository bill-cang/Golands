package main

type DogFc interface {
	//继承animal
	AnimalFc
	//独有方法
	Run()
}

type DogEntity struct {
	Name string
	Age int
}

func (dog *DogEntity) SetName(name string) {
	dog.Name = name
}

func (dog *DogEntity) GetName() string {
	return dog.Name
}




