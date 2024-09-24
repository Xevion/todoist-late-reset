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

type Item struct {
	ID             string   `json:"id"`
	UserID         string   `json:"user_id"`
	ProjectID      string   `json:"project_id"`
	Content        string   `json:"content"`
	Description    string   `json:"description"`
	Priority       int      `json:"priority"`
	Due            DueDate  `json:"due"`
	ParentID       *string  `json:"parent_id"`
	ChildOrder     int      `json:"child_order"`
	SectionID      *string  `json:"section_id"`
	DayOrder       int      `json:"day_order"`
	Collapsed      bool     `json:"collapsed"`
	Labels         []string `json:"labels"`
	AddedByUID     string   `json:"added_by_uid"`
	AssignedByUID  string   `json:"assigned_by_uid"`
	ResponsibleUID *string  `json:"responsible_uid"`
	Checked        bool     `json:"checked"`
	IsDeleted      bool     `json:"is_deleted"`
	SyncID         *string  `json:"sync_id"`
	AddedAt        string   `json:"added_at"`
	Duration       Duration `json:"duration"`
}

type Duration struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

type DueDate struct {
	Date        string  `json:"date"`
	Timezone    *string `json:"timezone"`
	String      string  `json:"string"`
	Lang        string  `json:"lang"`
	IsRecurring bool    `json:"is_recurring"`
}
