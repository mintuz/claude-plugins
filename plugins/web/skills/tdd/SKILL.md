---
name: tdd-best-practices
description: Write tests and code using Test-Driven Development principles. Use this skill when the user asks to write tests, implement features, or develop code that should follow TDD practices. Ensures behavior-focused testing and proper Red-Green-Refactor workflow.
---

# TDD Best Practices Knowledge Base

This skill provides comprehensive knowledge of Test-Driven Development principles, behavior-focused testing, and common anti-patterns. Use this knowledge when writing tests, implementing features, or reviewing code for TDD compliance.

## Core Principle

**Every single line of production code must be written in response to a failing test.**

This is non-negotiable. If you're typing production code without a failing test demanding it, you're not doing TDD.

## The Sacred Cycle: Red → Green → Refactor

### 1. RED - Write a Failing Test

Write a test that describes the desired behavior. The test must fail because the behavior doesn't exist yet.

**Rules:**
- Start with the simplest behavior
- Test ONE thing at a time
- Focus on business behavior, not implementation
- Use descriptive test names that document intent
- Use factory functions for test data

### 2. GREEN - Minimal Implementation

Write the **minimum** code to make the test pass. Nothing more.

**Rules:**
- Only enough code to pass the current test
- Resist "just in case" logic
- No speculative features
- If writing more than needed, STOP and question why

### 3. REFACTOR - Assess and Improve

With tests green, assess whether refactoring would add value.

**Rules:**
- Commit working code FIRST
- External APIs stay unchanged
- All tests must still pass
- Commit refactoring separately
- Not all code needs refactoring - if clean, move on

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

it("should use the PaymentGateway class", () => {
  // Testing internal wiring, not behavior
});
```

### Test Through Public APIs Only

Tests should only interact with the public interface. Internal methods and state are invisible to tests.

```typescript
// ✅ GOOD - Uses public API
const result = orderProcessor.processOrder(order);
expect(result.status).toBe("completed");

// ❌ BAD - Accesses internals
expect(orderProcessor._internalState.validated).toBe(true);
expect(orderProcessor.privateValidate).toHaveBeenCalled();
```

### Descriptive Test Names

Test names should document business behavior, not implementation steps.

```typescript
// ✅ GOOD - Documents behavior
"should reject payments with negative amounts"
"should apply free shipping for orders over £50"
"should charge shipping for orders exactly at £50"
"should calculate tax based on shipping address"

// ❌ BAD - Describes implementation
"should call validateAmount method"
"should set isValid to true"
"should use the correct formula"
"should invoke the callback"
```

## Test Data Factories

### Factory Functions Over let/beforeEach

Use factory functions with optional overrides for test data. Never use `let` declarations or `beforeEach` for test setup.

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

const getMockOrder = (overrides?: Partial<Order>): Order => {
  return {
    id: "order-789",
    items: [getMockOrderItem()],
    subtotal: 50,
    shippingCost: 5.99,
    status: "pending",
    ...overrides,
  };
};

// Usage in tests
it("should reject payments with negative amounts", () => {
  const payment = getMockPayment({ amount: -100 });
  const result = processPayment(payment);

  expect(result.success).toBe(false);
});

it("should apply discount for large orders", () => {
  const order = getMockOrder({ subtotal: 200 });
  const result = processOrder(order);

  expect(result.discount).toBeGreaterThan(0);
});
```

```typescript
// ❌ BAD - let and beforeEach
let payment: Payment;
let order: Order;

beforeEach(() => {
  payment = { id: "123", amount: 100, currency: "GBP" };
  order = { id: "456", items: [], subtotal: 50 };
});

it("should process payment", () => {
  // Where did payment come from? What's its state?
  // Hard to trace, easy to have shared state bugs
});
```

**Why factories are better:**

| let/beforeEach | Factory Functions |
| --- | --- |
| Shared mutable state | Fresh data each test |
| Hard to trace data origin | Explicit data creation |
| Implicit test coupling | Isolated tests |
| Mutation bugs possible | Immutable by design |
| Harder to customize | Easy overrides |

### Composing Factories

Build complex test data by composing simpler factories:

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

const getMockOrderWithCustomer = (
  overrides?: Partial<Order & { customer?: Partial<Customer> }>
): Order => {
  const { customer: customerOverrides, ...orderOverrides } = overrides ?? {};
  return {
    ...getMockOrder(orderOverrides),
    customer: getMockCustomer(customerOverrides),
  };
};

