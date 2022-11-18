package factory_method

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type ChinaCheesePizza struct {
}

func (c ChinaCheesePizza) Prepare() {
}

func (c ChinaCheesePizza) Bake() {
}

func (c ChinaCheesePizza) Cut() {
}

func (c ChinaCheesePizza) Box() {
}

type BasePizzaStore struct {
	store PizzaStore
}

//func (bps BasePizzaStore) CreatePizza(typ string) Pizza {
//	panic("can not be there") // 或者提供一个默认的披萨
//}

func (bps BasePizzaStore) OrderPizza(typ string) Pizza {
	pizza := bps.store.CreatePizza(typ)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

type PizzaStore interface {
	CreatePizza(typ string) Pizza
	//OrderPizza(typ string) Pizza
}

type ChinaPizzaStore struct {
	BasePizzaStore
}

func NewChinaPizzaStore() ChinaPizzaStore {
	store := ChinaPizzaStore{}
	store.BasePizzaStore = BasePizzaStore{store: store}
	return store
}

func (c ChinaPizzaStore) CreatePizza(typ string) Pizza {
	var pizza Pizza
	switch typ {
	case "cheese":
		pizza = new(ChinaCheesePizza)
	case "greek":
		pizza = new(ChinaGreekPizza)
	case "veggie":
		pizza = new(ChinaVeggiePizza)
	}
	return pizza
}

type AmericaPizzaStore struct {
	BasePizzaStore
}

func NewAmericaPizzaStore() AmericaPizzaStore {
	store := AmericaPizzaStore{}
	store.BasePizzaStore = BasePizzaStore{store: store}
	return store
}

func (c AmericaPizzaStore) CreatePizza(typ string) Pizza {
	var pizza Pizza
	switch typ {
	case "cheese":
		pizza = new(AmericaCheesePizza)
	case "greek":
		pizza = new(AmericaGreekPizza)
	case "veggie":
		pizza = new(AmericaVeggiePizza)
	}
	return pizza
}

func main() {
	chinaStore := NewChinaPizzaStore()
	chinaStore.OrderPizza("cheese")

	americaStore := NewAmericaPizzaStore()
	americaStore.OrderPizza("veggie")
}
