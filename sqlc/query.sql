-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = ? LIMIT 1;

-- name: FindAuthorByName :one
SELECT * FROM authors
WHERE name = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name
) VALUES (
  ?
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors
set name = ?
WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;

-- name: DeleteAuthorByName :exec
DELETE FROM authors
WHERE name = ?;

-- name: CreateQuote :one
INSERT into quotes (
  content,
  author_id
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetQuoteAndAuthor :one
SELECT * FROM quotes
INNER JOIN authors
  ON quotes.author_id = authors.id
WHERE quotes.id = ?
LIMIT 1;


-- name: CreateQuoteAuthor :one
-- BEGIN TRANSACTION
--   INSERT INTO authors (name)
--   SELECT ?
--   WHERE NOT EXISTS (SELECT 1 FROM authors WHERE name = ?);
--   WITH quote_author_id AS (
--     SELECT id FROM authors WHERE name = ?
--   )
--   RETURNING *;
--   INSERT INTO quotes(content, author_id)
--   VALUES(
--     ?,
--     quote_author_id
--   )
--   RETURNING *;
-- END TRANSACTION;
