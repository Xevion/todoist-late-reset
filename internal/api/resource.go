package api

type ResourceType string

const (
	Labels               ResourceType = "labels"
	Projects             ResourceType = "projects"
	Items                ResourceType = "items"
	Notes                ResourceType = "notes"
	Sections             ResourceType = "sections"
	Filters              ResourceType = "filters"
	Reminders            ResourceType = "reminders"
	RemindersLocation    ResourceType = "reminders_location"
	Locations            ResourceType = "locations"
	User                 ResourceType = "user"
	LiveNotifications    ResourceType = "live_notifications"
	Collaborators        ResourceType = "collaborators"
	UserSettings         ResourceType = "user_settings"
	NotificationSettings ResourceType = "notification_settings"
	UserPlanLimits       ResourceType = "user_plan_limits"
	CompletedInfo        ResourceType = "completed_info"
	Stats                ResourceType = "stats"
)
