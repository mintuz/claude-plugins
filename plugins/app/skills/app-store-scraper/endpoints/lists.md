# App Lists & Feeds

Browse curated App Store collections and categories.

## Endpoint

```bash
https://itunes.apple.com/{country}/rss/{collection}/[genre={category}]/limit={limit}/json
```

## Collections

Common collection values:

- `topfreeapplications` - Top Free Apps
- `toppaidapplications` - Top Paid Apps
- `topgrossingapplications` - Top Grossing Apps
- `newapplications` - New Apps
- `newfreeapplications` - New Free Apps

## Categories

Common genre IDs (use as `genre={id}`):

| ID | Category | ID | Category |
|----|----------|----|----------|
| `6014` | Games | `6000` | Business |
| `6015` | Finance | `6016` | Health & Fitness |
| `6017` | Lifestyle | `6018` | Medical |
| `6020` | Music | `6021` | Navigation |
| `6023` | Photo & Video | `6024` | Productivity |
| `6026` | Social Networking | `6027` | Sports |
| `6012` | Travel | `6002` | Utilities |

## Examples

### Top Free Apps (All Categories)

```bash
curl -s "https://itunes.apple.com/us/rss/topfreeapplications/limit=25/json" | \
  jq '.feed.entry[] | {
    name: .["im:name"].label,
    id: .id.attributes["im:id"],
    developer: .["im:artist"].label,
    category: .category.attributes.label
  }'
```

### Top Free Games

```bash
curl -s "https://itunes.apple.com/us/rss/topfreeapplications/genre=6014/limit=50/json" | \
  jq '.feed.entry[] | {
    name: .["im:name"].label,
    id: .id.attributes["im:id"],
    price: .["im:price"].attributes.amount
  }'
```

### New Apps in Productivity

```bash
curl -s "https://itunes.apple.com/us/rss/newapplications/genre=6024/limit=10/json" | \
  jq '.feed.entry[] | {
    name: .["im:name"].label,
    id: .id.attributes["im:id"],
    releaseDate: .["im:releaseDate"].label
  }'
```

### Top Grossing Apps

```bash
curl -s "https://itunes.apple.com/us/rss/topgrossingapplications/limit=20/json" | \
  jq '.feed.entry[] | {
    name: .["im:name"].label,
    id: .id.attributes["im:id"],
    category: .category.attributes.label
  }'
```

## Response Structure

```json
{
  "feed": {
    "title": {"label": "Top Free Applications"},
    "updated": {"label": "2024-01-01T12:00:00-07:00"},
    "entry": [
      {
        "im:name": {"label": "App Name"},
        "im:image": [
          {"label": "icon_url_small", "attributes": {"height": "53"}},
          {"label": "icon_url_medium", "attributes": {"height": "75"}},
          {"label": "icon_url_large", "attributes": {"height": "100"}}
        ],
        "im:artist": {"label": "Developer Name"},
        "category": {"attributes": {"label": "Category", "im:id": "6024"}},
        "id": {"attributes": {"im:id": "553834731"}},
        "im:releaseDate": {"label": "2024-01-01T00:00:00-07:00"},
        "im:price": {"attributes": {"amount": "0.00", "currency": "USD"}},
        "link": {"attributes": {"href": "https://apps.apple.com/..."}}
      }
    ]
  }
}
```

## Combining with Lookup

RSS feeds provide limited metadata. Get full details by extracting IDs and using lookup:

```bash
# Get top free app IDs
IDS=$(curl -s "https://itunes.apple.com/us/rss/topfreeapplications/limit=10/json" | \
  jq -r '.feed.entry[].id.attributes["im:id"]' | tr '\n' ',')

# Get full details
curl -s "https://itunes.apple.com/lookup?id=$IDS&entity=software" | \
  jq '.results[] | {name: .trackName, rating: .averageUserRating, description: .description}'
```

## Use Cases

- Track top-performing apps
- Monitor category trends
- Discover newly released apps
- Analyze competitive positioning
- Build app ranking dashboards
