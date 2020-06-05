package main

import (
	"log"
	"net/http"
	"os"

	"github.com/xiaowj/go-echarts/charts"
)

func barBase() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	return bar
}

func barTitle() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-标题", Subtitle: "我是副标题，相对来讲我会长一点", Right: "40%"},
		charts.LegendOpts{Right: "80%"},
		charts.ToolboxOpts{Show: true},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	return bar
}

func barShowLabel() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-显示 Label"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
	return bar
}

func barXYName() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-XY 轴名称"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
	bar.SetGlobalOptions(charts.XAxisOpts{Name: "商品名称"}, charts.YAxisOpts{Name: "销售量"})
	return bar
}

func barColor() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-设置系列颜色"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt(), charts.ColorOpts{"lightblue"}).
		AddYAxis("商家B", randInt(), charts.ColorOpts{"pink"})
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
	// 或者可以这样设置
	//bar.SetGlobalOptions(charts.ColorOpts{"lightblue", "pink"})
	return bar
}

func barSplitLine() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-显示分割线"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetGlobalOptions(charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}})
	return bar
}

func barGap() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-调整 bar 距离"})
	bar.AddXAxis(nameItems).AddYAxis("商家A", randInt())
	bar.SetSeriesOptions(charts.BarOpts{BarCategoryGap: "70%"})
	return bar
}

func barYAxis() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-Y 轴格式"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetGlobalOptions(charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}})
	return bar
}

func barMultiYAxis() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-多 Y 轴"},
		charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt(), charts.BarOpts{YAxisIndex: 0}).
		AddYAxis("商家B", randInt(), charts.BarOpts{YAxisIndex: 1})
	bar.ExtendYAxis(charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/月"}})
	return bar
}

func barMultiXAxis() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-多 X 轴"},
		charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt(), charts.BarOpts{XAxisIndex: 0}).
		AddYAxis("商家B", randInt(), charts.BarOpts{XAxisIndex: 1})
	bar.ExtendXAxis(charts.XAxisOpts{Data: foodItems})
	return bar
}

func barDataZoom() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-DataZoom"},
		charts.ToolboxOpts{Show: true},
		charts.DataZoomOpts{XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	return bar
}

func barReverse() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-翻转 XY 轴"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.XYReversal()
	return bar
}

func barStack() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-堆叠效果"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt(), charts.BarOpts{Stack: "stack"}).
		AddYAxis("商家B", randInt(), charts.BarOpts{Stack: "stack"})
	return bar

}

func barMark() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-标记线&标记点"},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetSeriesOptions(
		charts.MLNameTypeItem{Name: "最大值", Type: "max"},
		charts.MLStyleOpts{SymbolSize: 20, Label: charts.LabelTextOpts{Show: true}},
		charts.MPNameTypeItem{Name: "最大值点", Type: "max"},
		charts.MPStyleOpts{Label: charts.LabelTextOpts{Show: true}},
	)
	return bar
}

func barMarkCustom() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-自定义标记+主题"},
		charts.InitOpts{PageTitle: "Awesome", Theme: charts.ThemeType.Macarons},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt(),
			charts.MLNameYAxisItem{Name: "Y 值为 5", YAxis: 20},
			charts.MLNameCoordItem{Name: "自定义标记线",
				Coord0: []interface{}{"衬衫", 10}, Coord1: []interface{}{"羊毛衫", 40}},
			charts.MLStyleOpts{SymbolSize: 20, Label: charts.LabelTextOpts{Show: true}},
		)
	return bar
}

func barSize() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-画布大小"},
		charts.InitOpts{Width: "600px", Height: "400px"},
		charts.ToolboxOpts{Show: true},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	return bar
}

func barHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("bar")...)
	page.Add(
		barBase(),
		barTitle(),
		barShowLabel(),
		barXYName(),
		barColor(),
		barSplitLine(),
		barGap(),
		barYAxis(),
		barMultiYAxis(),
		barMultiXAxis(),
		barDataZoom(),
		barReverse(),
		barStack(),
		barMark(),
		barMarkCustom(),
		barSize(),
	)
	f, err := os.Create(getRenderPath("bar.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
