---
description: Generate a Mermaid diagram from code files and directories to visualize architecture, data flow, or component relationships.
---

# Mermaid Diagram Generator

Generate a Mermaid diagram based on supplied code files and directories to visualize system architecture, component relationships, data flow, or other structural aspects of the codebase.

## Parameters

- `$ARGUMENTS` - One or more paths to files or directories to analyze. Can include:
  - Individual files (e.g., `./src/auth/login.ts`)
  - Directories (e.g., `./src/components/`)
  - Glob patterns (e.g., `./src/**/*.ts`)
  - Multiple paths separated by spaces

**Examples:**
- `/project:mermaid ./src/services/`
- `/project:mermaid ./src/api/routes.ts ./src/middleware/`
- `/project:mermaid ./src/features/auth/ flowchart`
- `/project:mermaid ./src/ --type=sequence --focus=api-calls`

## Process

### 1. Analyze the Codebase

First, gather information about the specified files and directories:

1. **List all relevant files**
   - Use Glob to find files matching the specified paths
   - Filter by relevant file extensions (.ts, .tsx, .js, .jsx, .py, etc.)

2. **Read and parse the code**
   - Identify exports, imports, and dependencies
   - Extract class definitions, functions, and interfaces
   - Map relationships between modules
   - Identify data flow patterns

3. **Identify key architectural elements**
   - Entry points and main modules
   - Service layers and their interactions
   - Data models and their relationships
   - External integrations and APIs
   - Event flows and async patterns

### 2. Determine Diagram Type

Based on the code structure and user intent, select the most appropriate diagram type:

| Diagram Type | Best For |
|-------------|----------|
| `flowchart` | General architecture, module relationships, process flows |
| `sequenceDiagram` | API calls, request/response flows, async operations |
| `classDiagram` | Object-oriented code, type hierarchies, data models |
| `erDiagram` | Database schemas, entity relationships |
| `stateDiagram` | State machines, component lifecycle, workflow states |
| `graph` | Dependency trees, import/export relationships |

If the user specifies a diagram type, use that. Otherwise, infer the best type from the code structure.

### 3. Build the Diagram

Construct the Mermaid diagram with these principles:

- **Clarity over completeness** - Focus on the most important relationships; don't include every function
- **Meaningful groupings** - Use subgraphs to group related components
- **Consistent naming** - Use the actual names from the codebase
- **Appropriate detail level** - Match detail to the scope of files provided
- **Direction** - Choose flow direction (TB, LR, etc.) that best represents the architecture

### 4. Add Annotations

Enhance the diagram with:
- Brief labels on relationships describing the interaction
- Notes for complex or non-obvious connections
- Color coding for different types of components (if helpful)
- Links to actual files where appropriate

## Output Format

Provide the output in this structure:

```markdown
## Diagram Overview

[1-2 sentences describing what the diagram represents and the scope of analysis]

## Files Analyzed

- `path/to/file1.ts` - [brief description]
- `path/to/file2.ts` - [brief description]

## Mermaid Diagram

\`\`\`mermaid
[Generated diagram code]
\`\`\`

## Diagram Legend

| Symbol/Color | Meaning |
|-------------|---------|
| [element] | [description] |

## Key Relationships

1. **[Relationship Name]** - [explanation of this connection and its significance]
2. **[Relationship Name]** - [explanation]

## Notes

- [Any assumptions made during analysis]
- [Components intentionally omitted for clarity]
- [Suggestions for additional diagrams that might be useful]
```

## Diagram-Specific Guidelines

### For Flowcharts
```mermaid
flowchart TB
    subgraph Layer["Layer Name"]
        A[Component A] --> B[Component B]
    end
    A -->|"action"| C[Component C]
```
- Use subgraphs to represent layers or domains
- Show data flow direction with arrows
- Label edges with the type of interaction

### For Sequence Diagrams
```mermaid
sequenceDiagram
    participant C as Client
    participant S as Server
    participant D as Database
    C->>S: Request
    S->>D: Query
    D-->>S: Result
    S-->>C: Response
```
- Focus on key interactions, not every function call
- Show async operations with appropriate arrow types
- Include error paths if significant

### For Class Diagrams
```mermaid
classDiagram
    class ClassName {
        +property: Type
        +method(): ReturnType
    }
    ClassName <|-- SubClass : extends
    ClassName --> OtherClass : uses
```
- Include public interfaces
- Show inheritance and composition relationships
- Use appropriate relationship symbols

### For ER Diagrams
```mermaid
erDiagram
    USER ||--o{ ORDER : places
    ORDER ||--|{ LINE-ITEM : contains
```
- Focus on entities and their relationships
- Use proper cardinality notation
- Include key attributes

## Special Instructions

- **Be pragmatic about scope** - For large directories, focus on the main architectural patterns rather than every file
- **Infer intent** - If the user provides a specific subset of files, assume they want to understand that specific subsystem
- **Handle ambiguity** - If multiple diagram types would be appropriate, either ask the user or provide the most generally useful one with a note about alternatives
- **Keep diagrams renderable** - Mermaid has syntax limitations; ensure the output is valid Mermaid syntax
- **Consider the audience** - Diagrams should be understandable by someone unfamiliar with the codebase

## Error Handling

| Situation | Response |
|-----------|----------|
| No files found at specified paths | List what was searched and suggest corrections |
| Files contain no analyzable structure | Explain why and suggest what files might be more useful |
| Too many files for meaningful single diagram | Offer to create multiple focused diagrams or ask user to narrow scope |
| Mixed languages/frameworks | Note this and focus on the dominant pattern or ask for clarification |
| Circular dependencies detected | Include them in the diagram and highlight as a potential concern |
