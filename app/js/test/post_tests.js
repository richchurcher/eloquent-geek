describe('PostCtrl', function() {
  
  var mockPostResource, $httpBackend;
  beforeEach(angular.mock.module('post'));

  beforeEach(function() {
    angular.mock.inject(function ($injector) {
      $httpBackend = $injector.get('$httpBackend');
      mockPostResource = $injector.get('Post');
    })
  });

  describe('createPost', function () {
    
    it('should create a post', 
      inject(function (Post) {
        $httpBackend.expectPOST('/post/')
          .respond({
            id: 1,
            title: 'test title',
            body: 'test body',
            tags: ['tag1', 'tag2', 'tag3'],
            style: 'body{background-color:#cccccc'
          });

        var result = mockUserResource


      

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

