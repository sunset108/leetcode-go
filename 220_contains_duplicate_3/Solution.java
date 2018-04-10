class Solution {
    public boolean containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        if (k <= 0 || t < 0) {
            return false;
        }
        TreeSet<Long> tree = new TreeSet<>();
        for (int i = 0; i < nums.length; i++) {
            Long x = tree.ceiling((long)nums[i] - t);
            if (x != null && x <= (long)nums[i] + t) {
                return true;
            }
            if (i >= k) {
                tree.remove((long)nums[i - k]);
            }
            tree.add((long)nums[i]);
        }
        return false;
    }
}
