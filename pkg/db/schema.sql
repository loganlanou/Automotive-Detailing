-- Schema for C Auto Detailing Studio
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

-- Gallery groups (collections of images for a single vehicle/project)
CREATE TABLE IF NOT EXISTS gallery_groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    vehicle_make TEXT,
    vehicle_model TEXT,
    vehicle_year INTEGER,
    description TEXT,
    is_featured BOOLEAN DEFAULT 0,
    sort_order INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Media (images for gallery groups)
CREATE TABLE IF NOT EXISTS media (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    gallery_group_id INTEGER,
    url TEXT NOT NULL,
    kind TEXT DEFAULT 'gallery', -- before|after|hero|gallery
    sort_order INTEGER DEFAULT 0,
    alt_text TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (gallery_group_id) REFERENCES gallery_groups(id) ON DELETE CASCADE
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

-- Booking requests & calendar slots
CREATE TABLE IF NOT EXISTS bookings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    customer_name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT,
    vehicle_details TEXT,
    service_interest TEXT,
    notes TEXT,
    requested_start DATETIME NOT NULL,
    requested_end DATETIME NOT NULL,
    status TEXT DEFAULT 'pending', -- pending|confirmed|declined|cancelled
    source TEXT DEFAULT 'web',
    internal_notes TEXT,
    clerk_user_id TEXT, -- Clerk user ID if logged in
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX IF NOT EXISTS idx_gallery_groups_featured ON gallery_groups(is_featured);
CREATE INDEX IF NOT EXISTS idx_gallery_groups_sort ON gallery_groups(sort_order);
CREATE INDEX IF NOT EXISTS idx_media_gallery_group_id ON media(gallery_group_id);
CREATE INDEX IF NOT EXISTS idx_media_sort ON media(sort_order);
CREATE INDEX IF NOT EXISTS idx_bookings_requested_start ON bookings(requested_start);
CREATE INDEX IF NOT EXISTS idx_bookings_status ON bookings(status);
CREATE INDEX IF NOT EXISTS idx_reviews_featured ON reviews(is_featured);
