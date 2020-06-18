/*
* Добавьте функцию снятия со счета Withdraw(amount int) bool
* в програму gopl.io/ch9/bank1. Результат должен указывать,
* прошла ли транзакция успешно или произошла ошибка из-за
* нехватки средств. Сообщение, отправляемое go-подпрограме
* монитора, должно содержать как снимаемую сумму, так и новый
* канал, по которому go-подпрограмма монитора сможет отправить
* булев результат функции Withdraw.
 */

package main

type WithDrawMsg struct {
	amount   int
	accessCh chan<- bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withDrawCh = make(chan WithDrawMsg)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	accessCh := make(chan bool)
	withDrawCh <- WithDrawMsg{amount, accessCh}
	return <-accessCh
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case drawMsg := <-withDrawCh:
			ok := Balance() >= drawMsg.amount
			if ok {
				balance -= drawMsg.amount
			}
			drawMsg.accessCh <- ok
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {

}
