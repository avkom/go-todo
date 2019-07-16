define([
	'jquery',
	'underscore',
	'backbone',
	'text!scripts/taskListTemplate.html'
], function ($, _, Backbone, taskListTemplate) {

	var TaskListView = Backbone.View.extend({

		el:  '.todoapp',

		template: _.template(taskListTemplate),

		events: {
		},

		initialize: function () {
		},

		render: function () {
			this.$el.html(this.template());
			return this;
		},
	});

	return TaskListView;
});