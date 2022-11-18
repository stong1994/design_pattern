package abstract_factory

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type ChinaCheesePizza struct {
	ingredient PizzaIngredientFactory
	dough      Dough
	sauce      Sauce
}

func NewChinaCheesePizza(ingredient PizzaIngredientFactory) *ChinaCheesePizza {
	return &ChinaCheesePizza{
		ingredient: ingredient,
	}
}

func (c *ChinaCheesePizza) Prepare() {
	c.dough = c.ingredient.CreateDough()
	c.sauce = c.ingredient.CreateSauce()
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
	ingredient := ChinaPizzaIngredientFactory{}
	var pizza Pizza
	switch typ {
	case "cheese":
		//pizza = new(ChinaCheesePizza)
		pizza = NewChinaCheesePizza(ingredient)
	case "greek":
		//pizza = new(ChinaGreekPizza)
		pizza = NewChinaGreekPizza(ingredient)
	case "veggie":
		//pizza = new(ChinaVeggiePizza)
		pizza = NewChinaVeggiePizza(ingredient)
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

type Dough struct {
	Weight int
}

type Sauce struct {
	Weight int
}

type PizzaIngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
}

type ChinaPizzaIngredientFactory struct {
}

func (c ChinaPizzaIngredientFactory) CreateDough() Dough {
	return Dough{Weight: 10}
}

func (c ChinaPizzaIngredientFactory) CreateSauce() Sauce {
	return Sauce{Weight: 20}
}

type AmericaPizzaIngredientFactory struct {
}

func (c AmericaPizzaIngredientFactory) CreateDough() Dough {
	return Dough{Weight: 10}
}

func (c AmericaPizzaIngredientFactory) CreateSauce() Sauce {
	return Sauce{Weight: 20}
}
