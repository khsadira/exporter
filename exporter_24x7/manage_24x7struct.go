package exporter_24x7

func toCorrectFormat(metrics metrics24x7) [8]chartData {

	timeChart := metrics.Data.ChartData[1].LocationResponseTimeChart

	//creat num_struct
	num6 := timeChart[0].Num6
	num18 := timeChart[6].Num18
	num21 := timeChart[5].Num21
	num22 := timeChart[3].Num22
	num25 := timeChart[2].Num25
	num26 := timeChart[1].Num26
	num31 := timeChart[7].Num31
	num46 := timeChart[4].Num46

	var ret [8]chartData

	ret[0].nbStruct = 6
	ret[0].Label = num6.Label
	ret[0].Max = num6.Max[0]
	ret[0].Min = num6.Min[0]
	ret[0].Percentile = num6.Nine5Percentile[0]
	ret[0].Average = num6.Average[0]
	ret[0].Data = num6.ChartData[0][0]

	ret[1].nbStruct = 18
	ret[1].Label = num18.Label
	ret[1].Max = num18.Max[0]
	ret[1].Min = num18.Min[0]
	ret[1].Percentile = num18.Nine5Percentile[0]
	ret[1].Average = num18.Average[0]
	ret[1].Data = num18.ChartData[0][0]

	ret[2].nbStruct = 21
	ret[2].Label = num21.Label
	ret[2].Max = num21.Max[0]
	ret[2].Min = num21.Min[0]
	ret[2].Percentile = num21.Nine5Percentile[0]
	ret[2].Average = num21.Average[0]
	ret[2].Data = num21.ChartData[0][0]

	ret[3].nbStruct = 22
	ret[3].Label = num22.Label
	ret[3].Max = num22.Max[0]
	ret[3].Min = num22.Min[0]
	ret[3].Percentile = num22.Nine5Percentile[0]
	ret[3].Average = num22.Average[0]
	ret[3].Data = num22.ChartData[0][0]

	ret[4].nbStruct = 25
	ret[4].Label = num25.Label
	ret[4].Max = num25.Max[0]
	ret[4].Min = num25.Min[0]
	ret[4].Percentile = num25.Nine5Percentile[0]
	ret[4].Average = num25.Average[0]
	ret[4].Data = num25.ChartData[0][0]

	ret[5].nbStruct = 26
	ret[5].Label = num26.Label
	ret[5].Max = num26.Max[0]
	ret[5].Min = num26.Min[0]
	ret[5].Percentile = num26.Nine5Percentile[0]
	ret[5].Average = num26.Average[0]
	ret[5].Data = num26.ChartData[0][0]

	ret[6].nbStruct = 31
	ret[6].Label = num31.Label
	ret[6].Max = num31.Max[0]
	ret[6].Min = num31.Min[0]
	ret[6].Percentile = num31.Nine5Percentile[0]
	ret[6].Average = num31.Average[0]
	ret[6].Data = num31.ChartData[0][0]

	ret[7].nbStruct = 46
	ret[7].Label = num46.Label
	ret[7].Max = num46.Max[0]
	ret[7].Min = num46.Min[0]
	ret[7].Percentile = num46.Nine5Percentile[0]
	ret[7].Average = num46.Average[0]
	ret[7].Data = num46.ChartData[0][0]
	return ret
}
