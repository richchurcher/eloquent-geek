describe('PostCtrl', function() {
  beforeEach(module('post'));

  it("should create 'post' model",
    inject(function($controller) {
      var scope = {};
      var ctrl = $controller('PostCtrl', {$scope:scope});

      expect(scope.loadPosts).toBeDefined();
      expect(scope.deletePost).toBeDefined();
      expect(scope.createPost).toBeDefined();
    })
  );
});

//describe('PostCtrl', function() {
  //beforeEach(module('post'));

  //it("should populate 'posts' with posts",
     //inject(function($controller) {
       //var scope = {};
       //var ctrl = $controller('PostCtrl', {$scope:scope});

       //scope.createPost({title:"foo", body: "bar", tags: "wombat"});
       //scope.loadPosts();
       //expect(scope.posts).toBe(1);
    //})
  //);
//});
