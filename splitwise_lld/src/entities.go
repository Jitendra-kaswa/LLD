package src

import "encoding/json"

type User struct {
	ID    string
	Name  string
	Email string
}

func (u *User) String() string {
	data, _ := json.MarshalIndent(u, "", "  ")
	return string(data)
}

type Group struct {
	ID      string
	Name    string
	Members []*User
}

func (g *Group) String() string {
	data, _ := json.MarshalIndent(g, "", "  ")
	return string(data)
}

type Expense struct {
	ID            string
	Description   string
	Amount        float64
	Payer         *User
	Participants  []*User
	Splits        []*Split
	Status        ExpenseStatus
	SplitStrategy SplitStrategy
	Group         *Group // Optional, can be nil for non-group expenses
}

func (e *Expense) String() string {
	data, _ := json.MarshalIndent(e, "", "  ")
	return string(data)
}

type Split struct {
	User   *User
	Amount float64
	Status SplitStatus
}

func (s *Split) String() string {
	data, _ := json.MarshalIndent(s, "", "  ")
	return string(data)
}

func NewUser(name, email string, id string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func NewGroup(name string, members []*User, id string) *Group {
	return &Group{
		ID:      id,
		Name:    name,
		Members: members,
	}
}

func NewExpense(description string, amount float64, payer *User, participants []*User, splitStrategy SplitStrategy, group *Group, id string) *Expense {
	return &Expense{
		ID:            id,
		Description:   description,
		Amount:        amount,
		Payer:         payer,
		Participants:  participants,
		Status:        CREATED,
		SplitStrategy: splitStrategy,
		Group:         group,
	}
}

func NewSplit(user *User, amount float64) *Split {
	return &Split{
		User:   user,
		Amount: amount,
		Status: UNPAID,
	}
}
