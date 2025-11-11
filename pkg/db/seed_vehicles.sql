-- Seed data for placeholder vehicles and detailing jobs
-- Run this to populate the gallery with example work

-- First, insert packages if they don't exist
INSERT OR IGNORE INTO packages (slug, name, short_desc, price_min, price_max, duration_est, is_active, sort_order) VALUES
('interior-detail', 'Interior Detail', 'Deep vacuum, steam cleaning, leather conditioning, and odor elimination', 15000, 20000, 180, 1, 1),
('exterior-detail', 'Exterior Detail', 'Hand wash, clay bar, polish, and premium wax protection', 20000, 30000, 240, 1, 2),
('full-detail', 'Full Detail', 'Complete interior and exterior detailing for showroom finish', 30000, 50000, 360, 1, 3);

-- Insert vehicles from Courtesy Auto
INSERT INTO vehicles (slug, year, make, model, trim, color, dealership_name, dealership_logo_url, dealership_listing_url, dealership_location, status) VALUES
('2023-ford-f150-xlt', 2023, 'Ford', 'F-150', 'XLT', 'Agate Black', 'Courtesy Automotive Group', '/static/images/dealers/courtesy-auto-logo.png', 'https://www.courtesyautomotivegroup.com/inventory', 'Stanley, WI', 'available'),
('2022-chevrolet-silverado-1500', 2022, 'Chevrolet', 'Silverado 1500', 'LT', 'Summit White', 'Courtesy Automotive Group', '/static/images/dealers/courtesy-auto-logo.png', 'https://www.courtesyautomotivegroup.com/inventory', 'Stanley, WI', 'available'),
('2024-ram-1500-bighorn', 2024, 'Ram', '1500', 'Big Horn', 'Granite Crystal', 'Courtesy Automotive Group', '/static/images/dealers/courtesy-auto-logo.png', 'https://www.courtesyautomotivegroup.com/inventory', 'Stanley, WI', 'available'),
('2021-gmc-sierra-denali', 2021, 'GMC', 'Sierra 1500', 'Denali', 'Carbon Black', 'Courtesy Automotive Group', '/static/images/dealers/courtesy-auto-logo.png', 'https://www.courtesyautomotivegroup.com/inventory', 'Stanley, WI', 'available'),
('2023-chevrolet-tahoe-rst', 2023, 'Chevrolet', 'Tahoe', 'RST', 'Black', 'Courtesy Automotive Group', '/static/images/dealers/courtesy-auto-logo.png', 'https://www.courtesyautomotivegroup.com/inventory', 'Stanley, WI', 'available'),
('2022-ford-explorer-limited', 2022, 'Ford', 'Explorer', 'Limited', 'Iconic Silver', 'Courtesy Automotive Group', '/static/images/dealers/courtesy-auto-logo.png', 'https://www.courtesyautomotivegroup.com/inventory', 'Stanley, WI', 'available');

-- Insert non-dealership customer vehicles
INSERT INTO vehicles (slug, year, make, model, trim, color, status) VALUES
('2024-toyota-4runner-trd', 2024, 'Toyota', '4Runner', 'TRD Off-Road', 'Army Green', 'available'),
('2023-honda-accord-sport', 2023, 'Honda', 'Accord', 'Sport', 'Platinum White Pearl', 'available'),
('2022-toyota-camry-xse', 2022, 'Toyota', 'Camry', 'XSE', 'Supersonic Red', 'available'),
('2023-jeep-wrangler-rubicon', 2023, 'Jeep', 'Wrangler', 'Rubicon', 'Firecracker Red', 'available'),
('2024-mazda-cx5-touring', 2024, 'Mazda', 'CX-5', 'Touring', 'Soul Red Crystal', 'available'),
('2021-subaru-outback-wilderness', 2021, 'Subaru', 'Outback', 'Wilderness', 'Autumn Green', 'available');

