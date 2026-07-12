# UI Contracts: Checkout Readiness

## Overview

The AI Runtime generates these Dynamic UI schemas. The React Frontend interprets them and renders the corresponding UI components.

## 1. CheckoutSummary
Displays the locked-in product details, final price (VND), and Readiness Status.

```json
{
  "type": "CheckoutSummary",
  "props": {
    "status": "READY", // READY, MISSING_INFORMATION, OUT_OF_STOCK, etc.
    "product": {
      "id": "prod_123",
      "name": "Sony WH-1000XM5",
      "imageUrl": "https://...",
      "finalPriceVnd": 8490000
    },
    "checklist": [
      { "label": "In Stock", "passed": true },
      { "label": "Price Confirmed", "passed": true },
      { "label": "Customer Information", "passed": true }
    ]
  }
}
```

## 2. CustomerInformationForm
Secure form for collecting PII. Bypasses AI conversation history.

```json
{
  "type": "CustomerInformationForm",
  "props": {
    "requiredFields": ["name", "phoneNumber", "deliveryPreference"],
    "submitEndpoint": "/api/checkout-readiness/customer-info",
    "errorState": {
      "displayType": "inline",
      "fallbackMessage": "Failed to securely save your information. Please try again."
    }
  }
}
```

## 3. PromotionSummary
Highlights active promotions and includes the live countdown.

```json
{
  "type": "PromotionSummary",
  "props": {
    "promotionId": "promo_flash",
    "description": "Flash Sale: Save 500,000 VND",
    "discountAmountVnd": 500000,
    "expiresAt": "2026-07-12T15:30:00Z" // Absolute timestamp for countdown
  }
}
```

## 4. StorePicker
Facilitates hybrid geolocation for Store Pickup.

```json
{
  "type": "StorePicker",
  "props": {
    "allowGeolocation": true,
    "promptLocationConsent": {
      "title": "Use your location",
      "description": "We need your location to find the closest stores for pickup."
    },
    "searchEndpoint": "/api/stores/search",
    "nearbyEndpoint": "/api/stores/nearby"
  }
}
```

## 5. ActionBar
Contextual buttons for the user to proceed or take corrective action.

```json
{
  "type": "ActionBar",
  "props": {
    "actions": [
      {
        "id": "action_continue_checkout",
        "label": "Continue to Retail Checkout",
        "style": "primary",
        "disabled": false,
        "event": { "type": "CONTINUE_RETAIL_CHECKOUT" }
      },
      {
        "id": "action_save_session",
        "label": "Resume Later",
        "style": "secondary",
        "disabled": false,
        "event": { "type": "SAVE_SESSION" }
      }
    ]
  }
}
```

## 6. ErrorStateUI
Displays a structured error message when backend validation fails or times out.

```json
{
  "type": "ErrorStateUI",
  "props": {
    "errorCode": "TIMEOUT_ERROR",
    "message": "We couldn't reach the retailer to verify your checkout. Please try again.",
    "retryEndpoint": "/api/checkout-readiness/validate"
  }
}
```
