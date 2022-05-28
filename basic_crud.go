package main

import "fmt"

type Item struct{
	title string
	body string
}


var database []Item

func GetByName(title string) Item{
	var getItem Item
	for _, val := range database{
		if val.title == title{
			getItem = val
		}
	}
	return getItem
}

func AddItem(item Item) Item{
	database  = append(database, item)
	return item
}



func EditItem(title string, edit Item) Item{

	var changed Item

	for idx, val := range database {
		if val.title == title {

			database[idx] = edit
			changed = edit
		}
	}

	return changed
}


func DeleteItem(item Item) Item{
	var del Item

	for idx, val := range database{
		if val.title == item.title && val.body == item.body{
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	return del
}

func main(){

	fmt.Println("intial database:[no items] ", database)

	a := Item{"first", "the first item"}
	b := Item{"second", "the second item"}
	c := Item{"third", "the third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	//database after adding items
	fmt.Println("database 2:", database)

	//delete 'b'
	DeleteItem(b)
	//database after deleting item 'b'
	// now only 2 items left : 'a' and 'c'
	fmt.Println("database 3: ", database)


	//'b' is "second" item => {"second", "the second item"}

	//edit item "third" item => {"third", "the third item"}

	EditItem("third", Item{"fourth", "the fourth item"})

	fmt.Println("fourth database: ", database)

	x := GetByName("fourth") // 'a'
	y := GetByName("first") // 'c'
	fmt.Println(x, y)
	fmt.Println("\n", database)
}