-- Insert detailing jobs for these vehicles
-- Courtesy Auto vehicles
INSERT INTO jobs (slug, vehicle_id, package_id, notes, completed_at, featured, highlight_text, customer_testimonial, customer_name) VALUES
('2023-ford-f150-xlt-detail', 1, 3, 'Full detail on this beautiful 2023 F-150 XLT fresh from Courtesy Auto. Complete interior deep clean including leather conditioning and carpet shampooing. Exterior received clay bar treatment, paint correction, and ceramic coating for long-lasting protection. Turned out stunning!', '2024-01-15 14:30:00', 1, 'Featured Work', 'Absolutely incredible transformation! The truck looks better than when I first saw it on the lot.', 'Mike Johnson'),
('2022-chevrolet-silverado-1500-detail', 2, 3, 'Premium full detail package on this 2022 Silverado 1500. Addressed heavy road grime and salt buildup from winter driving. Interior steam cleaning, leather protection, and odor elimination. Exterior polish and wax brought back the factory shine.', '2024-01-20 11:00:00', 0, NULL, 'The team did an amazing job. My Silverado looks showroom fresh!', 'Dave Williams'),
('2024-ram-1500-bighorn-detail', 3, 2, 'Exterior detail focus on this brand new Ram 1500. Paint correction to remove dealer prep swirls, followed by premium ceramic coating application. Wheels detailed and protected. Ready for delivery in pristine condition.', '2024-01-25 16:00:00', 1, 'Ceramic Coated', NULL, NULL),
('2021-gmc-sierra-denali-detail', 4, 3, 'High-end full detail on this loaded GMC Sierra Denali. Special attention to the premium leather interior with conditioning and protection treatment. Exterior multi-stage paint correction brought out incredible depth in the Carbon Black finish.', '2024-02-01 13:30:00', 1, 'Award Winning Detail', 'Worth every penny! The Denali looks absolutely stunning.', 'Robert Martinez'),
('2023-chevrolet-tahoe-rst-detail', 5, 1, 'Interior-focused detail on this family-hauler Tahoe. Deep cleaning of all three rows, carpet extraction, and complete sanitization. Perfect for families looking for that new-car freshness.', '2024-02-05 10:00:00', 0, NULL, NULL, NULL),
('2022-ford-explorer-limited-detail', 6, 3, 'Complete full detail on this Explorer Limited. Interior leather conditioning, dashboard restoration, and deep carpet cleaning. Exterior paint correction and wax protection. Headlight restoration brought back crystal clarity.', '2024-02-10 15:00:00', 0, NULL, NULL, NULL);

-- Customer vehicles
INSERT INTO jobs (slug, vehicle_id, package_id, notes, completed_at, featured, highlight_text, customer_testimonial, customer_name) VALUES
('2024-toyota-4runner-trd-detail', 7, 3, 'Full detail on this adventure-ready 4Runner TRD Off-Road. Removed mud and trail debris, deep cleaned undercarriage, and protected all surfaces. Interior detailing included all cargo area cleaning. Army Green paint looks incredible after correction and coating.', '2024-02-12 14:00:00', 1, 'Adventure Ready', 'My 4Runner has never looked this good! Ready for the next trail.', 'Sarah Thompson'),
('2023-honda-accord-sport-detail', 8, 2, 'Exterior detail package on this sleek Accord Sport. Paint correction removed swirl marks and minor scratches. Premium wax application gives a deep, glossy finish to the Platinum White Pearl paint.', '2024-02-15 11:30:00', 0, NULL, 'Beautiful work on my Accord. The white paint is flawless!', 'James Chen'),
('2022-toyota-camry-xse-detail', 9, 1, 'Interior detail focusing on removing stains from the light-colored seats. Steam cleaning and extraction brought the interior back to like-new condition. Dashboard and console detailed and protected.', '2024-02-18 09:00:00', 0, NULL, NULL, NULL),
('2023-jeep-wrangler-rubicon-detail', 10, 3, 'Full detail on this off-road Wrangler Rubicon. Addressed mud, dirt, and trail debris throughout. Complete interior cleaning including removable floor mats. Exterior wash, clay bar, and protective coating applied to paint and undercarriage.', '2024-02-20 16:30:00', 1, 'Off-Road Clean', 'They got every bit of mud out! Looks factory fresh.', 'Tyler Anderson'),
('2024-mazda-cx5-touring-detail', 11, 2, 'Exterior detail on this gorgeous Soul Red Crystal CX-5. Paint correction enhanced the depth and clarity of Mazda''s signature red. Ceramic coating applied for long-term protection and easy maintenance.', '2024-02-22 13:00:00', 0, NULL, 'The red paint has never looked better! Amazing shine.', 'Jennifer Lee'),
('2021-subaru-outback-wilderness-detail', 12, 3, 'Complete full detail on this Outback Wilderness. Removed adventure debris from interior and exterior. Special attention to the rugged interior materials and cargo area. Autumn Green paint correction and protection.', '2024-02-25 10:30:00', 0, NULL, NULL, NULL);

