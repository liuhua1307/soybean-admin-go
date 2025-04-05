package location

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"os"
	"soybean-admin-go/config"
	"soybean-admin-go/db/gen"
	"soybean-admin-go/utils/log"
	"strconv"
)

var LocMap = map[string]Coordinate{}

type Coordinate struct {
	Longitude float64
	Latitude  float64
}

func init() {
	file, err := os.Open("orders_with_address.csv")
	if err != nil {
		panic("打开 CSV 文件失败: " + err.Error())
	}
	defer file.Close()

	// 创建 CSV 阅读器
	reader := csv.NewReader(file)

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		panic("读取 CSV 文件失败: " + err.Error())
	}

	// 定义列名到索引的映射
	columns := make(map[string]int)
	for i, header := range records[0] {
		columns[header] = i
	}

	// 检查必需列是否存在
	requiredColumns := []string{"detail_address", "longitude", "latitude"}
	for _, col := range requiredColumns {
		if _, exists := columns[col]; !exists {
			panic(fmt.Sprintf("CSV 文件缺少必需的列: %s", col))
		}
	}

	// 初始化映射
	LocMap = make(map[string]Coordinate)

	// 遍历数据行（跳过表头）
	for _, record := range records[1:] {
		// 提取详细地址
		detailAddress := record[columns["detail_address"]]

		// 解析经纬度
		lon, err := strconv.ParseFloat(record[columns["longitude"]], 64)
		if err != nil {
			fmt.Printf("解析经度失败: %s\n", err)
			continue
		}
		lat, err := strconv.ParseFloat(record[columns["latitude"]], 64)
		if err != nil {
			fmt.Printf("解析纬度失败: %s\n", err)
			continue
		}

		// 添加到映射（相同地址会被覆盖）
		LocMap[detailAddress] = Coordinate{Longitude: lon, Latitude: lat}
	}
	fmt.Println(LocMap)

}

