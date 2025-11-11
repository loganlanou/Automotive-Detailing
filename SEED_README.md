# Database Seeding Instructions

This project now includes placeholder vehicles and detailing jobs to showcase the gallery functionality.

## What Was Added

### Vehicles (12 total)
**From Courtesy Automotive Group (6 vehicles):**
1. 2023 Ford F-150 XLT
2. 2022 Chevrolet Silverado 1500 LT
3. 2024 Ram 1500 Big Horn
4. 2021 GMC Sierra 1500 Denali
5. 2023 Chevrolet Tahoe RST
6. 2022 Ford Explorer Limited

**Customer Vehicles (6 vehicles):**
7. 2024 Toyota 4Runner TRD Off-Road
8. 2023 Honda Accord Sport
9. 2022 Toyota Camry XSE
10. 2023 Jeep Wrangler Rubicon
11. 2024 Mazda CX-5 Touring
12. 2021 Subaru Outback Wilderness

### Features
- Each vehicle has a completed detailing job with package information
- Before/After images (using placeholder images from /static/images/work/)
- Customer testimonials for some vehicles
- Featured/highlighted work items
- Courtesy Auto vehicles link to dealership inventory with UTM tracking

## How to Seed the Database

### Option 1: Using the Seed Script (Recommended)
```bash
# Run the seed script to populate the database
go run cmd/seed/main.go
```

This will:
1. Clear existing vehicles, jobs, and media
2. Insert 12 placeholder vehicles
3. Insert 12 completed detailing jobs
4. Insert before/after images

### Option 2: Manual SQL Import
```bash
sqlite3 ./data/detailing.db < ./pkg/db/seed_vehicles.sql
```

## Files Added

- `pkg/db/seed_vehicles.sql` - SQL seed data for vehicles and jobs
- `cmd/seed/main.go` - Go script to seed the database
- `cmd/migrate/main.go` - Database migration script to add missing columns
- `cmd/check_schema/main.go` - Utility to check database schema

## Gallery Features

The gallery page now includes:
- **Book This Service** button - Links to contact/booking page
- **Browse Courtesy Auto Inventory** button - Links to Courtesy Automotive Group website
- All Courtesy Auto vehicles have:
  - Dealership name and location
  - Links to dealership website with UTM tracking
  - "View on Dealership Site" buttons

## Next Steps

1. Replace placeholder images in `/static/images/work/` with actual detailing photos
2. Update Google Review URL in `/web/templates/pages/reviews.templ` (currently set to placeholder)
3. Update phone number from (555) 123-4567 to actual business number
4. Create Courtesy Auto logo image at `/static/images/dealers/courtesy-auto-logo.png`
5. Update Courtesy Auto inventory URL if needed

## Database Schema Updates

The seed script automatically adds missing columns to match the schema:
- `vehicles.color` - Vehicle exterior color
- `vehicles.dealership_name` - Dealership name
- `vehicles.dealership_logo_url` - Logo image path
- `vehicles.dealership_location` - Dealership location
- `jobs.slug` - SEO-friendly URL slug
- `jobs.customer_testimonial` - Customer review text
- `jobs.customer_name` - Customer name for testimonial
- `jobs.highlight_text` - Promotional highlight badge

## Resetting the Data

To clear and re-seed:
```bash
go run cmd/seed/main.go
```

The script automatically clears existing data before inserting new seed data.
