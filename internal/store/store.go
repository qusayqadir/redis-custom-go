// the key-vallue expire 
// store is the database 

/* 
*/ 
package store 

type Store struct{
	data map[string]string
}

func New() *Store {
	return &Store{
		data : make(map[string]string)}
}

func GET(){

}

func SET() {

}

func DEL() {

}