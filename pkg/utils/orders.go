package utils

import (
	"database/sql"
	"inorder/pkg/types"
	"time"
)

func ParseOrderRows(rows *sql.Rows) ([]*types.Order, error) {
	var otpt []*types.Order = make([]*types.Order, 0)
	defer rows.Close()

	if exists := rows.Next(); !exists {
		return otpt, nil
	}
	for {
		var curr types.Order
		var rd types.MYSQLOrder
		err := rows.Scan(&curr.ID, &curr.IssuedBy, &rd.IssuedAt, &curr.Status, &rd.BillableAmount, &curr.TableNo, &rd.Waiter, &rd.PaidAt, &rd.Tip)
		if err != nil {
			return otpt, err
		}
		curr.IssuedAt, err = time.Parse(time.DateTime, string(rd.IssuedAt))
		if err != nil {
			return otpt, err
		}
		if len(rd.PaidAt) != 0 {
			curr.PaidAt, err = time.Parse(time.DateTime, string(rd.PaidAt))
			if err != nil {
				return otpt, err
			}
		}
		if rd.BillableAmount.Valid {
			curr.BillableAmount = float32(rd.BillableAmount.Float64)
		}
		if rd.Waiter.Valid {
			curr.Waiter = types.UserID(rd.Waiter.Int64)
		}
		if rd.Tip.Valid {
			curr.Tip = float32(rd.Tip.Float64)
		}

		otpt = append(otpt, &curr)

		if isNext := rows.Next(); !isNext {
			break
		}
	}
	return otpt, nil
}
