describe('PostCtrl', function() {
  beforeEach(module('post'));

  it('should create 'post' model',
    inject(function($controller) {
      var scope = {};
      expect($controller('PostCtrl', {$scope:scope})).toBeDefined();
    })
  );

  it('should create a post',
     inject(function($controller) {

       var scope = {};
       var ctrl = $controller('PostCtrl', {$scope:scope});
       scope.loadPosts();
       
       scope.createPost({title:'foo', body: 'bar', tags: 'wombat'});
       scope.loadPosts().then(function(returned) {
         actual = returned;
       });

       expect(actual).toBeDefined();
    })
  );
});

