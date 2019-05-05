/**
 * 有三个类 主类C的直接父类为B B的父类是A
 * 其中main函数定义了两个对象
 */
class A{
    public void a1(){
        System.out.println("A");
    }

    protected void b1() {
    }
}
class B extends A{
    public void a1(){
        System.out.println("B");
    }
    public void b1(){
        System.out.println("B1");
    }
}
public class C extends B{
    public static void main(String []args){
        A a=new B();
        a.a1();
        a.b1();
        ((B) a).b1();
        B b=new B();
        b.a1();
    }
}
