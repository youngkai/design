### 🔹 概念

抽象工厂模式（Abstract Factory Pattern）：
提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们的具体类。

👉 与工厂方法的区别：工厂方法只负责生产一种产品，而抽象工厂可以生产「一族」产品。

### 👉 通俗理解

就像你买 整套家具：

* 你选择「现代风格工厂」，它会给你「现代风格的椅子 + 现代风格的沙发」。

* 你选择「古典风格工厂」，它会给你「古典风格的椅子 + 古典风格的沙发」。

* 客户端不需要知道具体椅子、沙发怎么造，只要选择工厂就行。

### 🔹 角色

* AbstractFactory（抽象工厂接口）：定义创建一族产品的方法。

* ConcreteFactory（具体工厂）：实现抽象工厂，生产具体产品族。

* AbstractProduct（抽象产品接口）：定义产品规范。

* ConcreteProduct（具体产品）：不同风格/系列的具体实现。

* Client（客户端）：只依赖抽象工厂和抽象产品，选择工厂来获得具体产品。

### 🔹 优缺点

✅ 优点

* 保证产品族内部的一致性（风格统一）。

* 将产品的创建和使用解耦。

* 更容易扩展新的产品族。

⚠️ 缺点

* 新增产品等级结构时（比如再加「桌子」），要修改所有工厂，扩展成本高。

* 系统复杂度提高。

### 🔹 UML 图

```javascript
          Client
             │
             ▼
      ┌──────────────────┐
      │ AbstractFactory  │  <── 抽象工厂接口
      │ CreateProductA() │
      │ CreateProductB() │
      └───────┬──────────┘
              │
 ┌────────────┴─────────────┐
 │                          │
ConcreteFactory1        ConcreteFactory2
│ CreateProductA()       │ CreateProductA()
│ CreateProductB()       │ CreateProductB()
▼                          ▼
ConcreteProductA1        ConcreteProductA2
ConcreteProductB1        ConcreteProductB2
       ▲                        ▲
       │                        │
   ProductA                   ProductB  <── 抽象产品接口
```