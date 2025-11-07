-- ユーザーデータの挿入(passwordはpasswordをハッシュ化させたもの)
INSERT INTO users(nickname, email, password) VALUES
('テスト1', 'test1@example.com', '$2a$10$96yvIhyrN/EsONWFBEyRne8Fc2YA3yN5HqAQjgQ5l8kAXwOnuHHOe'),
('テスト2', 'test2@example.com', '$2a$10$96yvIhyrN/EsONWFBEyRne8Fc2YA3yN5HqAQjgQ5l8kAXwOnuHHOe'),
('テスト3', 'test3@example.com', '$2a$10$96yvIhyrN/EsONWFBEyRne8Fc2YA3yN5HqAQjgQ5l8kAXwOnuHHOe');