import java.util.Scanner;
class linkedList {
    Node head;
    class Node{
        int id;
        Node next;
        Node(int _id){
            id = _id;
            next = null;
        }
    }
    void push_bak(int _id){
        System.out.println("pushing new data : "+_id);
        if (head == null) {
            Node newNode = new Node(_id);
            head = newNode;
        }else{
            Node ptr = head;
            while(ptr.next != null){
                ptr = ptr.next;
            }
            ptr.next = new Node(_id);
        }
    }
    void print_data(){
        Node ptr = head;
        while(ptr != null){
             System.out.println("printing data : "+ptr.id);
             ptr = ptr.next;
        }
    }
    public static void main(String arg[]){
        System.out.println("main entered");
        linkedList obj = new linkedList();
       Scanner myObj = new Scanner(System.in);
        System.out.println("enter inputs size ");
        int len = myObj.nextInt();
        for(int i = 0 ; i < len ;  i++){
            obj.push_bak(myObj.nextInt());
        } 
        obj.print_data();
    }

}