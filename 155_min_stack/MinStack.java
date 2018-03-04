class MinStack {
    Node head;

    /** initialize your data structure here. */
    public MinStack() {
        head = new Node(0);
    }

    public void push(int x) {
        if (head.next == null) {
            head.next = new Node(x);
            head.next.min = x;
        } else {
            Node node = new Node(x);
            node.next = head.next;
            node.min = Math.min(node.next.min, x);
            head.next = node;
        }
    }

    public void pop() {
        if (head.next != null) {
            head.next = head.next.next;
        }
    }

    public int top() {
        return head.next == null ? 0 : head.next.val;
    }

    public int getMin() {
        return head.next == null ? 0 : head.next.min;
    }

    private class Node {
        int val;
        int min;
        Node next;
        public Node(int val) {
            this.val = val;
        }
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */