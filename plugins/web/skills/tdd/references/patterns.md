# TDD Testing Patterns and Examples

## Behavior-Focused Testing

### Test Behavior, Not Implementation

Tests should verify WHAT the code does, not HOW it does it.

```typescript
// ✅ GOOD - Tests business behavior
it("should reject payments with negative amounts", () => {
  const payment = getMockPayment({ amount: -100 });
  const result = processPayment(payment);

  expect(result.success).toBe(false);
  expect(result.error.message).toBe("Invalid amount");
});

it("should apply free shipping for orders over £50", () => {
  const order = getMockOrder({ subtotal: 60, shippingCost: 5.99 });
  const result = processOrder(order);

  expect(result.shippingCost).toBe(0);
  expect(result.total).toBe(60);
});

// ❌ BAD - Tests implementation details
it("should call validatePaymentAmount", () => {
  const spy = jest.spyOn(validator, "validateAmount");
  processPayment(payment);

  expect(spy).toHaveBeenCalled(); // Who cares if it's called?
});
```

### Test Through Public APIs Only

```typescript
// ✅ GOOD - Uses public API
const result = orderProcessor.processOrder(order);
expect(result.status).toBe("completed");

// ❌ BAD - Accesses internals
expect(orderProcessor._internalState.validated).toBe(true);
```

## Quick Start Patterns

### Factory with Schema Validation

```typescript
const getMockUser = (overrides?: Partial<User>): User => {
  return UserSchema.parse({
    id: "user-123",
    email: "[email protected]",
    role: "user",
    createdAt: new Date("2024-01-01"),
    ...overrides,
  });
};
```

### Composing Factories

```typescript
const getMockAddress = (overrides?: Partial<Address>): Address => ({
  line1: "123 Test Street",
  city: "London",
  postcode: "SW1A 1AA",
  country: "UK",
  ...overrides,
});

const getMockCustomer = (overrides?: Partial<Customer>): Customer => ({
  id: "customer-123",
  email: "test@example.com",
  name: "Test User",
  address: getMockAddress(),
  ...overrides,
});

// Usage with nested overrides
const customer = getMockCustomer({
  email: "vip@example.com",
  address: getMockAddress({ city: "Manchester" }),
});
```

### Basic Factory Pattern

```typescript
// ✅ GOOD - Factory with overrides
const getMockPayment = (overrides?: Partial<Payment>): Payment => {
  return {
    id: "payment-123",
    amount: 100,
    currency: "GBP",
    cardId: "card_456",
    ...overrides,
  };
};

// Usage in tests
it("should reject payments with negative amounts", () => {
  const payment = getMockPayment({ amount: -100 });
  const result = processPayment(payment);

  expect(result.success).toBe(false);
});

// ❌ BAD - let and beforeEach
let payment: Payment;

beforeEach(() => {
  payment = { id: "123", amount: 100, currency: "GBP" };
});

it("should process payment", () => {
  // Where did payment come from? What's its state?
});
```

## Test Naming and Organization

### Descriptive Test Names

```typescript
// ✅ GOOD - Documents behavior
"should reject payments with negative amounts"
"should apply free shipping for orders over £50"
"should charge shipping for orders exactly at £50"

// ❌ BAD - Describes implementation
"should call validateAmount method"
"should set isValid to true"
```

### Test Organization by Behavior

```typescript
// ✅ GOOD - Organized by business behavior
describe("Order processing", () => {
  describe("shipping calculations", () => {
    it("should charge standard shipping for orders under £50", () => {});
    it("should apply free shipping for orders over £50", () => {});
    it("should charge shipping for orders exactly at £50", () => {});
  });

  describe("discount application", () => {
    it("should apply 10% discount for orders over £100", () => {});
    it("should stack discounts with free shipping", () => {});
  });
});
```
