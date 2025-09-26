### 🔹 概念

备忘录模式的作用是 在不破坏对象封装性的前提下，捕获对象的内部状态，并在对象之外保存这个状态，以后可以在需要时将对象恢复到之前的状态。

典型应用：

1. 文本编辑器的 撤销/回滚

2. 游戏的 存档/读档

3. 事务的 回滚

### 🔹 角色

* Originator（发起人）：需要保存状态的对象

* Memento（备忘录）：存储 Originator 的内部状态（不可被外部随意修改）

* Caretaker（管理者）：负责保存和恢复备忘录

### 🔹 UML 图

```javascript
 ┌───────────────────┐
 │     Originator    │
 │-------------------│
 │- state: String    │
 │+ setState(s)      │
 │+ getState():String│
 │+ save(): Memento  │
 │+ restore(m:Memo)  │
 └─────────┬─────────┘
           │ uses
           ▼
 ┌───────────────────┐
 │     Memento       │
 │-------------------│
 │- state: String    │
 │+ getState():String│
 └───────────────────┘

 ┌───────────────────┐
 │     Caretaker     │
 │-------------------│
 │- mementoList:List │
 │+ add(m: Memento)  │
 │+ get(i:int):Memo  │
 └───────────────────┘

```