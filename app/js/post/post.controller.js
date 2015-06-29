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

  var PostCtrl = function() {
    var vm = this;

    vm.loadPosts = function() {
      var converter = new showdown.Converter();
      return postApiService.query(function (data) {
        // Markdown
        for (var i = 0; i < data.length; i++) {
          data[i].body = converter.makeHtml(data[i].body);  
          vm.style.css = data[i].style;
        }
        vm.posts = data;
      });
    };
    
    vm.deletePost = function(id, i) {
      return postApiService.delete({
        postId: id,
      }).$promise.then(function () {
          vm.posts.splice(i, 1);
      }, function (error) {
          // TODO: handle error
      });
    };

    vm.createPost = function(post) {
      if (!post.tags) post.tags = '';
      return postApiService.save({
        title: post.title,
        body: post.body,
        tags: post.tags.split(' '),
      }, function (response) {
        var converter = new showdown.Converter();
        response.body = converter.makeHtml(
          $sce.trustAsHtml(response.body)
        );
        vm.style.css = response.style;
        vm.posts.push(response);
      });
    };

    vm.loadPosts();
  };

  return {
    bindToController: true,
    controller: PostCtrl,
    controllerAs: 'vm',
    scope: {
      style: '=',
    },
    templateUrl: '/js/post/postDisplay.html',
  };
}

