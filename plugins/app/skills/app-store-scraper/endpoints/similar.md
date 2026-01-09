# Similar Apps

Find apps similar to a given app (based on "Customers Also Bought").

## Endpoint

```bash
https://apps.apple.com/{country}/app/id{id}
```

**Note:** This requires web scraping, not a structured API.

## How It Works

1. Fetch the App Store page HTML
2. Extract app IDs from "Customers Also Bought" section
3. Use Lookup API to get full details

## Examples

### Extract Similar App IDs

```bash
# Fetch app page and extract similar app IDs
curl -s "https://apps.apple.com/us/app/id553834731" | \
  grep -oE 'id[0-9]{9,}' | \
  grep -oE '[0-9]{9,}' | \
  sort -u | \
  head -20
```

### Get Full Details for Similar Apps

```bash
# Get similar app IDs
SIMILAR_IDS=$(curl -s "https://apps.apple.com/us/app/id553834731" | \
  grep -oE 'id[0-9]{9,}' | \
  grep -oE '[0-9]{9,}' | \
  sort -u | \
  head -10 | \
  tr '\n' ',')

# Remove trailing comma
SIMILAR_IDS=${SIMILAR_IDS%,}

# Lookup details
curl -s "https://itunes.apple.com/lookup?id=${SIMILAR_IDS}&entity=software" | \
  jq '.results[] | {
    name: .trackName,
    id: .trackId,
    developer: .artistName,
    rating: .averageUserRating
  }'
```

### Alternative: Using HTML Parser

If you have `pup` or `htmlq` installed:

```bash
# Using pup
curl -s "https://apps.apple.com/us/app/id553834731" | \
  pup 'a[href*="/app/id"] attr{href}' | \
  grep -oE 'id[0-9]+' | \
  grep -oE '[0-9]+' | \
  sort -u
```

### Filter Out Original App

```bash
ORIGINAL_ID=553834731

curl -s "https://apps.apple.com/us/app/id${ORIGINAL_ID}" | \
  grep -oE 'id[0-9]{9,}' | \
  grep -oE '[0-9]{9,}' | \
  sort -u | \
  grep -v "^${ORIGINAL_ID}$" | \
  head -10
```

### Get Similar Apps with Ratings

```bash
# Extract IDs
IDS=$(curl -s "https://apps.apple.com/us/app/id553834731" | \
  grep -oE 'id[0-9]{9,}' | \
  grep -oE '[0-9]{9,}' | \
  sort -u | \
  head -10 | \
  tr '\n' ',')

# Get details and filter by rating
curl -s "https://itunes.apple.com/lookup?id=${IDS}&entity=software" | \
  jq '.results[] | select(.averageUserRating >= 4.0) | {
    name: .trackName,
    rating: .averageUserRating,
    category: .primaryGenreName
  }'
```

## Limitations

- Requires HTML parsing (not a structured API)
- Page structure may change without notice
- "Customers Also Bought" section may not always be present
- Results depend on Apple's recommendation algorithm
- Rate limiting may apply for frequent requests

## Alternative Approaches

### By Category

Find similar apps by browsing the same category:

```bash
# Get app's category
CATEGORY=$(curl -s "https://itunes.apple.com/lookup?id=553834731&entity=software" | \
  jq -r '.results[0].primaryGenreName')

# Search for apps in same category
curl -s "https://itunes.apple.com/search?term=${CATEGORY}&media=software&entity=software&limit=20" | \
  jq '.results[] | {name: .trackName, category: .primaryGenreName}'
```

### By Developer

Find other apps from the same developer:

```bash
DEV_ID=$(curl -s "https://itunes.apple.com/lookup?id=553834731" | \
  jq -r '.results[0].artistId')

curl -s "https://itunes.apple.com/lookup?id=${DEV_ID}&entity=software" | \
  jq '.results[] | select(.wrapperType == "software") | .trackName'
```

## Use Cases

- App recommendation systems
- Competitive analysis
- Market research
- Related app discovery
- Cross-promotion opportunities

## Best Practices

1. Cache results to minimize requests
2. Add delays between scraping requests
3. Handle missing "Similar Apps" section gracefully
4. Validate extracted IDs before lookup
5. Consider using official API endpoints when possible
