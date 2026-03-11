SELECT 
    product_id,
    ROUND(
        COALESCE(
            (revenue)::numeric / 
            NULLIF(no_units,0),
            0
        ),
        2
    ) AS average_price
FROM (
    SELECT 
        p.product_id,
        SUM(p.price * u.units) AS revenue,
        SUM(u.units) AS no_units
    FROM Prices p
    LEFT JOIN UnitsSold u 
        ON p.product_id = u.product_id
        AND u.purchase_date BETWEEN p.start_date AND p.end_date
    GROUP BY p.product_id
) sub_t;