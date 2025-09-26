### 🔹 概念

迭代器模式是一种 行为型设计模式，它提供了一种方法，可以 顺序访问一个聚合对象中的元素，而不需要暴露其内部表示。

换句话说：

> 把遍历集合的逻辑独立出来，集合本身只负责存数据，迭代器负责遍历。

### 🔹 角色

#### 1. Iterator（迭代器接口）

* 定义访问和遍历元素的方法（HasNext()、Next()）

#### 2. ConcreteIterator（具体迭代器）

* 实现迭代逻辑，维护遍历位置

#### 3. Aggregate（聚合接口）

* 定义创建迭代器的方法

#### 4. ConcreteAggregate（具体聚合类）

* 实现聚合接口，返回具体迭代器

### 🔹 UML 图

```javascript
Aggregate ───► Iterator
   ▲               ▲
   │               │
ConcreteAggregate  ConcreteIterator
```