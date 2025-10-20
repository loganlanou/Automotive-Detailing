-- Schema for Detailing Pass
PRAGMA foreign_keys = ON;

-- Detailing packages/services
CREATE TABLE IF NOT EXISTS packages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    short_desc TEXT,
    long_desc TEXT,
    price_min INTEGER, -- in cents
    price_max INTEGER, -- in cents
    duration_est INTEGER, -- in minutes
    is_active BOOLEAN DEFAULT 1,
    sort_order INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Vehicles (trade-ins or customer cars)
CREATE TABLE IF NOT EXISTS vehicles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    vin TEXT,
    year INTEGER,
    make TEXT NOT NULL,
    model TEXT NOT NULL,
    trim TEXT,
    color TEXT, -- exterior color
    price INTEGER, -- vehicle price in cents (for dealer listings)
    stock_number TEXT,
    dealership_name TEXT, -- e.g., "Courtesy Auto Stanley"
    dealership_logo_url TEXT,
    dealership_listing_url TEXT,
    dealership_location TEXT, -- e.g., "Stanley, NC"
    status TEXT DEFAULT 'available', -- available|sold|archived
    posted_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Completed detailing jobs
CREATE TABLE IF NOT EXISTS jobs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE, -- SEO-friendly URL slug (auto-generated from vehicle+date)
    vehicle_id INTEGER,
    package_id INTEGER,
    technician TEXT,
    notes TEXT, -- detailed description/story of the work
    completed_at DATETIME,
    duration_actual INTEGER, -- actual time spent in minutes
    featured BOOLEAN DEFAULT 0,
    display_price INTEGER, -- optional price to show for this job
    highlight_text TEXT, -- promotional highlight (e.g., "Award Winning Detail")
    customer_testimonial TEXT, -- optional testimonial from customer
    customer_name TEXT, -- name for testimonial attribution
    meta_description TEXT, -- SEO meta description
    meta_keywords TEXT, -- SEO keywords
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (vehicle_id) REFERENCES vehicles(id) ON DELETE SET NULL,
    FOREIGN KEY (package_id) REFERENCES packages(id) ON DELETE SET NULL
);

-- Media (images/videos for jobs and vehicles)
CREATE TABLE IF NOT EXISTS media (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    job_id INTEGER,
    vehicle_id INTEGER,
    url TEXT NOT NULL,
    kind TEXT DEFAULT 'after', -- before|after|hero|gallery
    sort_order INTEGER DEFAULT 0,
    alt_text TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES vehicles(id) ON DELETE CASCADE
);

-- Customer reviews/testimonials
CREATE TABLE IF NOT EXISTS reviews (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author TEXT NOT NULL,
    rating INTEGER DEFAULT 5, -- 1-5
    body TEXT,
    source TEXT, -- google|facebook|manual
    is_featured BOOLEAN DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Blog posts (optional)
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    excerpt TEXT,
    body TEXT, -- Markdown or HTML
    author TEXT DEFAULT 'Detailing Pass Team',
    published_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX IF NOT EXISTS idx_vehicles_make ON vehicles(make);
CREATE INDEX IF NOT EXISTS idx_vehicles_model ON vehicles(model);
CREATE INDEX IF NOT EXISTS idx_vehicles_year ON vehicles(year);
CREATE INDEX IF NOT EXISTS idx_vehicles_status ON vehicles(status);
CREATE INDEX IF NOT EXISTS idx_jobs_featured ON jobs(featured);
CREATE INDEX IF NOT EXISTS idx_jobs_completed_at ON jobs(completed_at);
CREATE INDEX IF NOT EXISTS idx_media_job_id ON media(job_id);
CREATE INDEX IF NOT EXISTS idx_media_vehicle_id ON media(vehicle_id);
CREATE INDEX IF NOT EXISTS idx_posts_published_at ON posts(published_at);
