Mình nghĩ đây là **phần quan trọng nhất sau AI Architecture**.

Nếu làm tốt **Tool SDK + Retail SDK**, thì:

* AI Runtime (Python) không cần biết Go.
* Go không cần biết LangGraph/PydanticAI.
* Retail Connector không cần biết AI.
* Thêm Phong Vũ, Shopee, Tiki, Lazada chỉ cần implement interface.
* Sau này có thể mở SDK cho đối tác thứ ba.

Theo mình, đây nên là **Protocol chứ không chỉ là SDK**.

---

# Cấu trúc

```text
docs/13-sdk/

README.md

01-tool-sdk/
    overview.md
    lifecycle.md
    contracts.md
    registry.md
    permissions.md
    errors.md
    observability.md
    versioning.md

02-retail-sdk/
    overview.md
    provider.md
    catalog.md
    inventory.md
    promotion.md
    commerce.md
    order.md
    capabilities.md
    normalization.md

03-reference/
    go/
    python/
    typescript/
```

---

# Tổng quan

```text
                AI Runtime
                     │
         Tool Protocol (JSON)
                     │
──────────────────────────────────

             Go Backend

──────────────────────────────────
                     │
         Retail SDK Interface
                     │
──────────────────────────────────

        PhongVu Provider

        Shopee Provider

        Lazada Provider

        CellphoneS Provider

──────────────────────────────────
                     │
              Retail APIs
```

Toàn bộ hệ thống chỉ có **hai contract**.

* Tool Contract
* Retail Contract

---

# Phần I

# Tool SDK

## Triết lý

```
LLM

↓

Planning

↓

Tool Call

↓

Go Backend

↓

Observation

↓

Reasoning
```

LLM **không gọi REST API**.

LLM **không query Database**.

LLM chỉ gọi Tool.

---

# Tool Contract

```typescript
interface Tool {

    metadata: ToolMetadata

    input: JSONSchema

    output: JSONSchema

    execute()

}
```

---

## Tool Metadata

```typescript
interface ToolMetadata {

    id:string

    name:string

    description:string

    category:string

    version:string

    timeout:number

    permissions:string[]

}
```

Ví dụ

```json
{
  "id":"catalog.search",

  "version":"1",

  "category":"catalog"
}
```

---

# Tool Categories

```text
Catalog

Inventory

Promotion

Comparison

Recommendation

Commerce

Order

Memory

Knowledge

Analytics
```

---

# Tool Lifecycle

```text
Discover

↓

Validate Input

↓

Authorize

↓

Execute

↓

Normalize

↓

Observe

↓

Return
```

---

# Execute Contract

Go

```go
type Tool interface {

    Metadata() ToolMetadata

    InputSchema() any

    OutputSchema() any

    Execute(ctx context.Context, input any) (any,error)

}
```

---

Python

```python
class Tool:

    metadata: ToolMetadata

    async def execute(self,input):

        ...
```

---

TypeScript

```ts
interface Tool {

    execute(input:any):Promise<any>

}
```

---

# Tool Registry

Không hardcode.

```go
registry.Register(

    SearchProductsTool{}

)
```

Registry

```text
catalog.search

catalog.compare

catalog.detail

inventory.check

promotion.list

cart.add

checkout.prepare

checkout.submit

memory.read

memory.write

recommend.generate
```

---

# Tool Discovery

AI khởi động

↓

Backend

↓

```json
[
 "catalog.search",

 "cart.add",

 "checkout.prepare"
]
```

LLM biết tool nào tồn tại.

---

# Permission Model

```text
Public

↓

Authenticated

↓

Customer

↓

Admin

↓

System
```

Ví dụ

```text
catalog.search

↓

Public

----------------

checkout.submit

↓

Customer

----------------

promotion.sync

↓

Admin
```

---

# Error Contract

```json
{
  "code":"PRODUCT_NOT_FOUND",

  "message":"...",

  "retryable":false
}
```

---

# Observation

Quan trọng.

LLM không nhận raw response.

Nhận

```json
{
 "summary":"3 products found",

 "data":{

 },

 "confidence":0.98
}
```

Observation dễ reasoning hơn.

---

# Telemetry

Mỗi Tool

```
duration

↓

input_size

↓

output_size

↓

errors

↓

retry

↓

cost
```

---

# Phần II

# Retail SDK

Đây là SDK quan trọng hơn.

---

## Triết lý

Không có

```text
PhongVuAPI
```

mà có

```text
Retail Provider
```

---

# Provider

```go
type Provider interface {

    Catalog()

    Inventory()

    Promotion()

    Commerce()

    Orders()

}
```

---

# Capability

Không phải retailer nào cũng có.

