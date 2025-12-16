# Common TDD Violations and Code Smells

## Critical Violations (Must Fix)

| Violation | Problem | Fix |
| --- | --- | --- |
| Production code without failing test | Core TDD principle broken | Delete code, write test first |
| Multiple tests before making first pass | Batching, not TDD | Focus on one test at a time |
| More code than needed | Over-engineering | Remove excess, only pass current test |
| Implementation-focused tests | Brittle, don't verify behavior | Rewrite to test outcomes |

### Production Code Without Failing Test

**The core TDD principle:** Every single line of production code must be written in response to a failing test.

```typescript
// ❌ WRONG - Writing production code first
const calculateDiscount = (order: Order): number => {
  if (order.total > 100) {
    return order.total * 0.1;
  }
  return 0;
};

// ✅ CORRECT - Write the test first
it("should apply 10% discount for orders over £100", () => {
  const order = getMockOrder({ total: 150 });
  const discount = calculateDiscount(order);
  expect(discount).toBe(15);
});

// Then write minimal implementation to pass
const calculateDiscount = (order: Order): number => {
  if (order.total > 100) {
    return order.total * 0.1;
  }
  return 0;
};
```

## High Priority Issues

| Issue | Problem | Fix |
| --- | --- | --- |
| Using `let`/`beforeEach` | Shared mutable state | Use factory functions |
| Testing private methods | Coupling to implementation | Test through public API |
| `any` types in tests | Type safety disabled | Use proper types |
| Missing edge case tests | Incomplete coverage | Add boundary tests |
| Vague test names | Poor documentation | Use behavior-focused names |

### Testing Implementation Details

```typescript
// ❌ WRONG - Tests implementation details
it("should call validatePaymentAmount", () => {
  const spy = jest.spyOn(validator, "validateAmount");
  processPayment(payment);

  expect(spy).toHaveBeenCalled(); // Who cares if it's called?
});

it("should use the PaymentGateway class", () => {
  // Testing internal wiring, not behavior
});

// ✅ CORRECT - Tests business behavior
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
```

### Testing Through Public APIs Only

Tests should only interact with the public interface. Internal methods and state are invisible to tests.

```typescript
// ✅ GOOD - Uses public API
const result = orderProcessor.processOrder(order);
expect(result.status).toBe("completed");

// ❌ BAD - Accesses internals
expect(orderProcessor._internalState.validated).toBe(true);
expect(orderProcessor.privateValidate).toHaveBeenCalled();
```

## Style Issues

| Issue | Problem | Fix |
| --- | --- | --- |
| Large test files | Hard to navigate | Organize by behavior |
| Test duplication | Maintenance burden | Extract shared factories |
| Magic values in tests | Unclear intent | Use named constants or clear values |

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

## Behavior-Focused Testing Principles

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
