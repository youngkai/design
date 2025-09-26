### 🔹 概念

建造者模式（Builder Pattern）：
将一个复杂对象的构建过程与它的表示分离，使得同样的构建过程可以创建不同的表示。

### 👉 通俗理解

就像 点套餐：

* 服务员（Director）按流程来：先主食 → 配菜 → 饮料。

* 厨房里有不同的厨师（ConcreteBuilder），可以做「中餐套餐」或「西餐套餐」。

* 客户只管点套餐，不需要关心具体做法。

### 🔹 角色

* Product（产品）：最终要建造的复杂对象。

* Builder（抽象建造者）：定义建造产品的步骤。

* ConcreteBuilder（具体建造者）：实现建造步骤，组装出具体产品。

* Director（指挥者）：决定调用哪些建造步骤，以及顺序。

* Client（客户）：通过指挥者获得完整产品。

### 🔹 优缺点

✅ 优点

* 将构建和表示分离，客户端只管“要什么”，不用管“怎么造”。

* 同一构建过程可以创建不同的产品表示。

* 更符合单一职责原则，易于维护和扩展。

⚠️ 缺点

* 会增加类的数量，结构稍复杂。

* 适合“复杂对象”的构建场景，太简单的对象不需要用建造者。

### 🔹 UML 图

```javascript
           Director
               │
               ▼
       ┌──────────────┐
       │   Builder    │  <── 抽象建造者
       │ +BuildPartA()│
       │ +BuildPartB()│
       │ +GetResult() │
       └───────┬──────┘
               │
   ┌───────────┴───────────┐
   ▼                       ▼
ConcreteBuilder1       ConcreteBuilder2
│ BuildPartA()          │ BuildPartA()
│ BuildPartB()          │ BuildPartB()
│ GetResult()           │ GetResult()
       ▼                       ▼
       Product1                Product2

```