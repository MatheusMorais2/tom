ALTER TABLE keywords RENAME COLUMN name TO word;

DELETE FROM categories
WHERE name IN ('auditorium', 'bistro', 'lounge', 'cinema');

