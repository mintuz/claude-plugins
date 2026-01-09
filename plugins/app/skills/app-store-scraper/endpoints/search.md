# App Search

Search for apps using keywords and filters.

## Endpoint

```bash
https://itunes.apple.com/search
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `term` | string | Yes | Search query |
| `country` | string | No | Market code (default: `us`) |
| `media` | string | No | Set to `software` for apps |
| `entity` | string | No | Set to `software` for apps |
| `limit` | number | No | Max results (default: 50, max: 200) |
| `lang` | string | No | Language preference |

## Examples

### Basic Search

```bash
curl -s "https://itunes.apple.com/search?term=photo%20editor&country=us&media=software&entity=software&limit=10" | \
  jq '.results[] | {
    name: .trackName,
    id: .trackId,
    developer: .artistName,
    rating: .averageUserRating,
    price: .price
  }'
```

### Search with Pagination

```bash
# iTunes API doesn't support offset, so fetch max and slice client-side
curl -s "https://itunes.apple.com/search?term=fitness&country=us&media=software&entity=software&limit=100" | \
  jq '.results[20:30][] | {name: .trackName, id: .trackId}'
```

### Extract Only App IDs

```bash
curl -s "https://itunes.apple.com/search?term=games&country=us&media=software&entity=software&limit=50" | \
  jq '.results[].trackId'
```

### Search by Category Keywords

```bash
# Search for productivity apps
curl -s "https://itunes.apple.com/search?term=productivity&country=us&media=software&entity=software&limit=30" | \
  jq '.results[] | select(.primaryGenreName == "Productivity") | {
    name: .trackName,
    rating: .averageUserRating
  }'
```

## Response Structure

Same as App Lookup endpoint, but returns array in `results` field.

```json
{
  "resultCount": 50,
  "results": [
    {
      "trackId": 553834731,
      "trackName": "App Name",
      "bundleId": "com.example.app",
      "artistName": "Developer Name",
      "price": 0.0,
      "averageUserRating": 4.5,
      ...
    }
  ]
}
```

## Limitations

- No server-side pagination (no offset parameter)
- Maximum 200 results per request
- Results must be filtered and sorted client-side
- Search algorithm is controlled by Apple

## Use Cases

- Discover apps by keyword
- Find apps in specific categories
- Get competitor analysis data
- Build app recommendation systems
- Extract app IDs for batch processing
