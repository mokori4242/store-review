CREATE TABLE IF NOT EXISTS store_visit_stats (
     store_id BIGINT PRIMARY KEY REFERENCES stores(id) ON DELETE CASCADE,
     total_visit_count INTEGER DEFAULT 0,
     last_visited_at TIMESTAMP,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_store_visit_stats_store_id ON store_visit_stats(store_id);

COMMENT ON TABLE store_visit_stats IS '店舗の来店統計を管理するテーブル';
COMMENT ON COLUMN store_visit_stats.total_visit_count IS '合計訪問数';
COMMENT ON COLUMN store_visit_stats.last_visited_at IS '最終訪問日';
