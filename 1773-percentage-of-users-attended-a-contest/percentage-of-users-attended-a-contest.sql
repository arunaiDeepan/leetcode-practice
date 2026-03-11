WITH total_users AS (
    SELECT COUNT(*) AS cnt FROM Users
) 
SELECT
    r.contest_id,
    ROUND
    (
        COALESCE(
            COUNT(DISTINCT r.user_id) * 100.0 / t.cnt,
            0
        ),
        2
    ) AS percentage 
FROM Register r
CROSS JOIN total_users t
GROUP BY r.contest_id, t.cnt
ORDER BY percentage DESC, r.contest_id ASC;