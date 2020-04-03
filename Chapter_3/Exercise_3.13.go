/*
* Напишите объявления const для KB, MB и так далее
* до YB настолько компактно, насколько сможете.
 */

package main

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

}
