package main

import "github.com/go-packagist/container"

func main() {
	c := container.NewContainer()

	c.Instance("aa", "aaa")
	c.Bind("bb", func(c *container.Container) interface{} {
		return "bbb"
	}, false)
	c.Singleton("cc", func(c *container.Container) interface{} {
		return "ccc"
	})
	c.Singleton("dd", func(c *container.Container) interface{} {
		return func(hello string) string {
			return "hello " + hello
		}
	})

	var aa, bb, cc string
	c.Make("aa", &aa)
	c.Make("bb", &bb)
	c.Make("cc", &cc)
	println(aa, bb, cc) // aaa bbb ccc

	var dd func(name string) string
	c.Make("dd", &dd)
	println(dd("world")) // hello world
}
