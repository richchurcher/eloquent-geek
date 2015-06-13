describe('postCtrl function', function() {
  
  describe('postCtrl', function() {
    var $scope;

    beforeEach(module('post'));

    beforeEach(inject(function($rootScope, $controller) {
      $scope = $rootScope.$new();
      $controller('PostCtrl', {$scope: $scope});
    }));

    it('should have a loadPosts method', function () {
      expect($scope.loadPosts).toBeDefined();
    });
  });

});

  //beforeEach(module('post'));

  //it('should create model',
    //inject(function($controller) {
      //var scope = {};
      //expect($controller('PostCtrl', {$scope:scope})).toBeDefined();
    //})
  //);

  //it('should create and delete a post',
    //inject(function($controller) {
      //var scope = {};
      //var ctrl = $controller('PostCtrl', {$scope:scope});

      ////scope.createPost({
        ////title:'foo', 
        ////body: 'bar', 
        ////tags: 'wombat', 
        ////style:''
      ////}).$promise.then(function(returned) {
        ////expect(returned.title).toBe('blah');

        ////scope.deletePost({
          ////postId: returned.ID
        ////}, 0).$promise.then(function(result) {
          ////console.log(result);
        ////});

      ////});

      //scope.createPost({
        //title:'foo', 
        //body: 'bar', 
        //tags: 'wombat', 
        //style:''
      //}).(function(returned) {
      //});
    //})
  //);

  //it('should delete a post',
      //inject(function($controller) {
        //var scope = {};
        //var ctrl = $controller('PostCtrl', {$scope:scope});

      //scope.createPost({
        //title:'foo', 
        //body: 'bar', 
        //tags: 'wombat', 
        //style: ''
      //}).$promise.then(function(returned) {
        //expect(returned.title).toBe('foo');

        //scope.deletePost({
          //postId: returned.ID
        //}).$promise.then(function(deleteResult) {
          //console.log(deleteResult);
        //});
      //});
    //});
  //);        
//});

