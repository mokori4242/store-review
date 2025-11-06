CREATE TABLE IF NOT EXISTS store_business_hours (
    id BIGSERIAL PRIMARY KEY,
    store_id BIGINT NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    business_status SMALLINT NOT NULL,
    open_time TIME,
    close_time TIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE store_business_hours IS '店舗の営業時間を管理するテーブル';
COMMENT ON COLUMN store_business_hours.business_status IS '0:午前のみ,1:午後のみ,2:通し,3:休憩';
