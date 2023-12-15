package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Kaung Myat Han", City: "Yangon", Zipcode: "110011", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Han", City: "Yangon", Zipcode: "110011", DateOfBirth: "2000-01-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