-- Add media for the jobs (using existing placeholder images from f150)
-- Before images
INSERT INTO media (job_id, url, kind, sort_order, alt_text) VALUES
(1, '/static/images/work/f150-1.jpg', 'before', 1, '2023 Ford F-150 before detailing'),
(2, '/static/images/work/f150-2.jpg', 'before', 1, '2022 Chevrolet Silverado before detailing'),
(3, '/static/images/work/f150-3.jpg', 'before', 1, '2024 Ram 1500 before detailing'),
(4, '/static/images/work/f150-1.jpg', 'before', 1, '2021 GMC Sierra Denali before detailing'),
(5, '/static/images/work/f150-2.jpg', 'before', 1, '2023 Chevrolet Tahoe before detailing'),
(6, '/static/images/work/f150-3.jpg', 'before', 1, '2022 Ford Explorer before detailing'),
(7, '/static/images/work/f150-1.jpg', 'before', 1, '2024 Toyota 4Runner before detailing'),
(8, '/static/images/work/f150-2.jpg', 'before', 1, '2023 Honda Accord before detailing'),
(9, '/static/images/work/f150-3.jpg', 'before', 1, '2022 Toyota Camry before detailing'),
(10, '/static/images/work/f150-1.jpg', 'before', 1, '2023 Jeep Wrangler before detailing'),
(11, '/static/images/work/f150-2.jpg', 'before', 1, '2024 Mazda CX-5 before detailing'),
(12, '/static/images/work/f150-3.jpg', 'before', 1, '2021 Subaru Outback before detailing');

-- After/primary images (these are what show in the gallery)
INSERT INTO media (job_id, url, kind, sort_order, alt_text) VALUES
(1, '/static/images/work/f150-1.jpg', 'after', 1, '2023 Ford F-150 after full detail'),
(2, '/static/images/work/f150-2.jpg', 'after', 1, '2022 Chevrolet Silverado after full detail'),
(3, '/static/images/work/f150-3.jpg', 'after', 1, '2024 Ram 1500 after exterior detail'),
(4, '/static/images/work/f150-1.jpg', 'after', 1, '2021 GMC Sierra Denali after full detail'),
(5, '/static/images/work/f150-2.jpg', 'after', 1, '2023 Chevrolet Tahoe after interior detail'),
(6, '/static/images/work/f150-3.jpg', 'after', 1, '2022 Ford Explorer after full detail'),
(7, '/static/images/work/f150-1.jpg', 'after', 1, '2024 Toyota 4Runner after full detail'),
(8, '/static/images/work/f150-2.jpg', 'after', 1, '2023 Honda Accord after exterior detail'),
(9, '/static/images/work/f150-3.jpg', 'after', 1, '2022 Toyota Camry after interior detail'),
(10, '/static/images/work/f150-1.jpg', 'after', 1, '2023 Jeep Wrangler after full detail'),
(11, '/static/images/work/f150-2.jpg', 'after', 1, '2024 Mazda CX-5 after exterior detail'),
(12, '/static/images/work/f150-3.jpg', 'after', 1, '2021 Subaru Outback after full detail');

-- Gallery images (additional photos)
INSERT INTO media (job_id, url, kind, sort_order, alt_text) VALUES
(1, '/static/images/work/f150-2.jpg', 'gallery', 2, 'Interior detail close-up'),
(1, '/static/images/work/f150-3.jpg', 'gallery', 3, 'Wheel and tire detail'),
(4, '/static/images/work/f150-2.jpg', 'gallery', 2, 'Paint correction results'),
(4, '/static/images/work/f150-3.jpg', 'gallery', 3, 'Engine bay detail'),
(7, '/static/images/work/f150-2.jpg', 'gallery', 2, 'Cargo area cleaning'),
(10, '/static/images/work/f150-2.jpg', 'gallery', 2, 'Undercarriage protection');