```go
type Capability struct{

SupportsCheckout bool

SupportsCoupon bool

SupportsRealtimeInventory bool

SupportsGuestCheckout bool

}
```

Ví dụ

Shopee

```text
Checkout

YES

----------------

Realtime Inventory

NO
```

Phong Vũ

```text
Checkout

YES

Realtime Inventory

YES
```

---

# Catalog

```go
type CatalogProvider interface {

Search()

Get()

Compare()

Categories()

Brands()

}
```

---

# Inventory

```go
type InventoryProvider interface{

Availability()

Warehouse()

DeliveryEstimate()

}
```

---

# Promotion

```go
type PromotionProvider interface{

Campaigns()

Coupons()

Bundles()

}
```

---

# Commerce

Quan trọng.

```go
type CommerceProvider interface{

CreateCart()

AddItem()

RemoveItem()

UpdateQuantity()

Checkout()

}
```

---

# Order

```go
type OrderProvider interface{

Status()

Tracking()

Warranty()

Cancel()

}
```

---

# Normalize

Retail API

↓

Provider

↓

Canonical Schema

↓

Go

↓

Tool

↓

Observation

↓

LLM

---

Không retailer nào trả cùng schema.

Ví dụ

Shopee

```json
{
 "item_name":"Laptop"
}
```

Phong Vũ

```json
{
 "name":"Laptop"
}
```

Canonical

```json
{
 "name":"Laptop"
}
```

LLM chỉ thấy

Canonical Schema.

---

# Capability Negotiation

Ví dụ

AI

↓

Checkout()

↓

Provider

↓

```text
supports_checkout?
```

↓

YES

↓

Checkout

---

Nếu

NO

↓

Generate Link

↓

Open Retail App

---

# Retry Strategy

```text
Network

↓

Retry

----------------

Authentication

↓

Refresh Token

----------------

Validation

↓

Fail Fast
```

---

# Event Model

```text
InventoryChanged

PromotionUpdated

OrderCreated

OrderCompleted

PriceChanged
```

Backend subscribe.

AI không cần biết.

---

# Provider Registry

```go
Register(

PhongVuProvider{}

)

Register(

ShopeeProvider{}

)

Register(

LazadaProvider{}

)
```

---

# Adapter Pattern

```text
Retail API

↓

Provider Adapter

↓

Retail SDK

↓

Commerce Service

↓

Tool SDK

↓

AI Runtime
```

---

# Kiến trúc cuối cùng

```text
                    Python AI Runtime
                            │
                  Planning / Reasoning
                            │
                    Tool SDK Contract
                            │
────────────────────────────────────────────────────────
                    Go Business Backend
                            │
                     Tool Registry
                            │
               Catalog / Commerce / Memory
                            │
                    Retail SDK Contract
                            │
────────────────────────────────────────────────────────
         PhongVu     Shopee     Lazada     CellphoneS
             │           │           │            │
             └───────────┼───────────┴────────────┘
                         ▼
                    Retail APIs
```

# Mình đề xuất nâng cấp thêm: Tool Protocol + Retail SDK + MCP

Nếu mục tiêu của SHOPWISE là **Agent-native**, mình sẽ không dừng ở SDK mà nâng thành **giao thức chuẩn**.

## Tool Protocol

Tool không nên là interface Go thuần mà là một **Protocol** độc lập ngôn ngữ:

```json
{
  "tool": "catalog.search",
  "version": "1.0",
  "input_schema": { ... },
  "output_schema": { ... },
  "capabilities": [
    "streaming",
    "pagination"
  ]
}
```

Điều này cho phép AI Runtime Python, Go Backend, hoặc thậm chí một dịch vụ Node.js đều hiểu cùng một hợp đồng.

## Retail Capability Manifest

Mỗi Retail Provider nên công bố một manifest:

```yaml
provider: phongvu
version: 1.0

capabilities:
  catalog: true
  inventory: true
  promotions: true
  cart: true
  checkout: true
  guest_checkout: false
  realtime_inventory: true
```

Agent có thể đọc manifest để quyết định workflow thay vì hard-code logic theo từng nhà bán lẻ.

## MCP Bridge

Lớp cuối cùng là **MCP Server**. Tool Registry có thể được expose qua MCP:

```text
AI Runtime
      │
Tool Registry
      │
 MCP Bridge
      │
Claude / Codex / Cursor / OpenAI
```

Khi đó, toàn bộ Tool SDK của SHOPWISE vừa phục vụ agent nội bộ, vừa có thể được các AI client khác sử dụng mà không cần viết thêm API chuyên biệt. Đây là hướng mở rộng rất phù hợp nếu sau này SHOPWISE muốn trở thành một nền tảng AI Commerce thay vì chỉ là một ứng dụng đơn lẻ.
