package commands

import (
	"database/sql"
	db_utils "lite-sns/m/src/cmd/app_server/api_server_common/db"
	"lite-sns/m/src/cmd/app_server/server_configs"
	"log"
)

type GetTimelineCommand struct {
	UserId              int
	CurrentOldestPostId int // このIDよりも古い投稿を取得する。
	ResCh               chan<- *GetTimelineRes
}

type TimelineItem struct {
	PostId     int64
	UserId     int64
	UserName   string
	UserIconBg string
	CreatedAt  string
	UpdatedAt  string
	PostText   string
	GoodCount  int
}

type GetTimelineRes struct {
	Timeline []TimelineItem
	Error    error
}

func (c *GetTimelineCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("get timeline exec")
	log.Println("db command params:", c.CurrentOldestPostId, configs.App.TimelinePostNumber)

	// ユーザー以外の投稿をDBから取得する。
	selectData, err := db_utils.Query(
		db,
		"SELECT post.id, post.user_id, post.text, post.created_at, sns_user.name, sns_user.icon_background_color FROM post INNER JOIN sns_user ON post.user_id = sns_user.id WHERE post.id < $1 ORDER BY post.id DESC LIMIT $2",
		c.CurrentOldestPostId,
		configs.App.TimelinePostNumber,
	)
	if err != nil {
		// 何もせずコマンド終了。
		log.Printf("failed to get timeline data from DB | %s", err.Error())
		c.ResCh <- &GetTimelineRes{
			Timeline: nil,
			Error:    err,
		}
		return
	}

	timeline := make([]TimelineItem, 0, len(selectData))
	for _, postData := range selectData {
		postId := postData[0].(int64)    // post.id
		userId := postData[1].(int64)    // post.user_id
		postText := postData[2].(string) // post.text
		createdAt := postData[3].(int64) // post.created_at
		userName := postData[4].(string) // sns_user.name
		iconBg := postData[5].(string)   // sns_user.icon_background_color

		timeline = append(timeline, TimelineItem{
			PostId:     postId,
			UserId:     userId,
			UserName:   userName,
			UserIconBg: iconBg,
			CreatedAt:  db_utils.UnixTimeToString(createdAt),
			UpdatedAt:  "---",
			PostText:   postText,
			GoodCount:  0, // TODO: DBから値を取得する。
		})
	}

	c.ResCh <- &GetTimelineRes{
		Timeline: timeline,
		Error:    nil,
	}
}
