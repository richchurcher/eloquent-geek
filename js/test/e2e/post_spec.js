describe('Eloquent Geek', function() {

  describe('Post list view', function() {

    beforeEach(function() {
      browser.get(browser.params.app);
    });

    it('should delete all posts using the delete buttons',
       function() {
         element.all(by.css('button.deletePost')).then( function(elements) {
           elements.forEach( function(e) {
             e.click();
           });
         });
         var postList = element.all(by.repeater('post in posts'));
         expect(postList.count()).toBe(0);
       }
    );

    it('should create posts using the form',
       function() {
         for (var i = 0; i < 10; i++) {
           element(by.model('post.title')).clear().sendKeys('Protractor');
           element(by.model('post.body')).clear().sendKeys('Test');
           element(by.model('post.tags')).clear().sendKeys('tag1 tag2 tag3\n');
         }
         var postList = element.all(by.repeater('post in posts'));
         expect(postList.count()).toBe(10);
       }
    );
  });
});
