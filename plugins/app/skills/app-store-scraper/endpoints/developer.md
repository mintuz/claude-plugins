# Developer Apps

Retrieve all apps published by a specific developer.

## Endpoint

```bash
https://itunes.apple.com/lookup
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `id` | string | Yes | Developer/Artist ID |
| `country` | string | No | Market code (default: `us`) |
| `entity` | string | No | Set to `software` for apps |

## Examples

### Get All Apps by Developer

```bash
curl -s "https://itunes.apple.com/lookup?id=284882215&country=us&entity=software" | \
  jq '.results[] | select(.wrapperType == "software") | {
    name: .trackName,
    id: .trackId,
    bundleId: .bundleId,
    category: .primaryGenreName,
    rating: .averageUserRating
  }'
```

### Get Developer ID from an App

```bash
# First lookup an app to get developer ID
DEV_ID=$(curl -s "https://itunes.apple.com/lookup?bundleId=com.apple.Numbers" | \
  jq -r '.results[0].artistId')

echo "Developer ID: $DEV_ID"

# Then fetch all apps by that developer
curl -s "https://itunes.apple.com/lookup?id=${DEV_ID}&entity=software" | \
  jq '.results[] | select(.wrapperType == "software") | .trackName'
```

### Count Developer's Portfolio

```bash
curl -s "https://itunes.apple.com/lookup?id=284882215&entity=software" | \
  jq '[.results[] | select(.wrapperType == "software")] | length'
```

### Get Developer's Top-Rated Apps

```bash
curl -s "https://itunes.apple.com/lookup?id=284882215&entity=software" | \
  jq '.results[] | select(.wrapperType == "software") | {
    name: .trackName,
    rating: .averageUserRating,
    ratingCount: .userRatingCount
  } | select(.rating >= 4.0)' | \
  jq -s 'sort_by(.rating) | reverse'
```

### List Developer's Free vs Paid Apps

```bash
curl -s "https://itunes.apple.com/lookup?id=284882215&entity=software" | \
  jq '.results[] | select(.wrapperType == "software") | {
    name: .trackName,
    price: .price,
    type: (if .price == 0 then "Free" else "Paid" end)
  }'
```

## Response Structure

Returns the same structure as App Lookup, with multiple app results:

```json
{
  "resultCount": 15,
  "results": [
    {
      "wrapperType": "software",
      "artistId": 284882215,
      "artistName": "Developer Name",
      "trackId": 553834731,
      "trackName": "App Name",
      ...
    }
  ]
}
```

## Important Notes

- First result may contain developer info (not an app)
- Filter for `wrapperType == "software"` to get only apps
- Some developers may have apps in multiple countries
- Developer ID (artistId) is different from app ID (trackId)

## Finding Developer IDs

### Method 1: From an App

```bash
curl -s "https://itunes.apple.com/lookup?id=553834731" | \
  jq -r '.results[0] | "Developer: \(.artistName)\nID: \(.artistId)"'
```

### Method 2: From Bundle ID

```bash
curl -s "https://itunes.apple.com/lookup?bundleId=com.example.app" | \
  jq -r '.results[0].artistId'
```

### Method 3: From Search

```bash
curl -s "https://itunes.apple.com/search?term=developer%20name&media=software&entity=software&limit=1" | \
  jq -r '.results[0].artistId'
```

## Use Cases

- Analyze developer portfolios
- Track competitor app releases
- Monitor developer rating trends
- Build developer profile pages
- Discover related apps from same developer