// Usage
const order = getMockOrderWithCustomer({
  subtotal: 100,
  customer: { email: "vip@example.com" },
});
```

## TDD Workflow Examples

### Example 1: Adding Free Shipping Feature

**Step 1: RED - Write failing test for simplest behavior**

```typescript
it("should calculate total with shipping cost", () => {
  const order = getMockOrder({ subtotal: 30, shippingCost: 5.99 });

  const result = processOrder(order);

  expect(result.total).toBe(35.99);
  expect(result.shippingCost).toBe(5.99);
});
```

**Step 2: GREEN - Minimal implementation**

```typescript
const processOrder = (order: Order): ProcessedOrder => {
  return {
    ...order,
    total: order.subtotal + order.shippingCost,
  };
};
```

**Step 3: RED - Add test for free shipping behavior**

```typescript
it("should apply free shipping for orders over £50", () => {
  const order = getMockOrder({ subtotal: 60, shippingCost: 5.99 });

  const result = processOrder(order);

  expect(result.shippingCost).toBe(0);
  expect(result.total).toBe(60);
});
```

**Step 4: GREEN - Add conditional (now both paths tested)**

```typescript
const processOrder = (order: Order): ProcessedOrder => {
  const shippingCost = order.subtotal > 50 ? 0 : order.shippingCost;

  return {
    ...order,
    shippingCost,
    total: order.subtotal + shippingCost,
  };
};
```

**Step 5: RED - Add edge case test**

```typescript
it("should charge shipping for orders exactly at £50", () => {
  const order = getMockOrder({ subtotal: 50, shippingCost: 5.99 });

  const result = processOrder(order);

  expect(result.shippingCost).toBe(5.99);
  expect(result.total).toBe(55.99);
});
```

**Step 6: REFACTOR - Extract constant (if valuable)**

```typescript
const FREE_SHIPPING_THRESHOLD = 50;

const qualifiesForFreeShipping = (subtotal: number): boolean => {
  return subtotal > FREE_SHIPPING_THRESHOLD;
};

const processOrder = (order: Order): ProcessedOrder => {
  const shippingCost = qualifiesForFreeShipping(order.subtotal)
    ? 0
    : order.shippingCost;

  return {
    ...order,
    shippingCost,
    total: order.subtotal + shippingCost,
  };
};
```

### Example 2: Payment Validation

**RED → GREEN → RED → GREEN pattern:**

```typescript
// Test 1: RED
it("should process valid payments", () => {
  const payment = getMockPayment({ amount: 100 });
  const result = processPayment(payment);

  expect(result.success).toBe(true);
});

// GREEN: Minimal implementation
const processPayment = (payment: Payment): Result<Receipt> => {
  return { success: true, data: { id: "receipt-123" } };
};

// Test 2: RED
it("should reject payments with negative amounts", () => {
  const payment = getMockPayment({ amount: -100 });
  const result = processPayment(payment);

  expect(result.success).toBe(false);
  expect(result.error.message).toBe("Invalid amount");
});

// GREEN: Add validation
const processPayment = (payment: Payment): Result<Receipt> => {
  if (payment.amount < 0) {
    return { success: false, error: new Error("Invalid amount") };
  }
  return { success: true, data: { id: "receipt-123" } };
};

// Test 3: RED
it("should reject payments with zero amount", () => {
  const payment = getMockPayment({ amount: 0 });
  const result = processPayment(payment);

  expect(result.success).toBe(false);
});

