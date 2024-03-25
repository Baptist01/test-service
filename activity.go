package app

type Account struct {
	Name     string
	Email    string
	Role     string
	Projects string
	Devices  string
}

func CreateMockAccounts() ([]Account, error) {
	account1 := Account{Name: "Baptist", Email: "Baptist.o@outlook.com", Role: "Admin", Projects: "Project1", Devices: "first device"}
	account2 := Account{Name: "Aaron", Email: "Aaron@outlook.com", Role: "Guest", Projects: "Project1", Devices: "laptop"}
	account3 := Account{Name: "Quinn", Email: "Quinn@outlook.com", Role: "Researcher", Projects: "Project2", Devices: "Desktop"}
	account4 := Account{Name: "alexandru", Email: "Alexandru@outlook.com", Role: "Researcher", Projects: "Project1", Devices: "Device1"}
	accounts := []Account{account1, account2, account3, account4}
	return accounts, nil
}
