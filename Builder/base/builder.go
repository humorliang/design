package base

// Product()
// 迷宫
type Maze struct{}

// 方位
type Direction string

// 迷宫游戏
// 客户端
type MazeGame struct{}

// 房间
type Room struct{}

// Builder 角色
// 创建一个迷宫生成器
type MazeBuilder interface {
	BuildMaze()
	BuildRoom(no int)
	BuildDoor(roomFrom int, roomTo int)

	// 抽象出来的获取产品的接口
	GetMaze() *Maze
}

// Builder模式封装了对象时如何被创建的,我们可以重用MazeBuilder来创建
// 不同种类的迷宫

// Director 角色
// 客户端使用生成器创建迷宫
// 抽象出来的获取产品的接口
/*
这个CreateMaze 的版本与原来的相比，隐藏了迷宫的内部表示（即房间、门和墙壁的那些类）
这样使得迷宫的表示方式要更容易一些。这样所有的MazeBuilder 的客户端都不需要被改变。
*/
func (g *MazeGame) CreateMaze(builder MazeBuilder) *Maze {
	builder.BuildMaze()

	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 2)

	return builder.GetMaze()
}

// 还可以创建这样的迷宫
func (g *MazeGame) CreateCompleMaze(builder MazeBuilder) *Maze {
	builder.BuildRoom(1)
	// ...
	builder.BuildRoom(1001)
	return builder.GetMaze()
}

// MazeBuilder 自己不创建迷宫。它的存在目的就是为创建迷宫定义一个接口
// 它主要方便起见定义一些空的实现。实际工作是MazeBuilder它的子类也就是实现这个接口的struct。

var _ MazeBuilder = &StanderMazeBuilder{}

// ConcreteBuilder 角色
// 创建一个简单迷宫的实现
type StanderMazeBuilder struct {
	MazeBuilder

	currentMaze *Maze
}

// 构造函数
func NewStanderMazeBuilder() *StanderMazeBuilder {
	return &StanderMazeBuilder{}
}

func (s *StanderMazeBuilder) BuildMaze() {
	//todo 具体细节操作
}
func (s *StanderMazeBuilder) BuildRoom(no int) {
	//todo 具体细节操作
}
func (s *StanderMazeBuilder) BuildDoor(roomFrom int, roomTo int) {
	//todo 具体细节操作
}

func (s *StanderMazeBuilder) GetMaze() *Maze {
	return &Maze{}
}

// 描述两个房间的公共墙壁
func (s *StanderMazeBuilder) CommonWall() {
	//todo 具体细节操作
}

// 客户端可以使用CreateMaze和StanderMazeBuidler来创建一个迷宫
func Client() {
	var game = &MazeGame{}
	// 指定一个生成器
	var builder = &StanderMazeBuilder{}

	game.CreateMaze(builder)

	builder.GetMaze()
}

// 我们本可以把StanderMazeBuilder 操作放在Maze 中并让每一个Maze创建它自身，但是我们将
// maze 变得小一点使得它更容易被理解和修改， 而且StanderMazeBuilder 易于从Maze中分离，
// 这种分离之后我们可以有多种MazeBuilder, 每一种使用不同的房间、墙壁和门的类

// 还有一种更特殊的Builder 例子,这个生成器不创建迷宫，仅仅是对被已经创建的不同种类的组件进行计数
type CountingMazeBuilder struct {
	MazeBuilder

	doors int
	rooms int
}

// 重写部分方法只是增加相应的计数
func (c *CountingMazeBuilder) BuildRoom(no int) {
	c.rooms++
}

func (c *CountingMazeBuilder) BuildDoor(roomFrom int, roomTo int) {
	c.doors++
}

func (c *CountingMazeBuilder) GetCounts() (d int, r int) {
	return c.doors, c.rooms
}

func Client2() {
	var (
		game    = &MazeGame{}
		builder = &CountingMazeBuilder{}
	)

	game.CreateMaze(builder)

	builder.GetCounts()
}
