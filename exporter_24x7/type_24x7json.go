package exporter_24x7

type chartData struct {
	nbStruct   int
	Label      string
	Max        float64
	Min        float64
	Percentile float64
	Average    float64
	Data       interface{}
}

type metrics24x7 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		TableData struct {
		} `json:"table_data"`
		Info struct {
			FormattedEndTime       string `json:"formatted_end_time"`
			MonitorType            string `json:"monitor_type"`
			ResourceID             string `json:"resource_id"`
			ResourceTypeName       string `json:"resource_type_name"`
			PeriodName             string `json:"period_name"`
			GeneratedTime          string `json:"generated_time"`
			MetricAggregationName  string `json:"metric_aggregation_name"`
			ReportName             string `json:"report_name"`
			EndTime                string `json:"end_time"`
			MetricAggregation      int    `json:"metric_aggregation"`
			StartTime              string `json:"start_time"`
			SegmentType            int    `json:"segment_type"`
			ReportType             int    `json:"report_type"`
			Period                 int    `json:"period"`
			ResourceName           string `json:"resource_name"`
			SegmentTypeName        string `json:"segment_type_name"`
			FormattedStartTime     string `json:"formatted_start_time"`
			FormattedGeneratedTime string `json:"formatted_generated_time"`
			ResourceType           int    `json:"resource_type"`
		} `json:"info"`
		ChartData []struct {
			ResponseTimeReportChart   []interface{} `json:"ResponseTimeReportChart,omitempty"`
			LocationResponseTimeChart []struct {
				Num6 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"6,omitempty"`
				Num18 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"18,omitempty"`
				Num21 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"21,omitempty"`
				Num22 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"22,omitempty"`
				Num25 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"25,omitempty"`
				Num26 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"26,omitempty"`
				Num31 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"31,omitempty"`
				Num46 struct {
					Max             []float64       `json:"max"`
					Label           string          `json:"label"`
					Min             []float64       `json:"min"`
					Nine5Percentile []float64       `json:"95_percentile"`
					Average         []float64       `json:"average"`
					ChartData       [][]interface{} `json:"chart_data"`
				} `json:"46,omitempty"`
			} `json:"LocationResponseTimeChart,omitempty"`
		} `json:"chart_data"`
	} `json:"data"`
}
