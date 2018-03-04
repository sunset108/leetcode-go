public class LRUCache {
    Node head;
    Node tail;
    int capacity;
    HashMap<Integer, Node> map;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        map = new HashMap<>();
    }

    public int get(int key) {
        if (map.containsKey(key)) {
            Node node = map.get(key);
            node.moveToTail();
            return node.value;
        }
        return -1;
    }

    public void put(int key, int value) {
        if (map.isEmpty()) {
            head = new Node(key, value);
            tail = head;
            map.put(key, head);
            return;
        }
        if (map.containsKey(key)) {
            Node node = map.get(key);
            node.value = value;
            node.moveToTail();
        } else {
            Node node = new Node(key, value);
            node.addToTail();
            if (map.size() >= this.capacity) {
                map.remove(head.key);
                head = head.next;
                head.pre = null;
            }
            map.put(key, node);
        }
    }

    public class Node {
        int key;
        int value;
        Node pre;
        Node next;
        public Node(int key, int value) {
            this.key = key;
            this.value = value;
            this.pre = null;
            this.next = null;
        }

        public void moveToTail() {
            if (this != tail) {
                if (this != head) {
                    this.pre.next = this.next;
                    this.next.pre = this.pre;
                } else {
                    head = this.next;
                    head.pre = null;
                }
                tail.next = this;
                this.pre = tail;
                this.next = null;
                tail = this;
            }
        }

        public void addToTail() {
            tail.next = this;
            this.pre = tail;
            tail = this;
        }
    }
}