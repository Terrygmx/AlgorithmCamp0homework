class MedianFinder {
    Queue<Integer> A, B;
    public MedianFinder() {
        A = new PriorityQueue<>(); // 小顶堆，保存较大的一半
        B = new PriorityQueue<>((x, y) -> (y - x)); // 大顶堆，保存较小的一半
    }
    public void addNum(int num) {
        if(A.size() != B.size()) { // 当 m ！= n（即 和 为 奇数）：需向 B添加一个元素。实现方法：将新元素 num 插入至 A ，再将 A 堆顶元素插入至 B ；
            A.add(num);
            B.add(A.poll());
        } else { // 当 m = n（即 和 为 偶数）：需向 A添加一个元素。实现方法：将新元素 num 插入至 B ，再将 B 堆顶元素插入至 A ；
            B.add(num);
            A.add(B.poll());
        }
    }
    public double findMedian() {// 当 m = n（和 为 偶数）：则中位数为 ( A 的堆顶元素 + B 的堆顶元素 )/2。 当 m ！= n（ 和 为 奇数）：则中位数为 A 的堆顶元素。
        return A.size() != B.size() ? A.peek() : (A.peek() + B.peek()) / 2.0;
    }
}

