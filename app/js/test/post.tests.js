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

  //it("should equal 5",  inject(function(myservice, $timeout) {

      //var valueToVerify;
      //myservice.DoIt().then(function(returned) {
        //valueToVerify = returned;  
      //});  
      //$timeout.flush();        
      //expect(valueToVerify).toEqual(5);
  //}));

  it("should create a post",
     inject(function($controller) {

       //var actual;
       var scope = {};
       var ctrl = $controller('PostCtrl', {$scope:scope});
       scope.loadPosts();
       

       //scope.createPost({title:"foo", body: "bar", tags: "wombat"});
       //scope.loadPosts().then(function(returned) {
         //actual = returned;
       //});

       //expect(actual).toBeDefined();
    })
  );
});

