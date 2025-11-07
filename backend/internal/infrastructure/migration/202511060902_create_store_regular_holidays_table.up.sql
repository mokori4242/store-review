CREATE TABLE IF NOT EXISTS store_regular_holidays (
  id BIGSERIAL PRIMARY KEY,
  store_id BIGINT NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
  day_of_week SMALLINT NOT NULL CHECK (day_of_week BETWEEN 0 AND 6),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_store_regular_holidays_store_id ON store_regular_holidays(store_id);

COMMENT ON TABLE store_regular_holidays IS '店舗の定休日を管理するテーブル';
COMMENT ON COLUMN store_regular_holidays.day_of_week IS '0~6 = 月~日';
