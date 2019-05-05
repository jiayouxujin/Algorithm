public class Salary extends Employee {
    private double salary;
    public final int flag=2;

    public Salary(String name,String address,int number,double salary){
        super(name,address,number);
        setSalary(salary);
    }
    public void setSalary(double salary){
        if(salary>=0.0){
            this.salary=salary;
        }
    }
    public static void mailCheck(){
        System.out.println("with salary static method");
    }
    public void mailCheck2(){
        System.out.println("public method of salary");
    }
    private void mailCheck3(){
        System.out.println("private method of salary");
    }

}
