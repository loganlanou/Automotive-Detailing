-- Seed data for Ford vehicle gallery
-- Run this to populate the gallery with placeholder images

-- Clear existing gallery data
DELETE FROM media WHERE gallery_group_id IS NOT NULL;
DELETE FROM gallery_groups;

-- Reset auto-increment
DELETE FROM sqlite_sequence WHERE name='gallery_groups';
DELETE FROM sqlite_sequence WHERE name='media';

-- Insert Ford F-150 gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('2024 Ford F-150 XLT - Full Detail', 'f150-xlt-white-2024', 'Ford', 'F-150 XLT', 2024, 'Complete exterior and interior detail on this stunning white F-150. Paint correction, ceramic coating, and premium interior conditioning.', 1, 1),
('2023 Ford F-150 Lariat - Interior Reset', 'f150-lariat-black-2023', 'Ford', 'F-150 Lariat', 2023, 'Deep interior cleaning and leather conditioning on this black F-150 Lariat. Dashboard restoration and complete sanitization.', 0, 2),
('2024 Ford F-150 Raptor R - Performance Detail', 'f150-raptor-r-2024', 'Ford', 'F-150 Raptor R', 2024, 'High-performance detail package on this aggressive Raptor R. Paint protection film and ceramic coating for off-road protection.', 1, 3),
('2023 Ford F-150 Raptor - Adventure Ready', 'f150-raptor-white-2023', 'Ford', 'F-150 Raptor', 2023, 'Full detail and protection package on this white Raptor. Prepared for any adventure with premium protection.', 0, 4);

-- Insert Ford F-250/F-350 Super Duty gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('2024 Ford F-250 Super Duty - Work Truck Detail', 'f250-super-duty-2024', 'Ford', 'F-250 Super Duty', 2024, 'Heavy-duty detail for this workhorse F-250. Complete decontamination, paint correction, and protective coating.', 0, 5),
('2023 Ford F-350 Platinum - Premium Detail', 'f350-platinum-2023', 'Ford', 'F-350 Platinum', 2023, 'Luxury treatment for this premium F-350 Platinum. Multi-stage paint correction and top-tier ceramic coating.', 1, 6);

-- Insert Ford Mustang gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('2024 Ford Mustang GT - Track Ready', 'mustang-gt-green-2024', 'Ford', 'Mustang GT', 2024, 'Performance detail on this stunning green Mustang GT. Paint correction and ceramic coating for showroom shine.', 1, 7),
('2023 Ford Mustang GT - Urban Style', 'mustang-gt-black-2023', 'Ford', 'Mustang GT', 2023, 'Sleek black Mustang GT detailed to perfection. Full paint correction and glass coating.', 0, 8),
('2024 Ford Mustang Mach 1 - Orange Fury', 'mustang-mach1-orange-2024', 'Ford', 'Mustang Mach 1', 2024, 'Vibrant orange Mustang Mach 1 with complete detail package. Enhanced paint clarity and depth.', 1, 9),
('2023 Ford Mustang Shelby GT500 - Snow White', 'mustang-shelby-white-2023', 'Ford', 'Mustang Shelby GT500', 2023, 'Premium detail on this iconic Shelby GT500. Paint correction and ceramic coating protection.', 1, 10),
('1969 Ford Mustang Boss 302 - Classic Restoration', 'mustang-boss-302-vintage', 'Ford', 'Mustang Boss 302', 1969, 'Classic Mustang restoration detail. Careful paint correction preserving original character.', 1, 11),
('2024 Ford Mustang EcoBoost - Daily Driver Detail', 'mustang-ecoboost-red-2024', 'Ford', 'Mustang EcoBoost', 2024, 'Complete detail package for this red Mustang EcoBoost. Perfect daily driver protection.', 0, 12);

