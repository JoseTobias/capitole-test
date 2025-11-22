-- Add category_id column to products table
ALTER TABLE products ADD COLUMN IF NOT EXISTS category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL;

-- Insert the 3 categories
INSERT INTO categories (code, name) VALUES
('CLOTHING', 'Clothing'),
('SHOES', 'Shoes'),
('ACCESSORIES', 'Accessories')
ON CONFLICT (code) DO NOTHING;

-- Update products to link them to their respective categories
-- PROD001, PROD004, PROD007 belong to "Clothing"
UPDATE products SET category_id = (SELECT id FROM categories WHERE code = 'CLOTHING') WHERE code IN ('PROD001', 'PROD004', 'PROD007');

-- PROD002, PROD006 belong to "Shoes"
UPDATE products SET category_id = (SELECT id FROM categories WHERE code = 'SHOES') WHERE code IN ('PROD002', 'PROD006');

-- PROD003, PROD005, PROD008 belong to "Accessories"
UPDATE products SET category_id = (SELECT id FROM categories WHERE code = 'ACCESSORIES') WHERE code IN ('PROD003', 'PROD005', 'PROD008');

