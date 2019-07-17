define([
	'underscore',
	'backbone'
], function (_, Backbone) {

	var TaskModel = Backbone.Model.extend({
        urlRoot: '/api/tasks',
		defaults: {
            title: '',
			description: ''
		}
	});

	return TaskModel;
});