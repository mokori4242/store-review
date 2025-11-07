-- カテゴリーマスターデータの挿入
INSERT INTO store_category_master (name) VALUES
('カフェ'),
('レストラン'),
('居酒屋'),
('ラーメン'),
('イタリアン'),
('和食'),
('中華'),
('ファーストフード'),
('バー'),
('焼肉')
ON CONFLICT (name) DO NOTHING;

-- 店舗データの挿入
INSERT INTO stores (name, address) VALUES
('カフェ ブルームーン', '東京都渋谷区道玄坂1-2-3'),
('イタリアン トラットリア', '東京都港区六本木3-4-5'),
('居酒屋 さくら', '東京都新宿区歌舞伎町2-3-4'),
('ラーメン 一番', '東京都豊島区池袋1-2-3'),
('焼肉 大将', '東京都品川区大井町4-5-6'),
('和食 季節', '東京都中央区銀座5-6-7'),
('中華料理 龍門', '東京都台東区上野3-4-5'),
('カフェ サンシャイン', '東京都世田谷区三軒茶屋2-3-4'),
('バー ミッドナイト', '東京都渋谷区恵比寿1-2-3'),
('レストラン オーシャンビュー', '東京都港区お台場1-2-3');

-- 店舗カテゴリーの関連付け
INSERT INTO store_categories (store_id, category_id)
SELECT s.id, c.id
FROM stores s
CROSS JOIN store_category_master c
WHERE
    (s.name = 'カフェ ブルームーン' AND c.name = 'カフェ') OR
    (s.name = 'イタリアン トラットリア' AND c.name = 'イタリアン') OR
    (s.name = '居酒屋 さくら' AND c.name = '居酒屋') OR
    (s.name = 'ラーメン 一番' AND c.name = 'ラーメン') OR
    (s.name = '焼肉 大将' AND c.name = '焼肉') OR
    (s.name = '和食 季節' AND c.name = '和食') OR
    (s.name = '中華料理 龍門' AND c.name = '中華') OR
    (s.name = 'カフェ サンシャイン' AND c.name = 'カフェ') OR
    (s.name = 'バー ミッドナイト' AND c.name = 'バー') OR
    (s.name = 'レストラン オーシャンビュー' AND c.name = 'レストラン');

-- 営業時間の挿入
INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '10:00:00', '22:00:00'
FROM stores
WHERE name = 'カフェ ブルームーン';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '11:00:00', '23:00:00'
FROM stores
WHERE name = 'イタリアン トラットリア';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '17:00:00', '02:00:00'
FROM stores
WHERE name = '居酒屋 さくら';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '11:00:00', '23:00:00'
FROM stores
WHERE name = 'ラーメン 一番';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '17:00:00', '00:00:00'
FROM stores
WHERE name = '焼肉 大将';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '11:30:00', '22:00:00'
FROM stores
WHERE name = '和食 季節';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '11:00:00', '22:00:00'
FROM stores
WHERE name = '中華料理 龍門';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '08:00:00', '20:00:00'
FROM stores
WHERE name = 'カフェ サンシャイン';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '18:00:00', '03:00:00'
FROM stores
WHERE name = 'バー ミッドナイト';

INSERT INTO store_business_hours (store_id, business_status, open_time, close_time)
SELECT id, 2, '11:00:00', '21:00:00'
FROM stores
WHERE name = 'レストラン オーシャンビュー';

-- 定休日の挿入 (0~6 = 月~日)
-- カフェ ブルームーン: 月曜定休
INSERT INTO store_regular_holidays (store_id, day_of_week)
SELECT id, 0 FROM stores WHERE name = 'カフェ ブルームーン';

-- イタリアン トラットリア: 水曜定休
INSERT INTO store_regular_holidays (store_id, day_of_week)
SELECT id, 2 FROM stores WHERE name = 'イタリアン トラットリア';

-- 和食 季節: 日曜定休
INSERT INTO store_regular_holidays (store_id, day_of_week)
SELECT id, 6 FROM stores WHERE name = '和食 季節';

-- Webプロフィールの挿入 (0:website,1:x,2:instagram,3:tiktok,4:facebook)
INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 0, 'https://bluemoon-cafe.example.com'
FROM stores WHERE name = 'カフェ ブルームーン';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 2, 'https://instagram.com/bluemoon_cafe'
FROM stores WHERE name = 'カフェ ブルームーン';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 0, 'https://trattoria-italian.example.com'
FROM stores WHERE name = 'イタリアン トラットリア';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 1, 'https://x.com/sakura_izakaya'
FROM stores WHERE name = '居酒屋 さくら';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 0, 'https://ramen-ichiban.example.com'
FROM stores WHERE name = 'ラーメン 一番';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 4, 'https://facebook.com/yakiniku.taisho'
FROM stores WHERE name = '焼肉 大将';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 0, 'https://washoku-kisetsu.example.com'
FROM stores WHERE name = '和食 季節';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 2, 'https://instagram.com/ryumon_chinese'
FROM stores WHERE name = '中華料理 龍門';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 2, 'https://instagram.com/sunshine_cafe'
FROM stores WHERE name = 'カフェ サンシャイン';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 0, 'https://midnight-bar.example.com'
FROM stores WHERE name = 'バー ミッドナイト';

