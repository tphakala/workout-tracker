package timeago

func kaLocale(_ float64, index int) (ago string, in string) {
	var res = [][]string{
		{"ამ წამს", "ახლა"},
		{"%d წამის წინ", "%d წამში"},
		{"1 წუთის წინ", "1 წუთში"},
		{"%d წუთის წინ", "%d წუთში"},
		{"1 საათის წინ", "1 საათში"},
		{"%d საათის წინ", "%d საათში"},
		{"1 დღის წინ", "1 დღეში"},
		{"%d დღის წინ", "%d დღეში"},
		{"1 კვირის წინ", "1 კვირაში"},
		{"%d კვირის წინ", "%d კვირაში"},
		{"1 თვის წინ", "1 თვეში"},
		{"%d თვის წინ", "%d თვეში"},
		{"1 წლის წინ", "1 წელიწადში"},
		{"%d წლის წინ", "%d წელიწადში"},
	}[index]
	return res[0], res[1]
}