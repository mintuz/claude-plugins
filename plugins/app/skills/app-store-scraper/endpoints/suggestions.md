# Search Suggestions

Get autocomplete suggestions for search queries.

## Endpoint

```bash
https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `clientApplication` | string | Yes | Set to `Software` |
| `term` | string | Yes | Partial search query |

## Examples

### Basic Suggestions

```bash
curl -s "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints?clientApplication=Software&term=photo" | \
  jq '.hints[]'
```

### Format as List

```bash
curl -s "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints?clientApplication=Software&term=fitness" | \
  jq -r '.hints[] | "- \(.)"'
```

### Get Top 5 Suggestions

```bash
curl -s "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints?clientApplication=Software&term=game" | \
  jq -r '.hints[:5][]'
```

### Search Suggestions for Each Letter

```bash
for letter in p ph pho phot photo; do
  echo "Suggestions for: $letter"
  curl -s "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints?clientApplication=Software&term=$letter" | \
    jq -r '.hints[:3][]'
  echo ""
done
```

### Use Suggestions to Find Apps

```bash
# Get suggestion
SUGGESTION=$(curl -s "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints?clientApplication=Software&term=photo%20edit" | \
  jq -r '.hints[0]')

echo "Top suggestion: $SUGGESTION"

# Search using suggestion
curl -s "https://itunes.apple.com/search?term=$(echo $SUGGESTION | sed 's/ /%20/g')&media=software&entity=software&limit=5" | \
  jq '.results[] | {name: .trackName, developer: .artistName}'
```

## Response Structure

```json
{
  "hints": [
    "photo editor",
    "photo collage",
    "photo vault",
    "photo booth",
    "photography"
  ]
}
```

## Use Cases

- Build autocomplete search interfaces
- Discover popular search terms
- Improve search query accuracy
- Generate keyword ideas
- User experience enhancement

## Tips

- Suggestions are based on popular searches
- Results are language and region-specific
- Minimum 2-3 characters recommended for meaningful suggestions
- Suggestions update based on trending searches
- URL encode special characters in term parameter

## Example: Build Autocomplete UI

```bash
#!/bin/bash

get_suggestions() {
  local query="$1"
  curl -s "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints?clientApplication=Software&term=${query}" | \
    jq -r '.hints[]'
}

# Interactive search
read -p "Enter search term: " search_term
echo "Suggestions:"
get_suggestions "$search_term" | nl
```

## Combining with Search

```bash
# Get suggestions and search each
TERM="weather"

curl -s "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints?clientApplication=Software&term=$TERM" | \
  jq -r '.hints[:3][]' | \
  while read suggestion; do
    echo "=== Results for: $suggestion ==="
    curl -s "https://itunes.apple.com/search?term=$(echo $suggestion | sed 's/ /%20/g')&media=software&entity=software&limit=3" | \
      jq -r '.results[] | "  - \(.trackName) by \(.artistName)"'
    echo ""
  done
```
