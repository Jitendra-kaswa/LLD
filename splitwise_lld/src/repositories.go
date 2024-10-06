package src

import "fmt"

// different repos
type ExpenseRepository interface {
	Save(expense *Expense) error
	FindByID(id string) (*Expense, error)
	Update(expense *Expense) error
	Delete(id string) error
	FindByGroupID(groupID string) ([]*Expense, error)
	FindByUserID(userID string) ([]*Expense, error)
}

type UserRepository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

type GroupRepository interface {
	Save(group *Group) error
	FindByID(id string) (*Group, error)
	Update(group *Group) error
	Delete(id string) error
}

// Expense repo implementation
type InMemoryExpenseRepository struct {
	expenses map[string]*Expense
}

func NewInMemoryExpenseRepository() ExpenseRepository {
	return &InMemoryExpenseRepository{
		expenses: make(map[string]*Expense),
	}
}

func (r *InMemoryExpenseRepository) Save(expense *Expense) error {
	r.expenses[expense.ID] = expense
	return nil
}

func (r *InMemoryExpenseRepository) FindByID(id string) (*Expense, error) {
	expense, ok := r.expenses[id]
	if !ok {
		return nil, fmt.Errorf("expense with ID %s not found", id)
	}
	return expense, nil
}

func (r *InMemoryExpenseRepository) Update(expense *Expense) error {
	r.expenses[expense.ID] = expense
	return nil
}

func (r *InMemoryExpenseRepository) Delete(id string) error {
	delete(r.expenses, id)
	return nil
}

func (r *InMemoryExpenseRepository) FindByGroupID(groupID string) ([]*Expense, error) {
	var groupExpenses []*Expense
	for _, expense := range r.expenses {
		if expense.Group != nil && expense.Group.ID == groupID {
			groupExpenses = append(groupExpenses, expense)
		}
	}
	return groupExpenses, nil
}

func (r *InMemoryExpenseRepository) FindByUserID(userID string) ([]*Expense, error) {
	var userExpenses []*Expense
	for _, expense := range r.expenses {
		if expense.Payer.ID == userID {
			userExpenses = append(userExpenses, expense)
			continue
		}
		for _, split := range expense.Splits {
			if split.User.ID == userID {
				userExpenses = append(userExpenses, expense)
				break
			}
		}
	}
	return userExpenses, nil
}

// user repository implementation
type InMemoryUserRepository struct {
	users map[string]*User
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*User),
	}
}

func (r *InMemoryUserRepository) Save(user *User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (r *InMemoryUserRepository) Update(user *User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	delete(r.users, id)
	return nil
}

// Group repository implementation
type InMemoryGroupRepository struct {
	groups map[string]*Group
}

func NewInMemoryGroupRepository() GroupRepository {
	return &InMemoryGroupRepository{
		groups: make(map[string]*Group),
	}
}

func (r *InMemoryGroupRepository) Save(group *Group) error {
	r.groups[group.ID] = group
	return nil
}

func (r *InMemoryGroupRepository) FindByID(id string) (*Group, error) {
	group, ok := r.groups[id]
	if !ok {
		return nil, fmt.Errorf("group with ID %s not found", id)
	}
	return group, nil
}

func (r *InMemoryGroupRepository) Update(group *Group) error {
	r.groups[group.ID] = group
	return nil
}

func (r *InMemoryGroupRepository) Delete(id string) error {
	delete(r.groups, id)
	return nil
}
