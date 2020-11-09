package color

import "fmt"

func init() {
	c := map[string]string{
		"red":   "#ff0000",
		"green": "#0000ff",
		"white": "#ffffff",
	}

	Print(c)
}

func Print(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, v)
	}
}
