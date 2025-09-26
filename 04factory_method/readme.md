### 🔹 概念

工厂方法模式（Factory Method Pattern）：
定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法让一个类的实例化延迟到其子类。

👉 通俗理解

* 你去 餐厅点餐：

* 你只需要告诉服务员要「披萨」还是「汉堡」，不需要知道具体怎么做。

* 厨房（工厂）里会根据菜品类型，交给不同的厨师（具体工厂）来做。

* 你点菜（调用工厂方法） → 得到食物（具体对象）。

### 🔹 角色

* Product（抽象产品）：定义产品的公共接口。

* ConcreteProduct（具体产品）：实现 Product 的具体类，例如 Pizza、Hamburger。

* Factory（抽象工厂接口）：定义创建产品的工厂方法。

* ConcreteFactory（具体工厂）：实现工厂方法，返回具体产品实例。

* Client（客户端）：通过工厂获取产品对象，不关心具体创建逻辑。

### 🔹 优缺点

✅ 优点：

* 符合开闭原则（新增产品只需增加工厂和产品类，不修改原有代码）

* 将对象的创建和使用解耦

* 方便扩展

⚠️ 缺点：

* 类文件会增多（每个产品都要对应一个工厂）

* 系统复杂度增加

### 🔹 UML 图

```javascript
      Client
         │
         ▼
     ┌─────────────┐
     │   Creator   │   <── 抽象工厂接口
     │ + FactoryMethod() │
     └─────┬───────┘
           │
 ┌─────────┴─────────┐
 │                   │
ConcreteCreatorA   ConcreteCreatorB
│ FactoryMethod()   │ FactoryMethod()
▼                   ▼
ConcreteProductA   ConcreteProductB
       ▲                 ▲
       │                 │
      Product ────────────┘  <── 抽象产品接口

```