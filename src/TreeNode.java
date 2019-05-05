public class TreeNode {
    private int val;
    private TreeNode left;
    private TreeNode right;

    public TreeNode(){ }
    public TreeNode(int val) {
        this.val = val;
    }
    public static TreeNode solu(int[] preorder,int[] inorder,int prestart,int preend,int instart,int inend){
        if(prestart>preend) return null;
        if(prestart==preend) return new TreeNode(preorder[prestart]);

        TreeNode head=new TreeNode(preorder[prestart]);
        int low=instart;
        for(int i=low;i<=inend;i++){
            if(inorder[i]==preorder[prestart]){
                head.left=solu(preorder,inorder,prestart+1,prestart+i-instart,instart,i-1);
                head.right=solu(preorder,inorder,prestart+i-instart+1,preend,i+1,inend);
                break;
            }
        }
        return head;
    }
    public static void post(TreeNode head){
        if(head!=null){
            post(head.left);
            post(head.right);
            System.out.println(head.val);
        }
    }
    public static void main(String[] args){
        int[] preorder={1,2,4,8,9,5,10,3,6,7};
        int[] inorder={8,4,9,2,10,5,1,6,3,7};

        TreeNode head=new TreeNode();
        head=solu(preorder,inorder,0,preorder.length-1,0,inorder.length-1);
        post(head);
    }
}
