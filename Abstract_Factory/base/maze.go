package base



type MapSite interface {
	Enter()
}

//  实现了enter 方法 ，Mapsite 定义了迷宫构件之间的主要关系
var (
	_ MapSite = &Room{}
	_ MapSite = &Wall{}
	_ MapSite = &Door{}
)

// 方向
type Direction string

const (
	North Direction = "北"
	East  Direction = "东"
	South Direction = "南"
	West  Direction = "西"
)

// 房间
type Room struct {
	//房间号
	roomNumber int
	// 其他MapSite对象的引用
	side [4]MapSite
}

// Newxxx相当于c++ 的构造函数

// 普通的房间
func NewRoom(no int) *Room {
	return &Room{roomNumber: no}
}

func (r *Room) Enter() {
	return
}

func (r *Room) SetSide(d Direction, ms MapSite) {
	return
}

func (r *Room) GetSide(d Direction) MapSite {
	return nil
}

//墙
type Wall struct{}

func NewWall() *Wall {
	return &Wall{}
}

func (w *Wall) Enter() {
	return
}

// 门
type Door struct {
	roomIn  *Room
	roomOut *Room
	isOpen  bool
}

//
func NewDoor(in *Room, out *Room) *Door {
	return &Door{
		roomIn:  in,
		roomOut: out,
	}
}

func (d *Door) Enter() {
	return
}

func (d *Door) OtherSideFrom(r *Room) *Room {
	return nil
}

// 定义房间集合类
type Maze struct {
	rooms []*Room
}

func NewMaze() *Maze {
	return &Maze{}
}

// 添加房间
func (m *Maze) AddRoom(r *Room) {
	m.rooms = append(m.rooms, r)
}

// 根据房间号查找房间
func (m *Maze) RoomNo(no int) *Room {
	for i := 0; i < len(m.rooms); i++ {
		if m.rooms[i].roomNumber == no {
			return m.rooms[i]
		}
	}
	return nil
}

// 迷宫游戏
type MazeGame struct{}

func NewMazeGame() *MazeGame {
	return &MazeGame{}
}

// 创建迷宫 使用硬编码创建 maze 对象
func (mg *MazeGame) CreateMaze() *Maze {
	aMaze := NewMaze()

	r1 := NewRoom(1)
	r2 := NewRoom(2)

	theDoor := NewDoor(r1, r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	r1.SetSide(North, NewWall())
	r1.SetSide(East, theDoor)
	r1.SetSide(South, NewWall())
	r1.SetSide(West, NewWall())

	r1.SetSide(North, NewWall())
	r1.SetSide(East, NewWall())
	r1.SetSide(South, NewWall())
	r1.SetSide(West, theDoor)

	return aMaze

}
