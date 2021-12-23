package query

import (
	"fmt"
	"strings"

	"github.com/u-shylianok/ad-service/domain/model"
)

func BuildAdFilterQuery(filter model.AdFilter) (string, []interface{}) {

	var result string
	args := []interface{}{}
	argID := 1

	var subQuery string
	// Создаем подзапрос для фильтрации по датам
	if !filter.StartDate.IsZero() || !filter.EndDate.IsZero() {

		var sb strings.Builder

		if !filter.StartDate.IsZero() {
			sb.WriteString(fmt.Sprintf("ads.date >= $%d", argID))
			args = append(args, filter.StartDate)
			argID++

			if !filter.EndDate.IsZero() {
				sb.WriteString(" AND ")
			}
		}
		if !filter.EndDate.IsZero() {
			args = append(args, filter.EndDate)
			sb.WriteString(fmt.Sprintf("ads.date <= $%d", argID))
			argID++
		}

		subQuery = fmt.Sprintf("SELECT * FROM ads WHERE %s", sb.String())
	}

	const adsColumns = "ads.id, ads.user_id, ads.name, ads.date, ads.price, ads.description, ads.photo"

	// Фильтрация по остальным критериям
	var filterQuery string
	{
		var sb strings.Builder

		if filter.Username != "" {
			sb.WriteString(fmt.Sprintf(" INNER JOIN users ON ads.user_id = users.id AND username = $%d", argID))
			args = append(args, filter.Username)
			argID++
		}
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
