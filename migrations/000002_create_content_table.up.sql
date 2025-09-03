CREATE TABLE tb_content (
  id SERIAL PRIMARY KEY,
  video_title VARCHAR(255) NOT NULL,
  video_url TEXT NOT NULL,
  comment TEXT,
  rating INT,
  thumbnail_url TEXT,
  creator_name VARCHAR(255),
  user_id INT NOT NULL REFERENCES tb_user(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP DEFAULT NULL,
  deleted_at TIMESTAMP DEFAULT NULL
);
