-- name: IncrementMeetingCount :exec
INSERT INTO people (name, count)
VALUES (?, 1)
ON DUPLICATE KEY UPDATE
count = count + 1;

-- name: GetMeetingCount :one
SELECT count
FROM people
WHERE name = ?;