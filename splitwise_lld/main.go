package main

import (
	"fmt"

	"splitwise_lld.com/src"
)

func main() {
	balanceCalculationStrategy := src.NewSimpleBalanceCalculationStrategy()
	idGenerator := src.NewIdGenerationUsingUUID()
	expenseRepo := src.NewInMemoryExpenseRepository()
	userRepo := src.NewInMemoryUserRepository()
	groupRepo := src.NewInMemoryGroupRepository()
	splitwiseService := src.NewSplitwiseService(expenseRepo, userRepo, groupRepo, idGenerator, balanceCalculationStrategy)
	equalSplitStrategy := src.NewEqualSplitStrategy()

	// Create users
	alice, _ := splitwiseService.CreateUser("Alice", "alice@example.com")
	bob, _ := splitwiseService.CreateUser("Bob", "bob@example.com")
	charlie, _ := splitwiseService.CreateUser("Charlie", "charlie@example.com")

	// Create a group
	group, _ := splitwiseService.CreateGroup("Friends", []string{alice.ID, bob.ID, charlie.ID})

	// Create an expense
	expense, _ := splitwiseService.CreateExpense(
		"Dinner",
		9000.0,
		alice.ID,
		[]string{alice.ID, bob.ID, charlie.ID},
		equalSplitStrategy,
		group.ID,
	)

	fmt.Printf("Expense created: %+v\n", expense)

	_ = splitwiseService.SettleExpense(expense.ID, bob.ID)

	status, _ := splitwiseService.GetExpenseStatus(expense.ID)
	fmt.Printf("Expense status: %v\n", status)

	_ = splitwiseService.SettleExpense(expense.ID, charlie.ID)

	status, _ = splitwiseService.GetExpenseStatus(expense.ID)
	fmt.Printf("Expense status after full settlement: %v\n", status)

	groupBalances, _ := splitwiseService.GetGroupBalances(group.ID)
	fmt.Printf("Group balances: %v\n", groupBalances)

	dave, _ := splitwiseService.CreateUser("Dave", "dave@example.com")
	_ = splitwiseService.AddUserToGroup(group.ID, dave.ID)
	fmt.Printf("Added Dave to the group\n")

	_ = splitwiseService.RemoveUserFromGroup(group.ID, charlie.ID)
	fmt.Printf("Removed Charlie from the group\n")

	aliceExpenses, _ := splitwiseService.GetUserExpenses(alice.ID)
	fmt.Printf("Alice's expenses: %+v\n", aliceExpenses)

	groupExpenses, _ := splitwiseService.GetGroupExpenses(group.ID)
	fmt.Printf("Group expenses: %+v\n", groupExpenses)
}
