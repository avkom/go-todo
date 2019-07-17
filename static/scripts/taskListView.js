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
            this._collection.fetch().then(this.render.bind(this));
		},

		render: function () {
            var data = {
                tasks: this._collection.toJSON()
            };
			this.$el.html(this.template(data));
			return this;
		},
	});

	return TaskListView;
});