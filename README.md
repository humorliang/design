# 设计模式
这是一个设计模式的仓库，关于go 在各种设计模式中的体现

## 创建型设计模式

**理解**
创建型设计模式抽象了实例化过程，它们独立于一个系统独立于如何创建、组合、和表示它的那些对象。一个类的创建型模式是使用继承改变被示例化的类，而一个对象创建型模式将实例化委托给另外一个对象。

随着系统的演变转变为对象复合而不是继承，则从一组固定行为的硬编码变成了，定义一个较小的基本行为集，这些行为可以组合成任意数目的更复杂行为。


****
- [Abstract Factory(抽象工厂)](./Abstract_Factory/Readme.md)
- [Builder(生成器)](./Builder/Readme.md)
- [Factory Method(工厂方法)](./Factory_Method/Readme.md)
- [Prototype(原型)](./Prototype/Readme.md)
- [Singleton(单例)](./Singleton/Readme.md)

## 结构型设计模式
- [Adapter(适配器)](./Adapter/Readme.md)
- [Bridge(桥接)](./Bridge/Readme.md)
- [Composite(组合)](./Composite/Readme.md)
- [Decorator(装饰)](./Decorator/Readme.md)
- [Facade(外观)](./Facade/Readme.md)
- [Flyweight(享元)](./Flyweight/Readme.md)
- [Proxy(代理)](./Proxy/Readme.md)

## 行为模式
- [Chain Of Responsibility(责任链)](./Chain_Of_Responsibility/Readme.md)
- [Command(命令)](./Command/Readme.md)
- [Interpreter(解释器)](./Interpreter/Readme.md)
- [Iterator(迭代器)](./Iterator/Readme.md)
- [Mediator(中介者)](./Mediator/Readme.md)
- [Memento(备忘录)](./Memento/Readme.md)
- [Observer(观察者)](./Observer/Readme.md)
- [State(状态)](./State/Readme.md)
- [Strategy(策略)](./Strategy/Readme.md)
- [Template Method(模板方法)](./Template_Method/Readme.md)
- [Visitor(访问者)](./Visitor/Readme.md)


## 图型说明
1. 类继承表示为一个从子类(图中LineShape) 到父类(图中的shape) 的三角形连线;
![a](vx_images/vx20220515143133.png)
2. 代表部分或者聚集关系的对象引用表示为一个根部有菱形的箭头，指向被聚集的类(图中的Shape);

![b](vx_images/vx20220515143151.png)
3. 根部没有菱形的箭头表示相识关系(图中LineShape有一个指向Color的引用，而Color可能是多个Shape对象共享的)
![](vx_images/vx20220515143308.png)
4. 哪个类创建哪个类对象我们用虚线箭头标记这种情况，箭头指向被实例化的对象
5. 定义了一种实心圆点，表示多于一个。当圆点位于引用的头部，它表示指向或聚集多个对象，
6. 还有一些伪代码
![](vx_images/vx20220515143356.png)

