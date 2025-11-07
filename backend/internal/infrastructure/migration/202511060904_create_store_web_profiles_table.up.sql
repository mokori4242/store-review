CREATE TABLE IF NOT EXISTS store_web_profiles (
  id BIGSERIAL PRIMARY KEY,
  store_id BIGINT NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
  platform SMALLINT NOT NULL,
  url TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_store_web_profiles_store_id ON store_web_profiles(store_id);

COMMENT ON TABLE store_web_profiles IS '店舗のwebを管理するテーブル';
COMMENT ON COLUMN store_web_profiles.platform IS '0:website,1:x,2:instagram,3:tiktok,4:facebook,etc...';
