# 原型模式 (Prototype Pattern)

## 什么是原型模式？

原型模式就像"复印机"一样，可以快速复制现有对象，而不需要重新创建。

想象一下：
- 你有一份重要文件，需要给10个人看
- 与其重新写10遍，不如用复印机复印9份
- 原型模式就是代码中的"复印机"

## 为什么要用原型模式？

1. **创建对象成本高**：比如从数据库加载数据很慢
2. **对象结构复杂**：有很多嵌套的属性
3. **需要大量相似对象**：只是某些属性不同

## 生活中的例子

- 🖨️ 复印文件：原件不变，复印出新的
- 🍪 用模具做饼干：模具是原型，做出很多饼干
- 📋 填表格：有模板，每次复制后填不同内容

## 代码实现

### 基础版本
```go
// 可复制的接口
type Cloneable interface {
    Clone() Cloneable
}

// 简历对象
type Resume struct {
    Name string
    Age  int
}

// 实现复制方法
func (r *Resume) Clone() *Resume {
    return &Resume{
        Name: r.Name,
        Age:  r.Age,
    }
}
```

### 实际应用场景
```go
// 游戏角色原型
type Character struct {
    Name   string
    Level  int
    Skills []string
}

func (c *Character) Clone() *Character {
    // 深拷贝技能列表
    skills := make([]string, len(c.Skills))
    copy(skills, c.Skills)
    
    return &Character{
        Name:   c.Name,
        Level:  c.Level,
        Skills: skills,
    }
}
```

## 使用示例

```go
func main() {
    // 创建原型角色
    warrior := &Character{
        Name:   "战士模板",
        Level:  1,
        Skills: []string{"攻击", "防御"},
    }
    
    // 快速创建新角色
    player1 := warrior.Clone()
    player1.Name = "张三"
    
    player2 := warrior.Clone()
    player2.Name = "李四"
    player2.Level = 5
    
    fmt.Printf("玩家1: %+v\n", player1)
    fmt.Printf("玩家2: %+v\n", player2)
}
```

## 注意事项

### 浅拷贝 vs 深拷贝

```go
// ❌ 浅拷贝 - 危险！
func (c *Character) ShallowClone() *Character {
    return &Character{
        Name:   c.Name,
        Skills: c.Skills, // 共享同一个切片！
    }
}

// ✅ 深拷贝 - 安全
func (c *Character) DeepClone() *Character {
    skills := make([]string, len(c.Skills))
    copy(skills, c.Skills) // 创建新的切片
    
    return &Character{
        Name:   c.Name,
        Skills: skills,
    }
}
```

## 优缺点

### 优点
- ✅ 创建对象速度快
- ✅ 避免重复的初始化代码
- ✅ 可以动态配置对象

### 缺点
- ❌ 需要实现克隆方法
- ❌ 深拷贝可能比较复杂
- ❌ 循环引用时容易出错

## 适用场景

1. **对象创建成本高**：数据库查询、网络请求
2. **需要大量相似对象**：游戏中的NPC、UI组件
3. **对象状态变化频繁**：需要保存多个版本

## 与其他模式的区别

| 模式 | 用途 | 特点 |
|------|------|------|
| 原型模式 | 复制现有对象 | 基于已有实例 |
| 工厂模式 | 创建新对象 | 基于类型参数 |
| 建造者模式 | 构建复杂对象 | 分步骤构建 |

原型模式就是代码世界的"复印机"，让对象创建变得简单快捷！