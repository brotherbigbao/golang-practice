package animal

type Cat struct {
	name string
	age int
	color string
	price float32
}

func (c *Cat) GetName() string {
	return c.name
}

func (c *Cat) GetAge() int {
	return c.age
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetPrice() float32 {
	return c.price
}

func (c *Cat) SetName(name string) {
	c.name = name
}

func (c *Cat) SetAge(age int) {
	c.age = age
}

func (c *Cat) SetColor(color string) {
	c.color = color
}

func (c *Cat) SetPrice(price float32) {
	c.price = price
}