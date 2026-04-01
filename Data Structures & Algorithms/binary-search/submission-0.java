class Solution {
    public int search(int[] nums, int target) {
        int n = nums.length;

        int start = 0;
        int end = n - 1;
        while(start <= end) {
            int mid = start + (end - start) / 2;
            int focus = nums[mid];
            if(focus == target) return mid;
            else if(focus < target) start = mid + 1;
            else end = mid - 1;
        }
        return -1;
    }
}
