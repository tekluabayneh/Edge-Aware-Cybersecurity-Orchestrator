-- name: GetAllNotifications :many
SELECT * FROM notifications WHERE user_id = $1 AND is_read = false ORDER BY id DESC LIMIT 50;

-- name: GetNotificationById :one
SELECT * from notifications WHERE id = $1 AND user_id = $2;

-- name: CreateNotification :exec
INSERT INTO notifications(user_id, title, message, is_read) VALUES ($1, $2, $3, $4);

-- name: UpdateNotificationById :exec
UPDATE notifications SET is_read = true WHERE user_id = $1 AND id = $2;

-- name: UpdateAllNotificationByUserId :exec
UPDATE notifications SET is_read = true WHERE user_id = $1;

