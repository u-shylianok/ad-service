package query

import (
	"fmt"
	"strings"

	"github.com/u-shylianok/ad-service/internal/model"
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
			sb.WriteString(fmt.Sprintf(" INNER JOIN tags ON tags.id = ads_tags.tag_id AND tags.name IN (%s)", strings.Join(values, ",")))
			sb.WriteString(fmt.Sprintf(" GROUP BY %s HAVING COUNT(DISTINCT tags.name) = %d", adsColumns, len(values)))

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

// Old version
func BuildAdFilterQueryOLD(filter model.AdFilter) (string, []interface{}) {

	args := []interface{}{}
	argID := 1

	var dateFilter string
	{
		var sb strings.Builder
		if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
			if !filter.StartDate.IsZero() {
				sb.WriteString(fmt.Sprintf(" ads.date >= $%d", argID))
				args = append(args, filter.StartDate)
				argID++

				if !filter.EndDate.IsZero() {
					sb.WriteString(" AND")
				}
			}
			if !filter.EndDate.IsZero() {
				args = append(args, filter.EndDate)
				sb.WriteString(fmt.Sprintf(" ads.date <= $%d", argID))
				argID++
			}
		}
		dateFilter = sb.String()
	}
	var isDateFilterAdded bool

	var filterQuery string
	{
		var sb strings.Builder

		if filter.Username != "" {
			sb.WriteString(fmt.Sprintf(" INNER JOIN users ON ads.user_id = users.id AND username = $%d", argID))
			args = append(args, filter.Username)
			argID++
			if dateFilter != "" && !isDateFilterAdded {
				sb.WriteString(" AND")
				sb.WriteString(dateFilter)
				isDateFilterAdded = true
			}
		}
		if len(filter.Tags) > 0 {
			sb.WriteString(" INNER JOIN ads_tags ON ads.id = ads_tags.ad_id ")
			if dateFilter != "" && !isDateFilterAdded {
				sb.WriteString(" AND")
				sb.WriteString(dateFilter)
				isDateFilterAdded = true
			}
			values := []string{}
			for _, tag := range filter.Tags {
				values = append(values, fmt.Sprintf("$%d", argID))
				args = append(args, tag)
				argID++
			}
			sb.WriteString(fmt.Sprintf(" INNER JOIN tags ON tags.id = ads_tags.tag_id AND tags.name IN (%s)", strings.Join(values, ",")))
			sb.WriteString(fmt.Sprintf(" GROUP BY ads.* HAVING COUNT(DISTINCT tags.name) = %d", len(values)))
			args = append(args, strings.Join(filter.Tags, ","))
			argID++
		}
		if dateFilter != "" && !isDateFilterAdded {
			sb.WriteString(" WHERE")
			sb.WriteString(dateFilter)
			isDateFilterAdded = true
		}
		filterQuery = sb.String()
	}
	return filterQuery, args
}

// // query example 1
// SELECT ads.* FROM ads
// INNER JOIN users ON ads.user_id = users.id AND username='test'
// INNER JOIN ads_tags ON ads.id = ads_tags.ad_id
// INNER JOIN tags ON tags.id = ads_tags.tag_id AND tags.name = ALL ('ВАЖНОЕ')
// WHERE ads.date > '2021-10-9' AND ads.date < '2021-10-25'

// // query example 1
// SELECT ads.* FROM ads
// INNER JOIN users ON ads.user_id = users.id AND username='test'
// INNER JOIN ads_tags ON ads.id = ads_tags.ad_id
// INNER JOIN tags ON tags.id = ads_tags.tag_id AND tags.name = 'ВАЖНОЕ'
// WHERE ads.date > '2021-10-9' AND ads.date < '2021-10-25'
