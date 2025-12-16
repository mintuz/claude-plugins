# Component Patterns

## Colocation Principle

Keep components, functions, styles, and state as close as possible to where they're used:

```
// ✅ Good - colocated
src/features/checkout/
├── components/
│   ├── checkout-form.tsx
│   ├── checkout-form.test.tsx      # Test next to component
│   └── use-checkout-form.ts        # Hook used only by this component

// ❌ Bad - scattered
src/
├── components/checkout-form.tsx
├── hooks/use-checkout-form.ts      # Far from where it's used
├── tests/checkout-form.test.tsx    # Even further away
```

## No Nested Render Functions

Extract UI units into separate components:

```typescript
// ❌ Bad - nested render function
const UserList = ({ users }: Props) => {
  const renderUser = (user: User) => (
    <div className="user-card">
      <Avatar src={user.avatar} />
      <span>{user.name}</span>
    </div>
  );

  return <div>{users.map(renderUser)}</div>;
};

// ✅ Good - extracted component
const UserCard = ({ user }: { user: User }) => (
  <div className="user-card">
    <Avatar src={user.avatar} />
    <span>{user.name}</span>
  </div>
);

const UserList = ({ users }: Props) => (
  <div>
    {users.map((user) => (
      <UserCard key={user.id} user={user} />
    ))}
  </div>
);
```

## Composition Over Props

Use children and slots instead of excessive props:

```typescript
// ❌ Bad - prop drilling, inflexible
type CardProps = {
  title: string;
  subtitle: string;
  icon: ReactNode;
  actions: ReactNode;
  footer: ReactNode;
  // ... props keep growing
};

// ✅ Good - composition
type CardProps = {
  children: ReactNode;
};

const Card = ({ children }: CardProps) => (
  <div className="card">{children}</div>
);

const CardHeader = ({ children }: { children: ReactNode }) => (
  <div className="card-header">{children}</div>
);

const CardBody = ({ children }: { children: ReactNode }) => (
  <div className="card-body">{children}</div>
);

// Usage - flexible composition
<Card>
  <CardHeader>
    <Icon name="user" />
    <h2>User Profile</h2>
  </CardHeader>
  <CardBody>
    <UserDetails user={user} />
  </CardBody>
</Card>;
```

## Wrap Third-Party Components

Insulate your application from dependency changes:

```typescript
// ✅ Good - wrapped third-party component
// src/components/ui/date-picker.tsx
import { DatePicker as ThirdPartyDatePicker } from "some-library";

type DatePickerProps = {
  value: Date | null;
  onChange: (date: Date | null) => void;
  minDate?: Date;
  maxDate?: Date;
};

export const DatePicker = ({
  value,
  onChange,
  minDate,
  maxDate,
}: DatePickerProps) => {
  return (
    <ThirdPartyDatePicker
      selected={value}
      onSelect={onChange}
      minDate={minDate}
      maxDate={maxDate}
    />
  );
};

// If library changes, only update this file
```
