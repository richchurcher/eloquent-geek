angular

  .module('post')

  .directive('postDisplay', [
    '$sce',
    'postApiService',
    postDisplay
  ])

  .filter('trusted', ['$sce', function($sce) {
    return function (text) {
      return $sce.trustAsHtml(text);
    };
  }]);

function postDisplay($sce, postApiService) {

  var controller = function() {
    var vm = this;

    vm.loadPost = function(id, nav) {
      // Markdown
      var converter = new showdown.Converter();

      var params = { postId: id };
      if (nav) {
        params.nav = nav;
      }
      return postApiService.get(params, function (response) {
        response.body = converter.makeHtml(response.body);
        vm.post = response
        vm.style.css = response.style;
      }, function () {
        // Error: could be a 404, no posts exist
      });
    };
    
    //vm.deletePost = function(id, i) {
      //return postApiService.delete({
        //postId: id,
      //}).$promise.then(function () {
          //vm.posts.splice(i, 1);
      //}, function (error) {
          //// TODO: handle error
      //});
    //};

    vm.createPost = function() {
      if (!vm.post.tags) vm.post.tags = '';
      return postApiService.save({
        title: vm.post.title,
        body: vm.post.body,
        style: vm.post.style,
        image: vm.post.image,
        tags: vm.post.tags.split(' '),
      }, function (response) {
        var converter = new showdown.Converter();
        response.body = converter.makeHtml(response.body);
        $sce.trustAsHtml(response.body);
        vm.style.css = response.style;
        vm.post = response;
        vm.showCreateForm = false;
      });
    };

    if (!vm.post) {
      vm.loadPost('latest');
    }

    vm.newPost = function () {
      for (var k in vm.post) {
        vm.post[k] = null;
      }
      vm.showCreateForm = true;
      vm.style.css = '';
    }
  };

  return {
    bindToController: true,
    controller: controller,
    controllerAs: 'vm',
    scope: {
      style: '='
    },
    templateUrl: '/js/post/postDisplay.html'
  };
}

