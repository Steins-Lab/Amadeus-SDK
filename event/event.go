package event

type Event struct {
	PostType      string `json:"post_type"`
	MessageType   string `json:"message_type"`
	MetaEventType string `json:"meta_event_type"`
	Time          int    `json:"time"`
	SelfID        int64  `json:"self_id"`
	SubType       string `json:"sub_type"`
	GroupID       int    `json:"group_id"`
	Message       string `json:"message"`
	MessageSeq    int    `json:"message_seq"`
	RawMessage    string `json:"raw_message"`
	Sender        struct {
		Age      int    `json:"age"`
		Area     string `json:"area"`
		Card     string `json:"card"`
		Level    string `json:"level"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
		Sex      string `json:"sex"`
		Title    string `json:"title"`
		UserID   int    `json:"user_id"`
	} `json:"sender"`
	Anonymous interface{} `json:"anonymous"`
	Font      int         `json:"font"`
	UserID    int         `json:"user_id"`
	MessageID int         `json:"message_id"`
	Status    struct {
		AppEnabled     bool        `json:"app_enabled"`
		AppGood        bool        `json:"app_good"`
		AppInitialized bool        `json:"app_initialized"`
		Good           bool        `json:"good"`
		Online         bool        `json:"online"`
		PluginsGood    interface{} `json:"plugins_good"`
		Stat           struct {
			PacketReceived  int `json:"packet_received"`
			PacketSent      int `json:"packet_sent"`
			PacketLost      int `json:"packet_lost"`
			MessageReceived int `json:"message_received"`
			MessageSent     int `json:"message_sent"`
			DisconnectTimes int `json:"disconnect_times"`
			LostTimes       int `json:"lost_times"`
			LastMessageTime int `json:"last_message_time"`
		} `json:"stat"`
	} `json:"status"`
	Interval int `json:"interval"`
	TargetID int `json:"target_id"`
}

type Request struct {
	Action string `json:"action"`
	Params any    `json:"params"`
	Echo   string `json:"echo"`
}
type Response struct {
	Status  string      `json:"status"`
	RetCode int         `json:"retcode"`
	Data    interface{} `json:"data"`
	Echo    string      `json:"echo"`
}
type GroupMessage struct {
	GroupId    int    `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}
