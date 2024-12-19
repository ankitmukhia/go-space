package main

/* Note: bs (*bookStore) and (&bs) actually mean that we passed the actual address of the bookStore instance, allowing the method to directly manipulate the original object rather than working with a copy of it. */

import ( 
	"fmt" 
)

type store[t product] interface {
	Sell(t)
}

type product interface {
	Price() float64
	Name() string 
}

type book struct {
	title string
	author string
	price float64	
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
	name string
	price float64 
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

// it represents store, that seels books

type bookStore struct {
	// bookCon -> book counter
	bookCou []book
}

/* Sell is generic, so  it takes any kind of parameter that fulfills product interface */
// (bs *bookStore) translate taht method is pointer receiver, it operates on a pointer instance rather than copy  of the bookStore

// Sell adds a book to the bookStore's inventory.
func (bs *bookStore) Sell(b book) {
	bs.bookCou = append(bs.bookCou, b)
}

type toyStore struct {
	toyCou []toy
}

// Sell adds a toy to the toyStore's inventory.
func (ts *toyStore) Sell(t toy) {
	ts.toyCou = append(ts.toyCou, t)	
}


func sellProducts[P product](s store[P], product []P){
	for _, p := range product {
		s.Sell(p)
	}
}

func main() {
	bs := bookStore{
		bookCou: []book{},
	}
	// we are defining the type doing this [book], sellProducts func is generic so it can work will any type, so we are saying it will treat sellProducts as a book type, same goes for toy. And we don't have to explictly do this, Go compiler smart enough to know this, figure out type.
	sellProducts[book](&bs, []book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Println(bs.bookCou)

	ts := toyStore{
		toyCou: []toy{},
	}

	// Use this specific type(i,e: toy) for this function call.
	sellProducts[toy](&ts, []toy{
		{
			name:  "Lego",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Println(ts.toyCou)
}
