package prototype

// Cloneable 可复制接口
type Cloneable interface {
	Clone() Cloneable
}

// Character 游戏角色
type Character struct {
	Name   string
	Level  int
	Skills []string
}

// Clone 深拷贝角色
func (c *Character) Clone() *Character {
	// 复制技能切片
	skills := make([]string, len(c.Skills))
	copy(skills, c.Skills)

	return &Character{
		Name:   c.Name,
		Level:  c.Level,
		Skills: skills,
	}
}

// Resume 简历
type Resume struct {
	Name string
	Age  int
}

// Clone 复制简历
func (r *Resume) Clone() *Resume {
	return &Resume{
		Name: r.Name,
		Age:  r.Age,
	}
}