-- Insert Ford Explorer gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('2024 Ford Explorer ST - Family Performance', 'explorer-st-white-2024', 'Ford', 'Explorer ST', 2024, 'Full detail on this white Explorer ST. Family-friendly SUV with sport appeal, detailed inside and out.', 1, 13);

-- Insert Ford Escape gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('2024 Ford Escape Hybrid - Eco Detail', 'escape-hybrid-2024', 'Ford', 'Escape Hybrid', 2024, 'Compact SUV detail with eco-friendly products. Complete interior and exterior refresh.', 0, 14);

-- Insert Ford Expedition gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('2024 Ford Expedition Limited - Full Size Luxury', 'expedition-limited-2024', 'Ford', 'Expedition Limited', 2024, 'Premium detail on this full-size luxury SUV. All three rows deep cleaned and conditioned.', 0, 15);

-- Insert Ford Bronco gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('2024 Ford Bronco Badlands - Trail Ready', 'bronco-badlands-2024', 'Ford', 'Bronco Badlands', 2024, 'Adventure-ready detail package on this Bronco Badlands. Protected for off-road use.', 1, 16);

-- Insert Ford Maverick gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('2024 Ford Maverick Lariat - Compact Truck Detail', 'maverick-lariat-2024', 'Ford', 'Maverick Lariat', 2024, 'Full detail on this versatile compact truck. Perfect balance of utility and style.', 0, 17);

-- Insert Classic Ford Trucks gallery groups
INSERT INTO gallery_groups (title, slug, vehicle_make, vehicle_model, vehicle_year, description, is_featured, sort_order) VALUES
('1978 Ford F-100 - Classic Restoration', 'f100-classic-red-1978', 'Ford', 'F-100', 1978, 'Vintage truck restoration detail. Careful treatment preserving classic character.', 1, 18),
('1965 Ford F-250 - Heritage Detail', 'f250-heritage-blue-1965', 'Ford', 'F-250', 1965, 'Classic American truck detailed with care. Original paint enhanced and protected.', 0, 19);

-- Now insert media for each gallery group
-- F-150 XLT White (ID 1)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(1, '/static/images/gallery/f150-white-1.jpg', 'hero', 0, '2024 Ford F-150 XLT exterior'),
(1, '/static/images/gallery/f150-closeup-1.jpg', 'gallery', 1, 'F-150 front grille detail'),
(1, '/static/images/gallery/ford-interior-1.jpg', 'gallery', 2, 'F-150 interior dashboard'),
(1, '/static/images/gallery/car-interior-leather-1.jpg', 'gallery', 3, 'F-150 leather seats detail');

-- F-150 Lariat Black (ID 2)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(2, '/static/images/gallery/f150-black-1.jpg', 'hero', 0, '2023 Ford F-150 Lariat exterior'),
(2, '/static/images/gallery/ford-interior-1.jpg', 'gallery', 1, 'F-150 Lariat interior'),
(2, '/static/images/gallery/truck-detail-1.jpg', 'gallery', 2, 'F-150 detail work'),
(2, '/static/images/gallery/car-wheels-1.jpg', 'gallery', 3, 'F-150 wheel detail');

-- F-150 Raptor R (ID 3)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(3, '/static/images/gallery/raptor-1.jpg', 'hero', 0, '2024 Ford F-150 Raptor R exterior'),
(3, '/static/images/gallery/pickup-highway-1.jpg', 'gallery', 1, 'Raptor R on the road'),
(3, '/static/images/gallery/truck-detail-1.jpg', 'gallery', 2, 'Raptor R detail work'),
(3, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 3, 'Raptor R being detailed');

-- F-150 Raptor White (ID 4)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(4, '/static/images/gallery/raptor-white-1.jpg', 'hero', 0, '2023 Ford F-150 Raptor exterior'),
(4, '/static/images/gallery/pickup-black-night-1.jpg', 'gallery', 1, 'Raptor night shot'),
(4, '/static/images/gallery/car-wash-1.jpg', 'gallery', 2, 'Raptor wash process'),
(4, '/static/images/gallery/ford-logo-1.jpg', 'gallery', 3, 'Ford emblem detail');