INSERT INTO store_web_profiles (store_id, platform, url)
SELECT id, 0, 'https://oceanview-restaurant.example.com'
FROM stores WHERE name = 'レストラン オーシャンビュー';

-- 支払い方法の挿入 (0:PayPay,1:Cash,2:RakutenPay)
-- カフェ ブルームーン: PayPay, Cash
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 0 FROM stores WHERE name = 'カフェ ブルームーン'
UNION ALL
SELECT id, 1 FROM stores WHERE name = 'カフェ ブルームーン';

-- イタリアン トラットリア: すべて対応
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 0 FROM stores WHERE name = 'イタリアン トラットリア'
UNION ALL
SELECT id, 1 FROM stores WHERE name = 'イタリアン トラットリア'
UNION ALL
SELECT id, 2 FROM stores WHERE name = 'イタリアン トラットリア';

-- 居酒屋 さくら: Cash, PayPay
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 1 FROM stores WHERE name = '居酒屋 さくら'
UNION ALL
SELECT id, 0 FROM stores WHERE name = '居酒屋 さくら';

-- ラーメン 一番: Cash のみ
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 1 FROM stores WHERE name = 'ラーメン 一番';

-- 焼肉 大将: すべて対応
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 0 FROM stores WHERE name = '焼肉 大将'
UNION ALL
SELECT id, 1 FROM stores WHERE name = '焼肉 大将'
UNION ALL
SELECT id, 2 FROM stores WHERE name = '焼肉 大将';

-- 和食 季節: Cash, RakutenPay
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 1 FROM stores WHERE name = '和食 季節'
UNION ALL
SELECT id, 2 FROM stores WHERE name = '和食 季節';

-- 中華料理 龍門: PayPay, Cash
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 0 FROM stores WHERE name = '中華料理 龍門'
UNION ALL
SELECT id, 1 FROM stores WHERE name = '中華料理 龍門';

-- カフェ サンシャイン: すべて対応
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 0 FROM stores WHERE name = 'カフェ サンシャイン'
UNION ALL
SELECT id, 1 FROM stores WHERE name = 'カフェ サンシャイン'
UNION ALL
SELECT id, 2 FROM stores WHERE name = 'カフェ サンシャイン';

-- バー ミッドナイト: Cash のみ
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 1 FROM stores WHERE name = 'バー ミッドナイト';

-- レストラン オーシャンビュー: すべて対応
INSERT INTO store_payment_methods (store_id, method)
SELECT id, 0 FROM stores WHERE name = 'レストラン オーシャンビュー'
UNION ALL
SELECT id, 1 FROM stores WHERE name = 'レストラン オーシャンビュー'
UNION ALL
SELECT id, 2 FROM stores WHERE name = 'レストラン オーシャンビュー';

-- 来店統計の挿入
INSERT INTO store_visit_stats (store_id, total_visit_count, last_visited_at)
SELECT id, 125, '2025-11-05 14:30:00'::timestamp FROM stores WHERE name = 'カフェ ブルームーン'
UNION ALL
SELECT id, 89, '2025-11-06 19:45:00'::timestamp FROM stores WHERE name = 'イタリアン トラットリア'
UNION ALL
SELECT id, 203, '2025-11-06 22:15:00'::timestamp FROM stores WHERE name = '居酒屋 さくら'
UNION ALL
SELECT id, 156, '2025-11-07 12:30:00'::timestamp FROM stores WHERE name = 'ラーメン 一番'
UNION ALL
SELECT id, 78, '2025-11-04 20:00:00'::timestamp FROM stores WHERE name = '焼肉 大将'
UNION ALL
SELECT id, 45, '2025-11-03 18:30:00'::timestamp FROM stores WHERE name = '和食 季節'
UNION ALL
SELECT id, 112, '2025-11-06 13:00:00'::timestamp FROM stores WHERE name = '中華料理 龍門'
UNION ALL
SELECT id, 167, '2025-11-07 10:15:00'::timestamp FROM stores WHERE name = 'カフェ サンシャイン'
UNION ALL
SELECT id, 34, '2025-11-05 23:45:00'::timestamp FROM stores WHERE name = 'バー ミッドナイト'
UNION ALL
SELECT id, 92, '2025-11-06 20:30:00'::timestamp FROM stores WHERE name = 'レストラン オーシャンビュー';
