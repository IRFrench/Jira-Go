package jira

type Issue struct {
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Watcher struct {
			Self       string `json:"self"`
			IsWatching bool   `json:"isWatching"`
			WatchCount int    `json:"watchCount"`
			Watchers   []struct {
				Self        string `json:"self"`
				AccountID   string `json:"accountId"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
			} `json:"watchers"`
		} `json:"watcher"`
		Attachment []struct {
			ID       int    `json:"id"`
			Self     string `json:"self"`
			Filename string `json:"filename"`
			Author   struct {
				Self        string `json:"self"`
				Key         string `json:"key"`
				AccountID   string `json:"accountId"`
				AccountType string `json:"accountType"`
				Name        string `json:"name"`
				AvatarUrls  struct {
					Size48 string `json:"48x48"`
					Size24 string `json:"24x24"`
					Size16 string `json:"16x16"`
					Size32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
			} `json:"author"`
			Created   string `json:"created"`
			Size      int    `json:"size"`
			MimeType  string `json:"mimeType"`
			Content   string `json:"content"`
			Thumbnail string `json:"thumbnail"`
		} `json:"attachment"`
		SubTasks []struct {
			ID   string `json:"id"`
			Type struct {
				ID      string `json:"id"`
				Name    string `json:"name"`
				Inward  string `json:"inward"`
				Outward string `json:"outward"`
			} `json:"type"`
			OutwardIssue struct {
				ID     string `json:"id"`
				Key    string `json:"key"`
				Self   string `json:"self"`
				Fields struct {
					Status struct {
						IconUrl string `json:"iconUrl"`
						Name    string `json:"name"`
					} `json:"status"`
				} `json:"fields"`
			} `json:"outwardIssue"`
		} `json:"sub-tasks"`
		Description struct {
			Type    string `json:"type"`
			Version int    `json:"version"`
			Content []struct {
				Type    string `json:"type"`
				Content []struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"content"`
			} `json:"content"`
		} `json:"description"`
		Project struct {
			Self            string            `json:"self"`
			ID              string            `json:"id"`
			Key             string            `json:"key"`
			Name            string            `json:"name"`
			AvatarUrls      map[string]string `json:"avatarUrls"`
			ProjectCategory struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"projectCategory"`
			Simplified bool   `json:"simplified"`
			Style      string `json:"style"`
			Insight    struct {
				TotalIssueCount     int    `json:"totalIssueCount"`
				LastIssueUpdateTime string `json:"lastIssueUpdateTime"`
			} `json:"insight"`
		} `json:"project"`
		Comment []struct {
			Self         string      `json:"self"`
			ID           string      `json:"id"`
			Author       User        `json:"author"`
			Body         CommentBody `json:"body"`
			UpdateAuthor User        `json:"updateAuthor"`
			Created      string      `json:"created"`
			Updated      string      `json:"updated"`
			Visibility   struct {
				Type       string `json:"type"`
				Value      string `json:"value"`
				Identifier string `json:"identifier"`
			} `json:"visibility"`
		} `json:"comment"`
		IssueLinks []struct {
			ID           string   `json:"id"`
			Type         LinkType `json:"type"`
			OutwardIssue struct {
				ID     string `json:"id"`
				Key    string `json:"key"`
				Self   string `json:"self"`
				Fields struct {
					Status struct {
						IconUrl string `json:"iconUrl"`
						Name    string `json:"name"`
					} `json:"status"`
				} `json:"fields"`
			} `json:"outwardIssue"`
			InwardIssue struct {
				ID     string `json:"id"`
				Key    string `json:"key"`
				Self   string `json:"self"`
				Fields struct {
					Status struct {
						IconUrl string `json:"iconUrl"`
						Name    string `json:"name"`
					} `json:"status"`
				} `json:"fields"`
			} `json:"inwardIssue"`
		} `json:"issuelinks"`
		Worklog []struct {
			Self         string      `json:"self"`
			ID           string      `json:"id"`
			Author       User        `json:"author"`
			UpdateAuthor User        `json:"updateAuthor"`
			Comment      CommentBody `json:"comment"`
			Updated      string      `json:"updated"`
			Visibility   struct {
				Type       string `json:"type"`
				Value      string `json:"value"`
				Identifier string `json:"identifier"`
			} `json:"visibility"`
			Started          string `json:"started"`
			TimeSpent        string `json:"timeSpent"`
			TimeSpentSeconds int    `json:"timeSpentSeconds"`
			IssueID          string `json:"issueId"`
		} `json:"worklog"`
		Updated      int `json:"updated"`
		TimeTracking struct {
			OriginalEstimate         string `json:"originalEstimate"`
			RemainingEstimate        string `json:"remainingEstimate"`
			TimeSpent                string `json:"timeSpent"`
			OriginalEstimateSeconds  int    `json:"originalEstimateSeconds"`
			RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
			TimeSpentSeconds         int    `json:"timeSpentSeconds"`
		} `json:"timetracking"`
	} `json:"fields"`
}

type User struct {
	Self        string `json:"self"`
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
}

type CommentBody struct {
	Type    string `json:"type"`
	Version int    `json:"version"`
	Content []struct {
		Type    string `json:"type"`
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	} `json:"content"`
}

type LinkType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
}

type JiraReponse struct {
	Expand     string  `json:"expand"`
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}
