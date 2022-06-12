use snippetbox;
z CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);
-- Add an index on the created column. CREATE INDEX idx_snippets_created ON snippets(created);
CREATE INDEX idx_snippets_created ON snippets(created);
INSERT INTO snippets (title, content, created, expires)
VALUES (
        'An old silent pond',
        'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.',
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
    );
INSERT INTO snippets (title, content, created, expires) 
VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
    );


