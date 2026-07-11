# Node System

SAUP defines specific Node `type` strings. These act as the vocabulary for the UI Composer.

## Layout Nodes

Structural containers that define how children are positioned.

- `workspace`
- `page`
- `sidebar`
- `panel`
- `tabs`
- `split`
- `stack`
- `grid`
- `row`
- `column`
- `container`

## Decision Nodes

High-level contextual blocks that present the AI's reasoning.

- `recommendation.panel`
- `comparison.panel`
- `reasoning.panel`
- `decision.summary`
- `decision.timeline`
- `decision.memory`

## Commerce Nodes

Transactional and operational elements.

- `cart.panel`
- `checkout.panel`
- `promotion.banner`
- `shipping.info`
- `inventory.badge`

## Product Nodes

Merchandising and product display.

- `product.card`
- `product.grid`
- `product.carousel`
- `product.list`
- `product.gallery`

## Chat Nodes

Conversational interface elements.

- `chat.window`
- `chat.message`
- `chat.input`
- `chat.typing`

## Primitive Nodes

Basic building blocks for data display.

- `text`
- `button`
- `image`
- `icon`
- `badge`
- `divider`
- `progress`

All components map to strict native implementations via the `Component Registry`.
