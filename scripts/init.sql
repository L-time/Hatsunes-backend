-- 创建user表
CREATE TABLE "user" (
                        user_id SERIAL PRIMARY KEY,
                        user_password VARCHAR(255) NOT NULL
);

-- 创建personal_info表
CREATE TABLE personal_info (
                               user_id INTEGER NOT NULL,
                               user_name VARCHAR(255) NOT NULL,
                               gender INTEGER, -- 性别可以是 'Male', 'Female', 'Other' 等
                               birthday DATE,
                               signature TEXT, -- 个性签名
                               email VARCHAR(255),
                               avatar BYTEA, -- 头像，可以保存图像的二进制数据
                               FOREIGN KEY (user_id) REFERENCES "user" (user_id) ON DELETE CASCADE
);

-- 创建competence表
CREATE TABLE competence (
                            user_id INTEGER NOT NULL,
                            competence INTEGER NOT NULL,
                            FOREIGN KEY (user_id) REFERENCES "user" (user_id) ON DELETE CASCADE
);

-- 创建post表
CREATE TABLE post (
                      post_id SERIAL PRIMARY KEY,
                      user_id INTEGER NOT NULL,
                      date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      title VARCHAR(255) NOT NULL,
                      details TEXT,
                      FOREIGN KEY (user_id) REFERENCES "user" (user_id) ON DELETE CASCADE
);

-- 创建comment表
CREATE TABLE comment (
                         comment_id SERIAL PRIMARY KEY,
                         user_id INTEGER NOT NULL,
                         post_id INTEGER NOT NULL,
                         reply_id INTEGER,
                         date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         details TEXT,
                         FOREIGN KEY (user_id) REFERENCES "user" (user_id) ON DELETE CASCADE,
                         FOREIGN KEY (post_id) REFERENCES post (post_id) ON DELETE CASCADE,
                         FOREIGN KEY (reply_id) REFERENCES comment (comment_id) ON DELETE CASCADE
);
