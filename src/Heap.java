public class Heap {
    private int[] a;
    private int n;
    private int count;

    public Heap(int capacity){
        a=new int[capacity+1];
        n=capacity;
        count=0;
    }
    public void insert(int data){
        if(count>=n) return ;
        ++count;
        a[count]=data;
        int i=count;
        while(i/2>0&&a[i]>a[i/2]){
            swap(a,i,i/2);
            i=i/2;
        }
    }
    public void removeMax(){
        if(count==0) return ;
        a[1]=a[count];
        --count;
        heapify(a,count,1);
    }

    private static void heapify(int[] a, int count, int i) {
       while((i<<1)<=count){
           int next=i<<1;
           if(next+1<=count&&a[next+1]>a[next])
               next++;
           if(a[i]<a[next]){
               swap(a,i,next);
               i=next;
           }else{
               return;
           }
       }
    }

    private static void buildHeap(int [] a,int n){
        for(int i=n/2;i>=1;--i){
            heapify(a,n,i);
        }
    }

    public static void sort(int[] a,int n){
        buildHeap(a,n);
        int k=n;
        while(k>1){
            System.out.println(a[1]);
            swap(a,1,k);
            --k;
            heapify(a,k,1);
        }
    }

    private static void swap(int[] a, int i, int i1) {
        int t=a[i];
        a[i]=a[i1];
        a[i1]=t;
    }

    public static void main(String[] args){
        System.out.println("建立堆");
       Heap heap=new Heap(10);
       for(int i=1;i<=10;i++){
           heap.a[i]=i;
       }
       buildHeap(heap.a,10);
        for(int i=1;i<=10;i++)
            System.out.println(heap.a[i]);
    }
}

