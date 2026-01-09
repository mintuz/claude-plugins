# App Lookup

Retrieve detailed information for a specific app using its numeric ID or bundle identifier.

## Endpoint

```bash
https://itunes.apple.com/lookup
```

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `id` | string | Yes* | Numeric app ID (trackId) |
| `bundleId` | string | Yes* | Bundle identifier (e.g., com.example.app) |
| `country` | string | No | Market code (default: `us`) |
| `entity` | string | No | Filter type (use `software` for apps) |
| `lang` | string | No | Language preference |

*Use either `id` OR `bundleId`, not both

## Examples

### Lookup by App ID

```bash
curl -s "https://itunes.apple.com/lookup?id=553834731&country=us&entity=software" | \
  jq '.results[0] | {
    name: .trackName,
    developer: .artistName,
    bundleId: .bundleId,
    version: .version,
    price: .price,
    rating: .averageUserRating,
    ratingCount: .userRatingCount,
    description: .description,
    releaseDate: .releaseDate,
    size: .fileSizeBytes,
    languages: .languageCodesISO2A
  }'
```

### Lookup by Bundle ID

```bash
curl -s "https://itunes.apple.com/lookup?bundleId=com.apple.Numbers&country=us&entity=software" | \
  jq '.results[0] | {
    name: .trackName,
    appId: .trackId,
    developer: .artistName,
    category: .primaryGenreName
  }'
```

### Lookup Multiple Apps

```bash
curl -s "https://itunes.apple.com/lookup?id=553834731,361309726&country=us&entity=software" | \
  jq '.results[] | {name: .trackName, id: .trackId}'
```

## Key Response Fields

```json
{
  "trackId": 553834731,
  "trackName": "App Name",
  "bundleId": "com.example.app",
  "artistName": "Developer Name",
  "artistId": 12345,
  "price": 0.0,
  "currency": "USD",
  "version": "1.0.0",
  "averageUserRating": 4.5,
  "userRatingCount": 1234,
  "description": "Full app description...",
  "releaseDate": "2024-01-01T00:00:00Z",
  "fileSizeBytes": "52428800",
  "contentAdvisoryRating": "4+",
  "languageCodesISO2A": ["EN", "ES", "FR"],
  "genres": ["Productivity", "Business"],
  "primaryGenreName": "Productivity"
}
```

## Use Cases

- Get comprehensive app metadata
- Lookup apps by bundle identifier
- Fetch multiple apps in a single request
- Retrieve developer ID for further queries
- Check app availability in different countries
