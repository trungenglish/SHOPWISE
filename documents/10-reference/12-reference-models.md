# Reference Models

```text
Conversation
Message
Recommendation
Comparison
Decision
Memory
Cart
Checkout
Order
Promotion
Inventory
```

---

Example

```go
type Recommendation struct{
    Products []Product
    Reasoning string
    Confidence float64
    Alternatives []Product
}
```
