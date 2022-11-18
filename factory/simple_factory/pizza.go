package simple_factory

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type CheesePizza struct {
}

type GreekPizza struct {
}

type PepperoniPizza struct {
}

type VeggiePizza struct {
}

type PizzaStore struct {
	factory SimplePizzaFactory
}

func (p PizzaStore) OrderPizza(typ string) Pizza {
	pizza := p.factory.CreatePizza(typ)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

type SimplePizzaFactory struct {
}

func (spf SimplePizzaFactory) CreatePizza(typ string) Pizza {
	var pizza Pizza
	switch typ {
	case "cheese":
		pizza = new(CheesePizza)
	case "greek":
		pizza = new(GreekPizza)
	case "veggie":
		pizza = new(VeggiePizza)
	}
	return pizza
}
