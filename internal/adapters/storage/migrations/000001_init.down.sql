DROP INDEX IF EXISTS idx_post_keywords_keyword_id;
DROP INDEX IF EXISTS idx_post_keywords_post_id;
DROP INDEX IF EXISTS idx_posts_category_id;

DROP TABLE IF EXISTS post_keywords;
DROP TABLE IF EXISTS keywords;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS categories;

DROP EXTENSION IF EXISTS "pgcrypto";
