# Local Development with Cloudflare Tunnels

## Why Tunnels?

ChatGPT MCP servers require HTTPS endpoints. During local development, use Cloudflare Tunnels to expose your local server with a public HTTPS URL.

## Quick Tunnel vs Named Tunnel

| Type | Command | URL | Use Case |
|------|---------|-----|----------|
| Quick | `cloudflared tunnel --url http://localhost:3000` | Random subdomain each run | Quick testing |
| Named | `cloudflared tunnel run` | Persistent subdomain | Development workflow |

For MCP development, use **named tunnels** to avoid reconfiguring ChatGPT each session.

## Setup: Named Tunnel with Persistent URL

### 1. Install cloudflared

```bash
# macOS
brew install cloudflared

# Other platforms: https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/downloads/
```

### 2. Authenticate

```bash
cloudflared tunnel login
```

Opens browser to authenticate. Creates cert at `~/.cloudflared/cert.pem`.

### 3. Create Named Tunnel

```bash
cloudflared tunnel create mcp-dev
```

Note the tunnel ID (UUID) from the output.

### 4. Create Config File

Create `~/.cloudflared/config.yml`:

```yaml
tunnel: <TUNNEL_ID>
credentials-file: /Users/<USERNAME>/.cloudflared/<TUNNEL_ID>.json

ingress:
  - hostname: mcp-dev.yourdomain.com
    service: http://localhost:3000
  - service: http_status:404
```

Replace:
- `<TUNNEL_ID>` with UUID from step 3
- `<USERNAME>` with your system username
- `yourdomain.com` with a domain in your Cloudflare account

### 5. Route DNS (One-time)

```bash
cloudflared tunnel route dns <TUNNEL_ID> mcp-dev.yourdomain.com
```

Creates CNAME record pointing subdomain to your tunnel.

### 6. Run Tunnel

```bash
cloudflared tunnel run
```

Your MCP server is now accessible at `https://mcp-dev.yourdomain.com`.

## Environment Configuration

Update your MCP server to use the tunnel URL:

```typescript
const MCP_SERVER_URL = process.env.MCP_SERVER_URL || "http://localhost:3000";
const WIDGET_CDN_URL = process.env.WIDGET_CDN_URL || "http://localhost:5173";
```

For ChatGPT integration, configure your app to use the tunnel URL:
- MCP Server: `https://mcp-dev.yourdomain.com`

## Multiple Services

Route multiple local services through one tunnel:

```yaml
tunnel: <TUNNEL_ID>
credentials-file: /Users/<USERNAME>/.cloudflared/<TUNNEL_ID>.json

ingress:
  - hostname: mcp.yourdomain.com
    service: http://localhost:3000
  - hostname: widgets.yourdomain.com
    service: http://localhost:5173
  - service: http_status:404
```

Run DNS routing for each hostname:

```bash
cloudflared tunnel route dns <TUNNEL_ID> mcp.yourdomain.com
cloudflared tunnel route dns <TUNNEL_ID> widgets.yourdomain.com
```

## Troubleshooting

### ERR_NAME_NOT_RESOLVED

DNS not configured. Run the route command:

```bash
cloudflared tunnel route dns <TUNNEL_ID> <hostname>
```

### Tunnel runs but connection refused

Ensure your local server is running on the configured port.

### Permission denied on credentials file

```bash
chmod 600 ~/.cloudflared/<TUNNEL_ID>.json
```

## Useful Commands

```bash
# List tunnels
cloudflared tunnel list

# Check tunnel status
cloudflared tunnel info <TUNNEL_ID>

# Delete tunnel
cloudflared tunnel delete <TUNNEL_NAME>

# View config
cloudflared tunnel ingress validate
```
