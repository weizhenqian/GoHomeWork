/* （这是一种综合练习，使用到第 4、9、11 章和本章的内容。）写一个可交互的控制台程序，要求用户输入二位平面极坐标上的点（半径和角度（度））。计算对应的笛卡尔坐标系的点的 x 和 y 并输出。使用极坐标和笛卡尔坐标的结构体。

使用通道和协程：

channel1 用来接收极坐标
channel2 用来接收笛卡尔坐标
转换过程需要在协程中进行，从 channel1 中读取然后发送到 channel2。实际上做这种计算不提倡使用协程和通道，但是如果运算量很大很耗时，这种方案设计就非常合适了。 */
package main

import "fmt"

var (
	radius float64
	angle  float64
)

//读取用户输入的极坐标(r,θ)[半径,角度]
func readInput() {
	fmt.Println("Please input radius:")
	fmt.Scanln(&radius)
	fmt.Println("Please input angle:")
	fmt.Scanln(&angle)
}

func getPolarCoordinates() {}

//
func getCartesianCoordinates() {}

//极坐标转换为笛卡尔坐标公式：x = r × cos( θ )、y = r × sin( θ )
func pToc() {}
