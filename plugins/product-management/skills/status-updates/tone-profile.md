# Tone of Voice Profile

Detailed reference for replicating the warm, organized update style.

## Voice Characteristics

| Dimension | Style |
|-----------|-------|
| Formality | Professional-casual blend; technical when needed, approachable always |
| Warmth | Generous with credit, inclusive language ("we", team mentions) |
| Pace | Brisk; punchy sentences, scannable structure, no filler |
| Energy | High but controlled; enthusiasm channeled through structure |

## Signature Patterns

### Emoji as Section Anchors

Each topic gets a themed emoji prefix. Never decorative—always organizational.

```
:art: Design System & Templates
:rocket: Launch Updates
:cookie: Refactoring
:books: Documentation
:globe_with_meridians: Infrastructure
```

### Good First, Challenges Second

Bad news is sandwiched and contextualized, never buried.

> "Great progress and some less-than-ideal timeline changes. We'll cover the good first, as it's very easy to lose sight of just how much work is being shipped every day"

### Named Gratitude

Team members are called out by name at section ends, making appreciation specific rather than generic.

> "Big shout out to the team for working through some tough challenges: @person1, @person2, @person3..."

### Seasonal/Cultural Hooks

Updates often open with a themed invitation that humanizes the content.

> "Is it appropriate to drink mulled wine whilst at work and this close to Christmas?? If so, grab yourself a little Glühwein and settle in for a TEAM_NAME update"

### Sticky Analogies

Technical progress explained through memorable comparisons.

> "Remember that shopping cart analogy I've been whipping out at every opportunity? Well we're looking more 'race car' every day."

### Inline Evidence

Links to production, Storybook, docs woven naturally into claims—showing, not just telling.

> "shipped to production, to the X page ([our fastest growing page](link))"

### Honest Caveats

Footnotes or asterisks used for transparency without cluttering the main flow.

> "* This _is_ production, but this page does not resolve through the monzo.com domain; only we can see it"

## Replication Techniques

### Delivering Unwelcome News

Acknowledge the feeling directly, then pivot to context with bullet points that reframe.

**Pattern:** "[Bad news], which obviously [acknowledge feeling]. There is [nuance/context] though:"

**Example:** "It was necessary to revise our estimated delivery date, which obviously no one wants. There is a little more nuance to what exactly is being worked on, though:"

### Connecting Tactical to Strategic

Connect work to benefit in a single sentence.

**Pattern:** [What we did] + [immediate impact] + [future payoff framed as simplification]

**Example:** "This puts the foundational work and a real-world component in front of thousands of customers and makes the process of an identical FAQ section for the new product pages a question of configuration, not engineering"

### Crediting Innovation

Use a light touch of playfulness paired with a named individual.

**Example:** "Using some cutting edge AI wizardry from @person1 we've fast tracked the X page"

### Building Momentum

End sections with forward-looking statements that create anticipation.

**Examples:**
- "will kick off in January"
- "this month"
- "in the new year once we've validated the design"

## Extended Do's and Don'ts

### Do

- Open with a friendly hook or greeting before diving in
- Use emoji to create visual hierarchy, one per section header
- Credit individuals by name—specificity shows genuine appreciation
- Link to evidence (production, docs, Storybook) to back up claims
- Acknowledge challenges directly, then reframe with context
- Use bullet points and numbered lists liberally—optimize for scanning
- Use backticks for component/technical names (`Button`, `Dialog`)
- End sections with momentum ("will kick off in January", "this month")
- Use italics for soft emphasis or asides
- Include action items with owners when relevant

### Don't

- Bury bad news or avoid it—address it head-on with balance
- Use emoji as decoration; every emoji should serve a structural purpose
- Give vague team thanks ("thanks everyone")—name people
- Over-explain technical concepts; trust the audience's competence
- Use excessive exclamation marks; let structure convey enthusiasm
- Write dense paragraphs—break into scannable chunks
- Make claims without evidence or links when available
- Use overly formal or stiff language
- Forget to close the loop on previous updates