-- F-250 Super Duty (ID 5)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(5, '/static/images/gallery/ford-pickup-beach-1.jpg', 'hero', 0, '2024 Ford F-250 Super Duty'),
(5, '/static/images/gallery/truck-detail-1.jpg', 'gallery', 1, 'F-250 detail work'),
(5, '/static/images/gallery/car-wheels-1.jpg', 'gallery', 2, 'F-250 wheel detail'),
(5, '/static/images/gallery/ford-interior-1.jpg', 'gallery', 3, 'F-250 interior');

-- F-350 Platinum (ID 6)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(6, '/static/images/gallery/ford-single-cab-1.jpg', 'hero', 0, '2023 Ford F-350 Platinum'),
(6, '/static/images/gallery/car-interior-leather-1.jpg', 'gallery', 1, 'F-350 premium interior'),
(6, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 2, 'F-350 detailing process'),
(6, '/static/images/gallery/ford-logo-1.jpg', 'gallery', 3, 'Ford badge detail');

-- Mustang GT Green (ID 7)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(7, '/static/images/gallery/mustang-green-1.jpg', 'hero', 0, '2024 Ford Mustang GT green'),
(7, '/static/images/gallery/mustang-backlight-1.jpg', 'gallery', 1, 'Mustang GT taillights'),
(7, '/static/images/gallery/car-wheels-1.jpg', 'gallery', 2, 'Mustang GT wheel detail'),
(7, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 3, 'Mustang GT detailing');

-- Mustang GT Black (ID 8)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(8, '/static/images/gallery/mustang-black-gt-1.jpg', 'hero', 0, '2023 Ford Mustang GT black'),
(8, '/static/images/gallery/mustang-black-urban-1.jpg', 'gallery', 1, 'Mustang GT urban setting'),
(8, '/static/images/gallery/mustang-shiny-black-1.jpg', 'gallery', 2, 'Mustang GT shine detail'),
(8, '/static/images/gallery/ford-interior-1.jpg', 'gallery', 3, 'Mustang GT interior');

-- Mustang Mach 1 Orange (ID 9)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(9, '/static/images/gallery/mustang-orange-1.jpg', 'hero', 0, '2024 Ford Mustang Mach 1 orange'),
(9, '/static/images/gallery/mustang-backlight-1.jpg', 'gallery', 1, 'Mach 1 taillights'),
(9, '/static/images/gallery/car-wheels-1.jpg', 'gallery', 2, 'Mach 1 wheel detail'),
(9, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 3, 'Mach 1 detailing process');

-- Mustang Shelby GT500 White (ID 10)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(10, '/static/images/gallery/mustang-white-1.jpg', 'hero', 0, '2023 Ford Mustang Shelby GT500'),
(10, '/static/images/gallery/mustang-shelby-snow-1.jpg', 'gallery', 1, 'Shelby GT500 snow photo'),
(10, '/static/images/gallery/car-interior-leather-1.jpg', 'gallery', 2, 'Shelby GT500 interior'),
(10, '/static/images/gallery/ford-logo-1.jpg', 'gallery', 3, 'Shelby emblem');

-- Mustang Boss 302 Vintage (ID 11)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(11, '/static/images/gallery/mustang-vintage-red-1.jpg', 'hero', 0, '1969 Ford Mustang Boss 302'),
(11, '/static/images/gallery/classic-red-pickup-1.jpg', 'gallery', 1, 'Classic car show'),
(11, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 2, 'Vintage detail work'),
(11, '/static/images/gallery/car-wheels-1.jpg', 'gallery', 3, 'Classic wheel restoration');

