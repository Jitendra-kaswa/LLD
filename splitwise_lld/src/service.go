package src

import "errors"

type ISplitwiseService interface {
	CreateUser(name, email string) (*User, error)
	CreateGroup(name string, memberIDs []string) (*Group, error)
	CreateExpense(description string, amount float64, payerID string, participantIDs []string, splitStrategy SplitStrategy, groupID string) (*Expense, error)
	SettleExpense(expenseID string, settlerID string) error
	GetExpenseStatus(expenseID string) (*ExpenseStatus, error)
	GetGroupBalances(groupID string) (map[string]float64, error)
	AddUserToGroup(groupID string, userID string) error
	RemoveUserFromGroup(groupID string, userID string) error
	GetUserExpenses(userID string) ([]*Expense, error)
	GetGroupExpenses(groupID string) ([]*Expense, error)
}

type SplitwiseService struct {
	expenseRepo                ExpenseRepository
	userRepo                   UserRepository
	groupRepo                  GroupRepository
	idGenerator                IdGenerationStrategy
	balanceCalculationStrategy BalanceStrategy
}

func NewSplitwiseService(expenseRepo ExpenseRepository, userRepo UserRepository, groupRepo GroupRepository, idGenerator IdGenerationStrategy, balanceCalculationStrategy BalanceStrategy) ISplitwiseService {
	return &SplitwiseService{
		expenseRepo:                expenseRepo,
		userRepo:                   userRepo,
		groupRepo:                  groupRepo,
		idGenerator:                idGenerator,
		balanceCalculationStrategy: balanceCalculationStrategy,
	}
}

func (s *SplitwiseService) CreateUser(name, email string) (*User, error) {
	user := NewUser(name, email, s.idGenerator.GenerateId())
	err := s.userRepo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *SplitwiseService) CreateGroup(name string, memberIDs []string) (*Group, error) {
	members := make([]*User, len(memberIDs))
	for i, id := range memberIDs {
		user, err := s.userRepo.FindByID(id)
		if err != nil {
			return nil, err
		}
		members[i] = user
	}

	group := NewGroup(name, members, s.idGenerator.GenerateId())
	err := s.groupRepo.Save(group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (s *SplitwiseService) CreateExpense(description string, amount float64, payerID string, participantIDs []string, splitStrategy SplitStrategy, groupID string) (*Expense, error) {
	payer, err := s.userRepo.FindByID(payerID)
	if err != nil {
		return nil, err
	}

	participants := make([]*User, len(participantIDs))
	for i, id := range participantIDs {
		user, err := s.userRepo.FindByID(id)
		if err != nil {
			return nil, err
		}
		participants[i] = user
	}

	var group *Group
	if groupID != "" {
		group, err = s.groupRepo.FindByID(groupID)
		if err != nil {
			return nil, err
		}
	}

	expense := NewExpense(description, amount, payer, participants, splitStrategy, group, s.idGenerator.GenerateId())

	splits, err := splitStrategy.Split(expense)
	if err != nil {
		return nil, err
	}
	expense.Splits = splits

	err = s.expenseRepo.Save(expense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (s *SplitwiseService) SettleExpense(expenseID string, settlerID string) error {
	expense, err := s.expenseRepo.FindByID(expenseID)
	if err != nil {
		return err
	}

	for _, split := range expense.Splits {
		if split.User.ID == settlerID {
			split.Status = PAID
			break
		}
	}

	allSettled := true
	for _, split := range expense.Splits {
		if split.Status == UNPAID {
			allSettled = false
			break
		}
	}

	if allSettled {
		expense.Status = SETTLED
	} else {
		expense.Status = PARTIALLY_SETTLED
	}

	return s.expenseRepo.Update(expense)
}

func (s *SplitwiseService) GetExpenseStatus(expenseID string) (*ExpenseStatus, error) {
	expense, err := s.expenseRepo.FindByID(expenseID)
	if err != nil {
		return nil, err
	}
	return &expense.Status, nil
}

func (s *SplitwiseService) GetGroupBalances(groupID string) (map[string]float64, error) {
	_, err := s.groupRepo.FindByID(groupID)
	if err != nil {
		return nil, err
	}

	expenses, err := s.expenseRepo.FindByGroupID(groupID)
	if err != nil {
		return nil, err
	}

	return s.balanceCalculationStrategy.CalculateBalances(expenses)
}

func (s *SplitwiseService) AddUserToGroup(groupID string, userID string) error {
	group, err := s.groupRepo.FindByID(groupID)
	if err != nil {
		return err
	}

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	for _, member := range group.Members {
		if member.ID == userID {
			return errors.New("user already in group")
		}
	}

	group.Members = append(group.Members, user)
	return s.groupRepo.Update(group)
}

func (s *SplitwiseService) RemoveUserFromGroup(groupID string, userID string) error {
	group, err := s.groupRepo.FindByID(groupID)
	if err != nil {
		return err
	}

	for i, member := range group.Members {
		if member.ID == userID {
			group.Members = append(group.Members[:i], group.Members[i+1:]...)
			return s.groupRepo.Update(group)
		}
	}

	return errors.New("user not found in group")
}

func (s *SplitwiseService) GetUserExpenses(userID string) ([]*Expense, error) {
	return s.expenseRepo.FindByUserID(userID)
}

func (s *SplitwiseService) GetGroupExpenses(groupID string) ([]*Expense, error) {
	return s.expenseRepo.FindByGroupID(groupID)
}
