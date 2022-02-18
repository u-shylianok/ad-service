package query

import (
	"fmt"
	"strings"

	"github.com/u-shylianok/ad-service/svc-ads/domain/model"
)

func BuildAdFilterQuery(filter model.AdFilter) (string, []interface{}) {

	var result string
	args := []interface{}{}
	argID := 1

	var subQuery string
	// Создаем подзапрос для фильтрации по датам
	if !filter.StartDate.IsZero() || !filter.EndDate.IsZero() || filter.UserID != 0 {

		var sb strings.Builder

		if !filter.StartDate.IsZero() {
			sb.WriteString(fmt.Sprintf("ads.date >= $%d", argID))
			args = append(args, filter.StartDate)
			argID++
		}

		if !filter.EndDate.IsZero() {
			if sb.Len() != 0 {
				sb.WriteString(" AND ")
			}

			sb.WriteString(fmt.Sprintf("ads.date <= $%d", argID))
			args = append(args, filter.EndDate)
			argID++
		}

		if filter.UserID != 0 {
			if sb.Len() != 0 {
				sb.WriteString(" AND ")
			}

			sb.WriteString(fmt.Sprintf("ads.user_id = $%d", argID))
			args = append(args, filter.UserID)
			argID++
		}

		subQuery = fmt.Sprintf("SELECT * FROM ads WHERE %s", sb.String())
	}

	const adsColumns = "ads.id, ads.user_id, ads.name, ads.date, ads.price, ads.description, ads.photo"

	// Фильтрация по остальным критериям
	var filterQuery string
	{
		var sb strings.Builder

		if len(filter.Tags) > 0 {
			values := []string{}
			for _, tag := range filter.Tags {
				values = append(values, fmt.Sprintf("$%d", argID))
				args = append(args, tag)
				argID++
			}
			sb.WriteString(" INNER JOIN ads_tags ON ads.id = ads_tags.ad_id")
			sb.WriteString(fmt.Sprintf(" INNER JOIN tags ON tags.id = ads_tags.tag_id AND tags.name IN (%s)",
				strings.Join(values, ",")))
			sb.WriteString(fmt.Sprintf(" GROUP BY %s HAVING COUNT(DISTINCT tags.name) = %d",
				adsColumns, len(values)))

		}

		filterQuery = sb.String()
	}

	if subQuery != "" {
		result = fmt.Sprintf("SELECT %s FROM (%s) AS ads %s", adsColumns, subQuery, filterQuery)
	} else {
		result = fmt.Sprintf("SELECT %s FROM ads %s", adsColumns, filterQuery)
	}

	return result, args
}
