define([
	'jquery',
	'underscore',
    'backbone',
    'scripts/taskCollection',
	'text!scripts/taskListTemplate.html'
], function ($, _, Backbone, TaskCollection, taskListTemplate) {

	var TaskListView = Backbone.View.extend({

        _collection: new TaskCollection(),
        
        el: '.todoapp',

		template: _.template(taskListTemplate),

		events: {
		},

		initialize: function () {
            this._collection.fetch();
		},

		render: function () {
			this.$el.html(this.template());
			return this;
		},
	});

	return TaskListView;
});