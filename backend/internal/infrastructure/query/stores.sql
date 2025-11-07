-- name: GetListStores :many
SELECT s.id, s.name,
       COALESCE(ARRAY_AGG(DISTINCT srh.day_of_week) FILTER (WHERE srh.day_of_week IS NOT NULL), '{}')::text[] AS regular_holidays,
       COALESCE(ARRAY_AGG(DISTINCT c.name) FILTER (WHERE c.name IS NOT NULL), '{}')::text[] AS category_names,
       COALESCE(ARRAY_AGG(DISTINCT p.method) FILTER (WHERE p.method IS NOT NULL), '{}')::text[] AS payment_methods,
       COALESCE(ARRAY_AGG(DISTINCT w.url) FILTER (WHERE w.url IS NOT NULL), '{}')::text[] AS web_profiles
FROM stores s
         LEFT JOIN store_regular_holidays srh ON srh.store_id = s.id
         LEFT JOIN store_categories sc ON sc.store_id = s.id
         LEFT JOIN store_category_master c ON c.id = sc.category_id
         LEFT JOIN store_payment_methods p ON p.store_id = s.id
         LEFT JOIN store_web_profiles w ON w.store_id = s.id
GROUP BY s.id, s.name;
