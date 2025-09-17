package prototype

import (
	"fmt"
	"testing"
)

func TestCharacterClone(t *testing.T) {
	// 创建原型角色
	warrior := &Character{
		Name:   "战士模板",
		Level:  1,
		Skills: []string{"攻击", "防御"},
	}

	// 克隆新角色
	player1 := warrior.Clone()
	player1.Name = "张三"
	player1.Level = 5

	player2 := warrior.Clone()
	player2.Name = "李四"
	player2.Skills = append(player2.Skills, "治疗")

	// 验证原型没有被修改
	if warrior.Name != "战士模板" {
		t.Error("原型被意外修改")
	}

	// 验证克隆对象独立
	if len(warrior.Skills) != 2 {
		t.Error("原型技能被修改")
	}

	fmt.Printf("原型: %+v\n", warrior)
	fmt.Printf("玩家1: %+v\n", player1)
	fmt.Printf("玩家2: %+v\n", player2)
}

func TestResumeClone(t *testing.T) {
	// 创建简历模板
	template := &Resume{
		Name: "模板",
		Age:  0,
	}

	// 快速创建多份简历
	resume1 := template.Clone()
	resume1.Name = "张三"
	resume1.Age = 25

	resume2 := template.Clone()
	resume2.Name = "李四"
	resume2.Age = 30

	fmt.Printf("模板: %+v\n", template)
	fmt.Printf("简历1: %+v\n", resume1)
	fmt.Printf("简历2: %+v\n", resume2)
}

// 性能测试：对比直接创建 vs 克隆
func BenchmarkDirectCreate(b *testing.B) {
	var tmp *Character
	for i := 0; i < b.N; i++ {
		tmp = &Character{
			Name:   "新角色",
			Level:  1,
			Skills: []string{"攻击", "防御", "跳跃"},
		}
		tmp.Name = "新角色"
	}
}

func BenchmarkClone(b *testing.B) {
	prototype := &Character{
		Name:   "原型",
		Level:  1,
		Skills: []string{"攻击", "防御", "跳跃"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clone := prototype.Clone()
		clone.Name = "新角色"
	}
}
