# Reviews

Fetch user reviews for an app with pagination and sorting.

## Endpoint

```bash
https://itunes.apple.com/{country}/rss/customerreviews/page={page}/id={id}/sortby={sort}/json
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `id` | string | Yes | Numeric app ID |
| `country` | string | No | Market code (default: `us`) |
| `page` | number | No | Page number (1-10) |
| `sort` | string | No | Sort order: `mostrecent` or `mosthelpful` |

## Examples

### Recent Reviews

```bash
curl -s "https://itunes.apple.com/us/rss/customerreviews/page=1/id=553834731/sortby=mostrecent/json" | \
  jq '.feed.entry[1:] | .[] | {
    title: .title.label,
    author: .author.name.label,
    rating: .["im:rating"].label,
    content: .content.label,
    date: .updated.label
  }'
```

**Note:** Skip first entry (`[1:]`) as it contains app metadata, not a review.

### Most Helpful Reviews

```bash
curl -s "https://itunes.apple.com/us/rss/customerreviews/page=1/id=553834731/sortby=mosthelpful/json" | \
  jq '.feed.entry[1:] | .[] | {
    rating: .["im:rating"].label,
    title: .title.label,
    review: .content.label
  }'
```

### Fetch Multiple Pages

```bash
for page in {1..3}; do
  echo "=== Page $page ==="
  curl -s "https://itunes.apple.com/us/rss/customerreviews/page=${page}/id=553834731/sortby=mostrecent/json" | \
    jq -r '.feed.entry[1:] | .[] | "[\(.["im:rating"].label)★] \(.title.label) - \(.author.name.label)"'
  echo ""
done
```

### Filter by Rating

```bash
# Get only 5-star reviews
curl -s "https://itunes.apple.com/us/rss/customerreviews/page=1/id=553834731/sortby=mostrecent/json" | \
  jq '.feed.entry[1:] | .[] | select(.["im:rating"].label == "5") | {
    title: .title.label,
    author: .author.name.label,
    content: .content.label
  }'
```

### Reviews Summary

```bash
curl -s "https://itunes.apple.com/us/rss/customerreviews/page=1/id=553834731/sortby=mostrecent/json" | \
  jq '.feed.entry[1:] | group_by(.["im:rating"].label) | map({
    rating: .[0]["im:rating"].label,
    count: length
  })'
```

### Export to CSV

```bash
echo "Rating,Title,Author,Date" > reviews.csv

curl -s "https://itunes.apple.com/us/rss/customerreviews/page=1/id=553834731/sortby=mostrecent/json" | \
  jq -r '.feed.entry[1:] | .[] | [
    .["im:rating"].label,
    .title.label,
    .author.name.label,
    .updated.label
  ] | @csv' >> reviews.csv
```

### Get Review with Version Info

```bash
curl -s "https://itunes.apple.com/us/rss/customerreviews/page=1/id=553834731/sortby=mostrecent/json" | \
  jq '.feed.entry[1:] | .[] | {
    rating: .["im:rating"].label,
    title: .title.label,
    version: .["im:version"].label,
    date: .updated.label
  }'
```

## Response Structure

```json
{
  "feed": {
    "entry": [
      {
        "...": "app metadata (skip this - index 0)"
      },
      {
        "author": {
          "name": {"label": "Username"},
          "uri": {"label": "user-url"}
        },
        "updated": {"label": "2024-01-01T12:00:00-07:00"},
        "im:rating": {"label": "5"},
        "im:version": {"label": "1.0.0"},
        "id": {"label": "12345678"},
        "title": {"label": "Great app!"},
        "content": {
          "label": "This app is amazing. Highly recommend...",
          "attributes": {"type": "text"}
        },
        "link": {"attributes": {"href": "review-url"}},
        "im:voteSum": {"label": "10"},
        "im:voteCount": {"label": "12"}
      }
    ]
  }
}
```

## Pagination Limits

- Maximum 10 pages available
- Each page contains ~50 reviews
- First entry is always app metadata (not a review)
- Total accessible reviews: ~500 per app

## Sorting Options

### `mostrecent`
- Shows newest reviews first
- Best for monitoring recent feedback
- Useful for version-specific issues

### `mosthelpful`
- Shows most voted reviews first
- Better quality/detailed reviews
- Good for understanding common themes

## Use Cases

- Monitor user feedback over time
- Analyze sentiment by version
- Identify common issues or praise
- Generate review summaries
- Track rating trends
- Competitive analysis

## Tips

- Always skip first entry (`[1:]`) in results
- Add delay between page requests
- Reviews are region-specific
- Not all apps have reviews in all regions
- Review availability depends on app's release date
