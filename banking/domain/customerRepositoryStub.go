package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Random", City: "cityofgold", Zipcode: "9999", DateofBirth: "forever", Status: "1"},
		{Id: "1002", Name: "Person", City: "cityofgold", Zipcode: "9999", DateofBirth: "forever", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
