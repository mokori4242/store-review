CREATE TABLE IF NOT EXISTS store_payment_methods (
     id BIGSERIAL PRIMARY KEY,
     store_id BIGINT NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
     method SMALLINT NOT NULL, -- 例: 'PayPay', 'Cash', 'RakutenPay'
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE store_payment_methods IS '店舗の支払い方法を管理するテーブル';
COMMENT ON COLUMN store_payment_methods.method IS '0:PayPay,1:Cash,2:RakutenPay,etc...';
