public class Employee {
    private String name;
    private String address;
    private int number;
    public final int flag=1;

    public Employee(String name,String address,int number){
//        System.out.println("Constructing an Employee");
        this.name=name;
        this.address=address;
        this.number=number;
    }
    public static void mailCheck(){
        System.out.println("the static method of employee");
    }
    public void mailCheck2(){
        System.out.println("the public method of employee");
    }
    private void mailCheck3(){
        System.out.println("the private employee method");
    }
}

