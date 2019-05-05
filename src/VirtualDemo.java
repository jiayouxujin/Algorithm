
public class VirtualDemo {
    public static void main(String []args){
        Salary s=new Salary("jin","shanghai",1,3.0);
        Employee e=new Salary("xu","sh",2,3.0);
        Employee ee=new Employee("xu","sh",3);

        VirtualDemo virtualDemo=new VirtualDemo();
        virtualDemo.testOverloading(e);
    }
    private void testOverloading(Employee e){
        System.out.println("Employee");
    }
    private void testOverloading(Salary ss){
        System.out.println("Salary");
    }
}
