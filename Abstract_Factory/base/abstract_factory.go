package base

// MazeFactory创建迷宫的组件，它建造房间，墙壁和房门之间的门。

// 有个工厂可以创建迷宫的组件, 它可以创建门，墙壁和房间之间的门。

// 它可以从一个文件当中读取迷宫说明图并创建相应的迷宫的程序，或者可以用于一个随机建造迷宫的程序

// 抽象迷宫工厂 （AbstractFactory）
type AbstractMazeFactory interface {
	MakeMaze() *Maze
	MakeWall() *Wall
	MakeRoom(no int) *Room
	MakeDoor(r1, r2 *Room) *Door
}

// 工厂方法的一个集合，这是最常见的实现Abstract Factory 模式的方式。
// MazeFactory 不是抽象类，因此它既是Abstract Factory 也是ConcreteFactory
// 这个是 Abstract Factory 的简单应用的体现。
// 因为MazeFactory 是一个完全由工厂方法组成的具体类，它可以生成一个字类并重定义
// 需要改变的操作，这样就很容易生成一个新的MazeFactory
type MazeFactory struct{}

func NewMazeFactory() *MazeFactory {
	return &MazeFactory{}
}

// 每个方法可以理解为一个产品制造

func (m *MazeFactory) MakeMaze() *Maze {
	return NewMaze()
}
func (m *MazeFactory) MakeWall() *Wall {
	return NewWall()
}
func (m *MazeFactory) MakeRoom(no int) *Room {
	return NewRoom(no)
}
func (m *MazeFactory) MakeDoor(r1, r2 *Room) *Door {
	return NewDoor(r1, r2)
}

// 使用工厂方法创建Maze
func (mg *MazeGame) CreateMazeByFactory(factory AbstractMazeFactory) *Maze {
	aMaze := factory.MakeMaze()
	r1 := factory.MakeRoom(1)
	r2 := factory.MakeRoom(2)
	aDoor := factory.MakeDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(North, factory.MakeWall())
	r1.SetSide(East, aDoor)
	r1.SetSide(South, factory.MakeWall())
	r1.SetSide(West, factory.MakeWall())

	r1.SetSide(North, factory.MakeWall())
	r1.SetSide(East, factory.MakeWall())
	r1.SetSide(South, factory.MakeWall())
	r1.SetSide(West, aDoor)

	return aMaze
}

// 创建一个施了魔法的迷宫 (ConcreteFactory)
type EnchantedMazeFactory struct {
	// 继承工厂方法
	MazeFactory
}

// 施魔法
type Spell struct{}

func (e *EnchantedMazeFactory) MakeRoom(no int) *Room {
	// 创建被施了魔法的房间
	return EnchantedRoom(no, e.CastSpell())
}

func (e *EnchantedMazeFactory) MakeDoor(r1, r2 *Room) *Door {
	return DoorNeedingSpell(r1, r2)
}

// 施魔法
func (e *EnchantedMazeFactory) CastSpell() *Spell {
	return &Spell{}
}

// 创建被施了魔法的房间
func EnchantedRoom(no int, spell *Spell) *Room {
	return &Room{roomNumber: no}
}

//创建施魔法的门
func DoorNeedingSpell(r1, r2 *Room) *Door {
	return &Door{
		roomIn:  r1,
		roomOut: r2,
	}
}

// 创建一个有炸弹的迷宫

// 炸弹迷宫工厂 (ConcreteFactory)
type BombedMazeFactory struct {
	MazeFactory
}

func NewBombedMazeFactory() *BombedMazeFactory {
	return &BombedMazeFactory{}
}

// 增加产品生成方法
func BombedWall() *Wall {
	return &Wall{}
}

func RoomWithABoomb(no int) *Room {
	return &Room{roomNumber: no}
}

// 重写工厂部分方法
func (bmf *BombedMazeFactory) MakeWall() *Wall {
	return BombedWall()
}

func (bmf *BombedMazeFactory) MakeRoom(no int) *Room {
	return RoomWithABoomb(no)
}

// 使用工厂创建一个包含炸弹的简单迷宫

func CreateMaze() *Maze {
	game := NewMazeGame()

	factory := NewBombedMazeFactory()

	return game.CreateMazeByFactory(factory)
}

// 标准抽象工厂模式示例
//
//
//
// ----------------------
//
//
//
// -----------------------

// 抽象工厂（AbstractFactory）
type WidgetFactory interface {
	// 创建抽象ScrollBar产品
	CreateScrollBar() ScrollBar
	// 创建抽象Window产品
	CreateWindow() Window
}

// client 使用函数
func NewWidgetFactory(style string) WidgetFactory {
	switch style {
	case "motify":
		return &MotifWidgetFactory{}
	case "pm":
		return &PMWidgetFactory{}
	default:
		return nil
	}
}

var (
	_ WidgetFactory = &MotifWidgetFactory{}
	_ WidgetFactory = &PMWidgetFactory{}
)

// ConcreteFactory1
type MotifWidgetFactory struct {
	style string
}

func (mwf *MotifWidgetFactory) CreateScrollBar() ScrollBar {
	return &MotifyScrollBar{}
}
func (mwf *MotifWidgetFactory) CreateWindow() Window {
	return &MotifyWindow{}
}

// ConcreteFactory2
type PMWidgetFactory struct{}

func (mwf *PMWidgetFactory) CreateScrollBar() ScrollBar {
	return &PMScrollBar{}
}
func (mwf *PMWidgetFactory) CreateWindow() Window {
	return &PMWindow{}
}

// AbstractProduct1
type Window interface {
	Show()
}

var (
	_ Window = &MotifyWindow{}
	_ Window = &PMWindow{}
)

// client 使用的函数
func NewWidows(style string) Window {
	switch style {
	case "motify":
		return &MotifyWindow{}
	case "pm":
		return &PMWindow{}
	default:
		return nil
	}
}

//ConcreteProduct1
type MotifyWindow struct{}

func (m *MotifyWindow) Show() {}

type PMWindow struct{}

func (m *PMWindow) Show() {}

// AbstractProduct2
type ScrollBar interface {
	Scroll()
}

var (
	_ ScrollBar = &MotifyScrollBar{}
	_ ScrollBar = &PMScrollBar{}
)

// client 客户端使用的函数
func NewScrollBar(style string) ScrollBar {
	switch style {
	case "motify":
		return &MotifyScrollBar{}
	case "pm":
		return &PMScrollBar{}
	default:
		return nil
	}
}

//ConcreteProduct2
type MotifyScrollBar struct{}

func (m *MotifyScrollBar) Scroll() {}

type PMScrollBar struct{}

func (m *PMScrollBar) Scroll() {}

// Client 只使用了 AbstractFactory 和AbstractProduct 接口中的方法
func CreateWidgte() {
	// 创建工厂
	motifyFactory := NewWidgetFactory("motify")
	pmFactory := NewWidgetFactory("pm")

	// 生产Window产品
	motifyWindow := motifyFactory.CreateWindow()
	pmWindow := pmFactory.CreateWindow()

	// 生产Scroll产品
	motifyScroll := motifyFactory.CreateScrollBar()
	pmScroll := pmFactory.CreateScrollBar()

	// 产品功能
	motifyWindow.Show()
	pmWindow.Show()

	motifyScroll.Scroll()
	pmScroll.Scroll()

}
