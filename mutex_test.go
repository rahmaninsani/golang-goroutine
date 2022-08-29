package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
TestMutex is Race Condition Solution with Mutex
*/
func TestMutex(t *testing.T) {
	var mutex sync.Mutex
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", x)
}

// RWMutex
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

// Write
func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

// Read
func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("Balance:", account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}

// Deadlock
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	fmt.Println("Unlock user1", user1.Name)

	user2.Unlock()
	fmt.Println("Unlock user2", user2.Name)
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{Name: "Alice", Balance: 1000000}
	user2 := UserBalance{Name: "Bob", Balance: 1000000}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(10 * time.Second)

	fmt.Println("User", user1.Name, "- Balance:", user1.Balance)
	fmt.Println("User", user2.Name, "- Balance:", user2.Balance)

}
