IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'books' LIMIT 1) THEN
    CREATE TABLE books (
        id BIGINT UNSIGNED PRIMARY KEY,
        title VARCHAR(191),
        cover VARCHAR(191),
        type TINYINT UNSIGNED
    );

    INSERT INTO books (id, title, cover, type) VALUES
        (1, 'To Kill a Mockingbird', 'Cover 1', 1),
        (2, '1984', 'Cover 2', 2),
        (3, 'The Great Gatsby', 'Cover 3', 1),
        (4, 'Moby-Dick', 'Cover 4', 2),
        (5, 'Pride and Prejudice', 'Cover 5', 1);
END IF;