func GotLocation(ctx *gin.Context) {
	var (
		loc = [][]Coordinate{
			{
				{113.393116, 23.039404},
				{104.232717, 33.479594},
				{104.430758, 31.003392},
				{103.927263, 31.193496},
				{103.438729, 34.965523},
				{104.393699, 34.948906},
				{103.970505, 33.586967},
				{105.157181, 32.606498},
				{105.784729, 32.705778},
				{105.573431, 31.369635},
				{102.088984, 34.463923},
				{102.802166, 32.542946},
				{102.555499, 33.378795},
				{104.087572, 34.508873},
				{104.323051, 34.496253},
			},
			{
				{113.393116, 23.039404},
				{113.018329, 33.929822},
				{114.429896, 36.813588},
				{114.048155, 34.369034},
				{115.873625, 33.287552},
				{113.516811, 32.820886},
				{115.383027, 34.546204},
				{113.1608, 32.853814},
				{114.431543, 32.581248},
				{113.50135, 34.494667},
				{112.295128, 35.835092},
				{115.315797, 32.642001},
				{113.700049, 35.251709},
				{113.59653, 36.63738},
				{111.340934, 35.338143},
				{112.4973, 36.032048},
				{114.884567, 34.147756},
			},
			{
				{113.393116, 23.039404},
				{109.573229, 25.483631},
				{115.377363, 24.863425},
				{115.965296, 24.915374},
				{111.791596, 26.182979},
				{112.347575, 27.088419},
				{109.016579, 25.302082},
				{111.73243, 24.133888},
				{108.846107, 25.06465},
			},
			{
				{113.393116, 23.039404},
				{102.392239, 26.96934},
				{102.400643, 27.03194},
				{102.623426, 25.913311},
				{103.642157, 29.511465},
				{102.254306, 26.737535},
				{103.795522, 27.324707},
				{103.017603, 26.517472},
				{104.53349, 29.69273},
				{105.904493, 29.348341},
				{103.31803, 25.118059},
				{104.611754, 24.823445},
				{105.549822, 28.493425},
			},
			{
				{113.393116, 23.039404},
				{106.289195, 34.702739},
				{105.49052, 34.589967},
				{108.715635, 34.618736},
				{108.855549, 35.220372},
				{108.948912, 35.429555},
				{106.150165, 33.868653},
				{108.413586, 33.259111},
				{107.623473, 34.665121},
				{106.569021, 32.850985},
				{107.282801, 34.032081},
				{106.445613, 35.475546},
			},
			{
				{113.393116, 23.039404},
				{115.283939, 39.89803},
				{113.880829, 37.644576},
				{114.782624, 40.156078},
				{113.047502, 40.253311},
			},
			{
				{113.393116, 23.039404},
				{115.543556, 31.094238},
				{113.908427, 31.698098},
				{113.671556, 30.015452},
				{115.641066, 27.983804},
				{113.600775, 30.28966},
				{113.914715, 27.683904},
				{115.596809, 29.875854},
				{112.276866, 29.790961},
				{115.247318, 30.161926},
				{113.279697, 29.802188},
				{113.576533, 30.528186},
			},
			{
				{113.393116, 23.039404},
				{109.636162, 29.232887},
				{109.384547, 29.40236},
				{109.297849, 28.197519},
				{110.525547, 29.512873},
				{110.04536, 31.214203},
				{108.795578, 29.961139},
				{111.648624, 28.960358},
			},
			{
				{113.393116, 23.039404},
				{109.76065, 37.843845},
				{108.467335, 38.385428},
				{110.833578, 39.175419},
				{110.737582, 38.292025},
				{109.873364, 40.065333},
				{107.662772, 40.656805},
				{111.463175, 38.260045},
				{108.418025, 39.932767},
			},
			{
				{113.393116, 23.039404},
				{104.181535, 38.325001},
				{105.505058, 36.242676},
				{105.364783, 40.080729},
				{106.118362, 38.087807},
				{106.942945, 37.314161},
				{104.904663, 37.27945},
				{102.230568, 36.545403},
				{102.067912, 37.51353},
				{104.840022, 38.337916},
			},
		}
		resOr  []Coordinate
		resp   [][]float64
		orders = gen.Q.Order
	)
	i, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		config.Logger.Error("Failed to parse id", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	first, err := orders.WithContext(ctx).Preload(orders.CustomerInfo).Where(orders.ID.Eq(i)).First()
	if err != nil {
		config.Logger.Error("Failed to get order", log.Field{
			Key:   "error",
			Value: err,
		})
		ctx.JSON(500, gin.H{"error": "Failed to get order"})
		return
	}
	locInfo := first.CustomerInfo.Address
	locAddr := LocMap[locInfo]
	fmt.Println(locInfo == "山东省济宁市邹城市中心店街道鑫岳旧货市场")
	resOr = getCoordinates(loc, locAddr)
	if len(resOr) == 0 {
		resp = append(resp, []float64{loc[0][0].Longitude, loc[0][0].Latitude}, []float64{LocMap[locInfo].Longitude, LocMap[locInfo].Latitude})
	}
	// 把 resOr 转换为二维数组
	for _, coordinate := range resOr {
		resp = append(resp, []float64{coordinate.Longitude, coordinate.Latitude})
	}

	ctx.JSON(200, gin.H{
		"data": map[string]interface{}{
			"points": resp,
		},
		"code": "0000",
		"msg":  "请求成功",
	})

}

// 定义坐标点结构体
type Point struct {
	X float64
	Y float64
}

// 计算两点间的欧几里得距离
func distance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// 生成所有可能的路径排列
func permute(path []int) [][]int {
	var result [][]int
	var permuteHelper func([]int, int)
	permuteHelper = func(path []int, start int) {
		if start == len(path) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		} else {
			for i := start; i < len(path); i++ {
				path[start], path[i] = path[i], path[start]
				permuteHelper(path, start+1)
				path[start], path[i] = path[i], path[start]
			}
		}
	}
	permuteHelper(path, 0)
	return result
}

// 计算路径的总长度
func pathLength(points []Point, path []int) float64 {
	length := 0.0
	for i := 0; i < len(path)-1; i++ {
		length += distance(points[path[i]], points[path[i+1]])
	}
	return length
}

// 找出最短路径
func findShortestPath(points []Point, start, end int) ([]int, float64) {
	var waypoints []int
	for i := range points {
		if i != start && i != end {
			waypoints = append(waypoints, i)
		}
	}
	allPaths := permute(waypoints)
	shortestPath := make([]int, 0)
	shortestLength := math.MaxFloat64

	for _, path := range allPaths {
		fullPath := []int{start}
		fullPath = append(fullPath, path...)
		fullPath = append(fullPath, end)
		length := pathLength(points, fullPath)
		if length < shortestLength {
			shortestLength = length
			shortestPath = fullPath
		}
	}
	return shortestPath, shortestLength
}

func getCoordinates(loc [][]Coordinate, locInfo Coordinate) []Coordinate {
	//找到对应的坐标 截取
	for i, coordinate := range loc {
		for j, c := range coordinate {
			if c.Latitude == locInfo.Latitude && c.Longitude == locInfo.Longitude {
				return loc[i][:j+1]
			}
		}
	}
	return nil

}
