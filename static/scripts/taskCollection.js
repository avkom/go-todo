define([
	'underscore',
	'backbone',
	'scripts/taskModel'
], function (_, Backbone, TaskModel) {

	var TaskCollection = Backbone.Collection.extend({
		model: TaskModel,
        comparator: 'id',
        url: '/api/tasks'
	});

	return TaskCollection;
});