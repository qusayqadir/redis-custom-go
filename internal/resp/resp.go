// redis serialization protocol (*resp) - wire format that redis clients 
// and servers use to talk to each other over tcp connection. 

/*
every byte of data over the wire comes in the format \r\n 
:1000\r\n

+ simple string  
- error 
: integer 
$ Bulk String	$5\r\nhello\r\n (the 5 is the byte length) 
* Array			*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n 

string commands 
set [ key, value] - set a key 
get [ value ] - retrieve key 
delete key [key... ] one or more keys 
exits key - checks if a key exists 
expire key seconds - set a ttl on a key 
ttl key - check remaining time to live 

list commands 
hash commands 
set commands 
sorted commands 
sorted set commands 
connection / server commands 
pub/sub  (after server can handle multiple tcp connections )

ping -> get/set/del -> expiry/ttl -> etc 



decode client raw byte into a command
encode reply back into RESP format 
*/ 

package resp 