-- Mustang EcoBoost Red (ID 12)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(12, '/static/images/gallery/mustang-red-1.jpg', 'hero', 0, '2024 Ford Mustang EcoBoost red'),
(12, '/static/images/gallery/mustang-red-2.jpg', 'gallery', 1, 'Mustang EcoBoost profile'),
(12, '/static/images/gallery/mustang-backlight-1.jpg', 'gallery', 2, 'EcoBoost taillights'),
(12, '/static/images/gallery/ford-interior-1.jpg', 'gallery', 3, 'EcoBoost interior');

-- Explorer ST White (ID 13)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(13, '/static/images/gallery/explorer-white-1.jpg', 'hero', 0, '2024 Ford Explorer ST'),
(13, '/static/images/gallery/car-interior-leather-1.jpg', 'gallery', 1, 'Explorer ST interior'),
(13, '/static/images/gallery/car-wheels-1.jpg', 'gallery', 2, 'Explorer ST wheels'),
(13, '/static/images/gallery/car-wash-1.jpg', 'gallery', 3, 'Explorer ST wash');

-- Escape Hybrid (ID 14)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(14, '/static/images/gallery/explorer-white-1.jpg', 'hero', 0, '2024 Ford Escape Hybrid'),
(14, '/static/images/gallery/ford-interior-1.jpg', 'gallery', 1, 'Escape interior'),
(14, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 2, 'Escape detailing'),
(14, '/static/images/gallery/ford-logo-1.jpg', 'gallery', 3, 'Ford emblem');

-- Expedition Limited (ID 15)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(15, '/static/images/gallery/f150-white-1.jpg', 'hero', 0, '2024 Ford Expedition Limited'),
(15, '/static/images/gallery/car-interior-leather-1.jpg', 'gallery', 1, 'Expedition leather interior'),
(15, '/static/images/gallery/car-wheels-1.jpg', 'gallery', 2, 'Expedition wheel detail'),
(15, '/static/images/gallery/ford-interior-1.jpg', 'gallery', 3, 'Expedition dashboard');

-- Bronco Badlands (ID 16)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(16, '/static/images/gallery/bronco-rear-1.jpg', 'hero', 0, '2024 Ford Bronco Badlands'),
(16, '/static/images/gallery/bronco-rear-2.jpg', 'gallery', 1, 'Bronco rear angle'),
(16, '/static/images/gallery/truck-detail-1.jpg', 'gallery', 2, 'Bronco detail work'),
(16, '/static/images/gallery/car-wash-1.jpg', 'gallery', 3, 'Bronco wash');

-- Maverick Lariat (ID 17)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(17, '/static/images/gallery/maverick-desert-1.jpg', 'hero', 0, '2024 Ford Maverick Lariat'),
(17, '/static/images/gallery/pickup-highway-1.jpg', 'gallery', 1, 'Maverick on the road'),
(17, '/static/images/gallery/ford-interior-1.jpg', 'gallery', 2, 'Maverick interior'),
(17, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 3, 'Maverick detailing');

-- F-100 Classic Red (ID 18)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(18, '/static/images/gallery/ford-vintage-red-1.jpg', 'hero', 0, '1978 Ford F-100 classic'),
(18, '/static/images/gallery/classic-red-pickup-1.jpg', 'gallery', 1, 'Classic F-100 side view'),
(18, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 2, 'Vintage detail process'),
(18, '/static/images/gallery/car-wheels-1.jpg', 'gallery', 3, 'Classic wheel restoration');

-- F-250 Heritage Blue (ID 19)
INSERT INTO media (gallery_group_id, url, kind, sort_order, alt_text) VALUES
(19, '/static/images/gallery/ford-pickup-beach-1.jpg', 'hero', 0, '1965 Ford F-250 heritage'),
(19, '/static/images/gallery/ford-single-cab-1.jpg', 'gallery', 1, 'Classic F-250 profile'),
(19, '/static/images/gallery/car-detailing-1.jpg', 'gallery', 2, 'Heritage detail work'),
(19, '/static/images/gallery/ford-logo-1.jpg', 'gallery', 3, 'Vintage Ford emblem');
