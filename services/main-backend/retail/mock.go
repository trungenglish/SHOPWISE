package retail

type Product struct {
	ProductID string            `json:"productId"`
	Name      string            `json:"name"`
	Brand     string            `json:"brand"`
	Category  string            `json:"category"`
	Price     float64           `json:"price"` // VND
	Specs     map[string]string `json:"specs"`
	InStock   bool              `json:"inStock"`
	Images    []string          `json:"images"`
}

type MockRetailProvider struct {
	products map[string]Product
}

func NewMockRetailProvider() *MockRetailProvider {
	provider := &MockRetailProvider{
		products: make(map[string]Product),
	}
	// Seed some mock data
	provider.products["123"] = Product{
		ProductID: "123",
		Name:      "Gaming Laptop X",
		Brand:     "TechBrand",
		Category:  "Laptop",
		Price:     25000000,
		Specs:     map[string]string{"RAM": "16GB", "CPU": "Intel i7", "GPU": "RTX 4060"},
		InStock:   true,
		Images:    []string{"https://example.com/laptop-x.jpg"},
	}
	provider.products["456"] = Product{
		ProductID: "456",
		Name:      "Gaming Laptop Y",
		Brand:     "TechBrand",
		Category:  "Laptop",
		Price:     30000000,
		Specs:     map[string]string{"RAM": "32GB", "CPU": "Intel i9", "GPU": "RTX 4070"},
		InStock:   true,
		Images:    []string{"https://example.com/laptop-y.jpg"},
	}
	return provider
}

func (m *MockRetailProvider) Search(query string) []Product {
	// Simple mock implementation
	var results []Product
	for _, p := range m.products {
		// Just return all for now or do basic filtering
		results = append(results, p)
	}
	return results
}

func (m *MockRetailProvider) GetProducts(ids []string) []Product {
	var results []Product
	for _, id := range ids {
		if p, ok := m.products[id]; ok {
			results = append(results, p)
		}
	}
	return results
}
