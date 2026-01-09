# Ratings Histogram

Get detailed rating distribution (1-5 star breakdown).

## Endpoint

```bash
https://itunes.apple.com/{country}/customer-reviews/id{id}?displayable-kind=11
```

## Required Headers

```bash
X-Apple-Store-Front: {storefront_id}-1,29 l=en
```

## Common Storefront IDs

| Country | Code | Storefront ID | Country | Code | Storefront ID |
|---------|------|---------------|---------|------|---------------|
| United States | `us` | `143441` | United Kingdom | `gb` | `143444` |
| Germany | `de` | `143443` | France | `fr` | `143442` |
| Japan | `jp` | `143462` | Australia | `au` | `143460` |
| Canada | `ca` | `143455` | Spain | `es` | `143454` |
| Italy | `it` | `143450` | Brazil | `br` | `143503` |
| India | `in` | `143467` | Mexico | `mx` | `143468` |
| South Korea | `kr` | `143466` | China | `cn` | `143465` |

## Examples

### Basic Rating Histogram

```bash
curl -s "https://itunes.apple.com/us/customer-reviews/id553834731?displayable-kind=11" \
  -H "X-Apple-Store-Front: 143441-1,29 l=en" | \
  jq '.userReviewList.userReviewStatistics | {
    avgRating: .averageUserRating,
    totalRatings: .userRatingCount,
    ratingCountList: .ratingCountList
  }'
```

### Extract Star Breakdown

```bash
curl -s "https://itunes.apple.com/us/customer-reviews/id553834731?displayable-kind=11" \
  -H "X-Apple-Store-Front: 143441-1,29 l=en" | \
  jq '.userReviewList.userReviewStatistics.ratingCountList[] | {
    stars: .ratingCount,
    count: .userRatingCount
  }'
```

### Calculate Percentages

```bash
curl -s "https://itunes.apple.com/us/customer-reviews/id553834731?displayable-kind=11" \
  -H "X-Apple-Store-Front: 143441-1,29 l=en" | \
  jq '
    .userReviewList.userReviewStatistics |
    .totalCount = .userRatingCount |
    .ratingCountList[] |= (
      . + {percentage: ((.userRatingCount / .totalCount) * 100 | round)}
    )
  '
```

### Format as Chart

```bash
curl -s "https://itunes.apple.com/us/customer-reviews/id553834731?displayable-kind=11" \
  -H "X-Apple-Store-Front: 143441-1,29 l=en" | \
  jq -r '
    .userReviewList.userReviewStatistics.ratingCountList[] |
    "[\(.ratingCount)★] \("█" * (.userRatingCount / 100 | floor)) \(.userRatingCount)"
  '
```

### Multiple Countries

```bash
for country in us gb de jp; do
  case $country in
    us) storefront=143441 ;;
    gb) storefront=143444 ;;
    de) storefront=143443 ;;
    jp) storefront=143462 ;;
  esac

  echo "=== $country ==="
  curl -s "https://itunes.apple.com/$country/customer-reviews/id553834731?displayable-kind=11" \
    -H "X-Apple-Store-Front: $storefront-1,29 l=en" | \
    jq '.userReviewList.userReviewStatistics | {
      avgRating,
      totalRatings: .userRatingCount
    }'
  echo ""
done
```

### Compare Current vs All-Time Ratings

Some responses include both current version and all-time ratings:

```bash
curl -s "https://itunes.apple.com/us/customer-reviews/id553834731?displayable-kind=11" \
  -H "X-Apple-Store-Front: 143441-1,29 l=en" | \
  jq '{
    currentVersion: .userReviewList.userReviewStatistics.averageUserRating,
    allVersions: .userReviewList.userReviewStatistics.averageUserRatingForCurrentVersion
  }'
```

## Response Structure

```json
{
  "userReviewList": {
    "userReviewStatistics": {
      "averageUserRating": 4.5,
      "userRatingCount": 12345,
      "ratingCountList": [
        {
          "ratingCount": 5,
          "userRatingCount": 8000
        },
        {
          "ratingCount": 4,
          "userRatingCount": 2500
        },
        {
          "ratingCount": 3,
          "userRatingCount": 1000
        },
        {
          "ratingCount": 2,
          "userRatingCount": 500
        },
        {
          "ratingCount": 1,
          "userRatingCount": 345
        }
      ]
    }
  }
}
```

## Important Notes

- **Requires X-Apple-Store-Front header** - Request will fail without it
- Storefront ID must match the country in the URL
- Response structure may vary slightly by region
- Some apps may not have rating histogram data
- Data may include current version vs all-time ratings

## Finding Storefront IDs

If you need a storefront ID not listed:

```bash
# Visit an App Store page and inspect headers
curl -I "https://apps.apple.com/us/app/id553834731" | grep -i "x-apple"
```

Or use these common patterns:
- US: 143441
- UK: 143444
- Most EU countries: 1434XX range
- Asia-Pacific: 1434XX-1434XX range

## Use Cases

- Detailed rating analysis
- Compare ratings across regions
- Calculate rating distribution
- Monitor rating trends over time
- Generate rating visualizations
- Quality metrics reporting

## Combining with Other Data

```bash
APP_ID=553834731

# Get basic info
APP_NAME=$(curl -s "https://itunes.apple.com/lookup?id=$APP_ID" | jq -r '.results[0].trackName')

# Get rating histogram
echo "Rating breakdown for: $APP_NAME"
curl -s "https://itunes.apple.com/us/customer-reviews/id$APP_ID?displayable-kind=11" \
  -H "X-Apple-Store-Front: 143441-1,29 l=en" | \
  jq -r '.userReviewList.userReviewStatistics.ratingCountList[] |
    "\(.ratingCount) stars: \(.userRatingCount) ratings"'
```

## Error Handling

```bash
RESPONSE=$(curl -s "https://itunes.apple.com/us/customer-reviews/id999999999?displayable-kind=11" \
  -H "X-Apple-Store-Front: 143441-1,29 l=en")

if echo "$RESPONSE" | jq -e '.errorMessage' > /dev/null 2>&1; then
  echo "Error: App not found or no ratings available"
else
  echo "$RESPONSE" | jq '.userReviewList.userReviewStatistics'
fi
```
