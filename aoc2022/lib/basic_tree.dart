class TreeNode<T> {
  T value;

  TreeNode(this.value);

  List<TreeNode<T>> children = [];

  void add(TreeNode<T> child) {
    children.add(child);
  }
}
