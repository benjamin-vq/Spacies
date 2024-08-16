package main

type Config struct {
	screenWidth  int
	screenHeight int
}

func NewConfig(width, height int) Config {
	return Config{
		screenWidth:  width,
		screenHeight: height,
	}
}