// GREEN: Adjust condition
const processPayment = (payment: Payment): Result<Receipt> => {
  if (payment.amount <= 0) {
    return { success: false, error: new Error("Invalid amount") };
  }
  return { success: true, data: { id: "receipt-123" } };
};
```

## Refactoring Assessment

### When to Refactor

After tests are green, assess whether refactoring would add value:

| Signal | Refactoring Action |
| --- | --- |
| Magic numbers repeated | Extract named constants |
| Unclear names | Improve naming |
| Complex logic | Extract functions |
| Knowledge duplication | Create single source of truth |
| Nested structure | Use early returns |
| Long functions | Split into smaller functions |

### When NOT to Refactor

Not all code needs refactoring. If the code is already clean:

- Clear function names ✓
- No magic numbers ✓
- Simple structure ✓
- Self-documenting ✓

Then commit and move to the next test.

### Refactoring Rules

1. **Commit working code FIRST** - Never refactor uncommitted code
2. **Keep tests green** - All tests must pass throughout
3. **Preserve external API** - Don't change public interfaces
4. **Commit refactoring separately** - Clean git history
5. **Small steps** - Refactor incrementally

## Common TDD Violations

### Critical Violations

| Violation | Problem | Fix |
| --- | --- | --- |
| Production code without failing test | Core TDD principle broken | Delete code, write test first |
| Multiple tests before making first pass | Batching, not TDD | Focus on one test at a time |
| More code than needed | Over-engineering | Remove excess, only pass current test |
| Implementation-focused tests | Brittle, don't verify behavior | Rewrite to test outcomes |

### High Priority Issues

| Issue | Problem | Fix |
| --- | --- | --- |
| Using `let`/`beforeEach` | Shared mutable state | Use factory functions |
| Testing private methods | Coupling to implementation | Test through public API |
| `any` types in tests | Type safety disabled | Use proper types |
| Missing edge case tests | Incomplete coverage | Add boundary tests |
| Vague test names | Poor documentation | Use behavior-focused names |

### Style Issues

| Issue | Problem | Fix |
| --- | --- | --- |
| Large test files | Hard to navigate | Organize by behavior |
| Test duplication | Maintenance burden | Extract shared factories |
| Magic values in tests | Unclear intent | Use named constants or clear values |

## Test Organization

### Organize by Behavior, Not Implementation

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

// ❌ BAD - Organized by implementation
describe("OrderProcessor", () => {
  describe("calculateShipping method", () => {});
  describe("applyDiscount method", () => {});
  describe("validateOrder method", () => {});
});
```

### No 1:1 Test File to Implementation File Mapping

Tests should be organized by feature/behavior, not mirroring implementation structure:

```
// ❌ BAD - 1:1 mapping
src/
  payment-validator.ts
  payment-validator.test.ts
  payment-processor.ts
  payment-processor.test.ts

// ✅ GOOD - By feature/behavior
src/
  payment/
    payment-processor.ts      # Public API
    payment-validator.ts      # Internal (implementation detail)
    payment-processor.test.ts # Tests ALL payment behavior
```

The validator is an implementation detail. Its logic is fully covered by testing the processor's behavior.

## 100% Coverage Through Behavior

Achieve complete coverage by testing all business behaviors, not by targeting implementation:

```typescript
// payment-validator.ts (implementation detail - no direct tests)
export const validateAmount = (amount: number): boolean => {
  return amount > 0 && amount <= 10000;
};

// payment-processor.ts (public API)
export const processPayment = (payment: Payment): Result<Receipt> => {
  if (!validateAmount(payment.amount)) {
    return { success: false, error: new PaymentError("Invalid amount") };
  }
  // ... process payment
};

// payment-processor.test.ts - tests achieve 100% coverage of validator
// without directly testing validateAmount

it("should reject payments with negative amounts", () => {
  const payment = getMockPayment({ amount: -100 });
  const result = processPayment(payment);
  expect(result.success).toBe(false);
});

it("should reject payments with zero amount", () => {
  const payment = getMockPayment({ amount: 0 });
  const result = processPayment(payment);
  expect(result.success).toBe(false);
});

it("should reject payments exceeding maximum", () => {
  const payment = getMockPayment({ amount: 10001 });
  const result = processPayment(payment);
  expect(result.success).toBe(false);
});

it("should process valid payment amounts", () => {
  const payment = getMockPayment({ amount: 100 });
  const result = processPayment(payment);
  expect(result.success).toBe(true);
});
```

## Quality Gates

Before committing, verify:

- ✅ All production code has a test that demanded it
- ✅ Tests verify behavior, not implementation
- ✅ Implementation is minimal (only what's needed)
- ✅ Refactoring assessment completed
- ✅ All tests pass
- ✅ Factory functions used (no `let`/`beforeEach`)
- ✅ Test names describe business behavior
- ✅ Edge cases covered

## Quick Reference: Decision Trees

### Should I write this code?

```
Is there a failing test demanding this code?
├── Yes → Write minimal code to pass
└── No → Write the failing test first
```

### Is my test good?

```
Does the test verify a business outcome?
├── Yes → Does it use the public API only?
│   ├── Yes → Does it use factory functions?
│   │   ├── Yes → Good test ✓
│   │   └── No → Refactor to use factories
│   └── No → Rewrite to avoid internals
└── No → Rewrite to focus on behavior
```

### Should I refactor?

```
Are all tests green?
├── Yes → Is the code already clean?
│   ├── Yes → Commit and move on
│   └── No → Commit first, then refactor
└── No → Make tests pass first
```

### How much code should I write?

```
Does this code make the current failing test pass?
├── Yes → Is there any code that could be removed
│         and tests still pass?
│   ├── Yes → Remove it
│   └── No → Done, commit
└── No → Keep writing minimal code
```
