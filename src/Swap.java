import java.util.List;

public class Swap {
    public static ListNode swapPairs(ListNode head){
        if(head==null||head.next==null) return head;
        ListNode p=head;
        while(p!=null&&p.next!=null){
            ListNode temp=p.next.next;
            p.next.next=p;printListNode(head);
            p.next=temp;printListNode(head);
            p=temp;
            printListNode(head);
        }
        return head;
    }
    public static void printListNode(ListNode head){
        ListNode temp=head;
        while (temp!=null){
            System.out.print(temp.val);
            System.out.print(" ");
            temp=temp.next;
        }
        System.out.println();
    }
    public static void main(String[] args){
        ListNode head=new ListNode(1);
        ListNode curr=head;
        for (int i=2;i<=4;i++){
            curr.next=new ListNode(i);
            curr=curr.next;
        }
        printListNode(head);
        printListNode(swapPairs(head));
    }
}
