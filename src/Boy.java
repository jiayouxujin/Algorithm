/**
 * 总共两个类 主类Boy继承Human
 * 类Human里有个方法walk() Boy对该方法进行重写(overriding)
 * 然后有三个对象obj obj1 obj2
 */
class Human{
    public void walk(){
        System.out.println("Human walks");
    }
}
public class Boy extends Human {
    public void walk(){
        System.out.println("Boy walks");
    }
    public static void main(String []args){
        Human obj=new Boy();
        Human obj2=new Human();
        Boy obj3=new Boy();

        obj.walk();
        obj2.walk();
        obj3.walk();
    }
}
