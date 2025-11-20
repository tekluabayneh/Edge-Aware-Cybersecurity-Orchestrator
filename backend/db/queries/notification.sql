-- name: GetAllNotifications:one
SELECT * FROM  notifications LIMIT =50;

-- name: GetNotificationById:one
SELECT * from notifications WHERE id = $1;

-- name: CreateNotification:exec
INSERT INTO notifications(user_id, title, message, is_read) VALUES ($1, $2, $3, $4)

-- name: UpdateNotificationById:exec
UPDATE notifications
SET
    user_id       = COALESCE($1, user_id),
    title = COALESCE($2, title),
    message = COALESCE($3, message),
    is_read = COALESCE($4, is_read)
WHERE id = $5;
