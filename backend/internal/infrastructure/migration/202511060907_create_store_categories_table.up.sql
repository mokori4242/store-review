CREATE TABLE IF NOT EXISTS store_categories (
   id BIGSERIAL PRIMARY KEY,
   store_id BIGINT NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
   category_id BIGINT NOT NULL REFERENCES store_category_master(id) ON DELETE CASCADE,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   UNIQUE (store_id, category_id)
);

COMMENT ON TABLE store_categories IS '店舗と店舗カテゴリーの中間テーブル';
