# Detailing Pass

Premium automotive detailing website built with Go, Echo, Templ, and Tailwind CSS.

## Features

- Server-side rendered with Go + Echo + Templ
- SQLite database with type-safe SQLC queries
- Tailwind CSS v4 for styling
- Progressive enhancement (minimal JavaScript)
- Filterable work gallery
- Admin dashboard (Clerk authentication)
- Dealer sync API + CSV export
- Contact form with spam protection
- Cal.com booking integration ready
- Fully accessible (WCAG AA)
- SEO optimized with structured data

## Tech Stack

- **Backend:** Go 1.23, Echo web framework
- **Database:** SQLite + SQLC
- **Templates:** Templ (type-safe Go templates)
- **Styling:** Tailwind CSS v4 + PostCSS
- **Auth:** Clerk (admin only)
- **Dev Tools:** Air (hot reload), direnv, Make

## Prerequisites

- Go 1.23+
- Node.js 18+ (for Tailwind CLI)
- Make
- SQLite3

## Quick Start

1. **Clone and setup:**
   ```bash
   git clone <repository>
   cd Automotive-Detailing
   cp .env.example .env
   ```

2. **Install dependencies:**
   ```bash
   make install
   ```

3. **Setup database:**
   ```bash
   make migrate
   make sqlc
   ```

4. **Run development server:**
   ```bash
   make dev
   ```

   The site will be available at http://localhost:8080

## Available Commands

```bash
make install      # Install all dependencies (Go + Node)
make dev          # Start dev server with hot reload
make build        # Build production binary
make migrate      # Run database migrations
make sqlc         # Generate SQLC code
make templ        # Generate Templ templates
make tailwind     # Build Tailwind CSS
make test         # Run tests
make fmt          # Format code
make clean        # Clean build artifacts
```

## Deployment

- **Vercel (default):** `vercel.json` + `api/index.go` configure the project for Go Serverless Functions. See `DEPLOYMENT.md` for the exact checklist (Git integration, build command, env vars, and troubleshooting).
- **Other targets:** Render or Fly.io instructions are also documented in `DEPLOYMENT.md` if you prefer a long-running server with persistent storage.

## Project Structure

```
.
├── cmd/server/           # Application entry point
├── pkg/
│   ├── server/          # HTTP server, routes, middleware
│   │   └── handlers/    # Request handlers
│   ├── db/              # Database schema, queries, SQLC config
│   └── models/          # Data models
├── web/
│   ├── templates/       # Templ templates
│   │   ├── layout.templ
│   │   ├── partials/    # Reusable components
│   │   └── pages/       # Page templates
│   └── static/
│       ├── css/         # Tailwind styles
│       ├── js/          # Vanilla JavaScript
│       └── uploads/     # User uploads
├── public/              # Static assets (robots.txt, sitemap.xml)
├── data/                # SQLite database (gitignored)
└── Makefile            # Build commands
```

## Environment Variables

Copy `.env.example` to `.env` and configure:

- `PORT` - Server port (default: 8080)
- `DATABASE_PATH` - SQLite database path
- `CLERK_SECRET_KEY` - Clerk authentication
- `SMTP_*` - Email configuration for contact form
- `DEALER_*` - Dealer sync API configuration
- `CALCOM_EMBED_URL` - Cal.com booking URL

## Database

The project uses SQLite with SQLC for type-safe queries.

**Schema:** See `pkg/db/schema.sql`

**Tables:**
- `packages` - Detailing service packages
- `vehicles` - Vehicles (with dealer listing URLs)
- `jobs` - Completed detailing work
- `media` - Before/after images
- `reviews` - Customer testimonials
- `posts` - Blog posts (optional)

To modify the database:
1. Edit `pkg/db/schema.sql`
2. Run `make migrate`
3. Update queries in `pkg/db/queries.sql`
4. Run `make sqlc` to regenerate Go code

## Design System

The design system uses the following brand colors:

- Primary: `#003DA5` (Ford Blue)
- Secondary: `#111827` (Slate 900)
- Accent: `#60A5FA` (Sky 400)
- Background: `#0B0F13` (Near-black)
- Foreground: `#F3F4F6` (Gray 100)

**Fonts:**
- Headings: Poppins (600, 700)
- Body: Inter (400, 500, 600)

View the complete design system at `/style` when running the server.

## Image Guidelines

For optimal performance:

1. **Format:** Use WebP when possible, fallback to JPEG
2. **Sizing:**
   - Hero images: 1920x1080px (max 200KB)
   - Gallery thumbnails: 800x600px (max 100KB)
   - Before/after: 1200x900px (max 150KB each)
3. **Alt text:** Always provide descriptive alt text for accessibility
4. **Lazy loading:** Enabled by default on gallery images

## Content Management

### Adding Detailing Packages

1. Go to `/admin/packages` (requires Clerk authentication)
2. Create new package with pricing, description, duration
3. Set `is_active` to show on the public site

### Adding Work Items

1. Create or select a vehicle in `/admin/vehicles`
2. Create a job in `/admin/jobs` linking the vehicle and package
3. Upload before/after media in `/admin/media`
4. Mark as `featured` to show on homepage
5. Optionally add dealership listing URL

### Dealer Sync

**API Endpoint:** `POST /api/dealer/sync`

**Payload:**
```json
{
  "vehicle": {
    "vin": "1HGBH41JXMN109186",
    "year": 2024,
    "make": "Ford",
    "model": "F-150",
    "trim": "XLT",
    "price": 45000,
    "stock_number": "F150-001",
    "dealership_listing_url": "https://dealer.com/inventory/12345"
  },
  "job": {
    "package_slug": "full-detail",
    "completed_at": "2025-01-15T10:00:00Z",
    "featured": true
  }
}
```

**CSV Export:** Available at `/api/dealer/export`

## Deployment

### Option 1: Render / Fly.io (Recommended)

1. Build the Docker image or use the binary
2. Set environment variables
3. Ensure SQLite volume is persistent
4. Deploy!

### Option 2: VPS (Traditional)

1. Build binary: `make build`
2. Copy `bin/server` to VPS
3. Copy `.env` with production values
4. Run behind Caddy or NGINX as reverse proxy
5. Use systemd for process management

**Example systemd service:**
```ini
[Unit]
Description=Detailing Pass Web Server
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/var/www/detailingpass
ExecStart=/var/www/detailingpass/bin/server
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

## Testing

Run Go tests:
```bash
make test
```

Optional Playwright tests for accessibility and smoke testing:
```bash
# TODO: Add Playwright setup
```

## Launch Checklist

Before going live:

- [ ] Replace all `{{REPLACE_ME}}` placeholders with real content
- [ ] Add real images (logo, hero, before/after gallery)
- [ ] Configure SMTP for contact form
- [ ] Set up Clerk authentication for admin
- [ ] Add Cal.com booking embed
- [ ] Test contact form submission
- [ ] Run Lighthouse audit (target: Performance ≥85, A11y ≥95, SEO ≥95)
- [ ] Run axe accessibility checks
- [ ] Test on mobile devices
- [ ] Configure production domain in `.env`
- [ ] Update `robots.txt` and `sitemap.xml` with production domain
- [ ] Set up SSL certificate (automatic with Caddy/Render)
- [ ] Configure backup strategy for SQLite database
- [ ] Test dealer sync API
- [ ] Add Google Analytics (optional)

## Support

For questions or issues, please open an issue on GitHub or contact the development team.

## License

Proprietary - All rights reserved
