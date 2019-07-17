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
            'click .new': '_onNewClicked',
            'click .delete': '_onDeleteClicked'
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
        
        _onNewClicked: function () {
            Backbone.history.navigate('tasks/new', { trigger: true });
        },

        _onDeleteClicked: function (e) {
            var taskId = $(e.currentTarget).data('task-id');
            var task = this._collection.get(taskId);
            task.destroy({ 
                error: this._onTaskDeleted.bind(this),
                success: this._onTaskDeleted.bind(this)
            });
        },

        _onTaskDeleted: function () {
            this._collection.fetch().then(this.render.bind(this));
        }
	});

	return TaskListView;
});