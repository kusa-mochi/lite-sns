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
	PostId    int
	UserId    int
	CreatedAt string
	UpdatedAt string
	PostText  string
	GoodCount int
}

type GetTimelineRes struct {
	Timeline []TimelineItem
	Error    error
}

func (c *GetTimelineCommand) Exec(configs *server_configs.ServerConfigs, db *sql.DB) {
	log.Println("get timeline exec")

	// ユーザー以外の投稿をDBから取得する。
	selectData, err := db_utils.Query(
		db,
		"SELECT id, user_id, text, created_at, updated_at FROM post WHERE id < $1 ORDER BY id DESC LIMIT $2",
		c.CurrentOldestPostId,
		10, // TODO: 設定ファイルで設定するようにする。
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
		postId := postData[0].(int)
		userId := postData[1].(int)
		postText := postData[2].(string)
		createdAt := postData[3].(int64)
		updatedAt := postData[4].(int64)
		timeline = append(timeline, TimelineItem{
			PostId:    postId,
			UserId:    userId,
			CreatedAt: db_utils.UnixTimeToString(createdAt),
			UpdatedAt: db_utils.UnixTimeToString(updatedAt),
			PostText:  postText,
			GoodCount: 0, // TODO: DBから値を取得する。
		})
	}
	// selectData, err := db_utils.SelectFrom(
	// 	db,
	// 	[]string{"name", "icon_type", "icon_background_color"},
	// 	"post",
	// 	"WHERE id = $1",
	// 	c.UserId,
	// )
	// if err != nil {
	// 	// 何もせずコマンド終了。
	// 	log.Printf("failed to get a user info corresponding to the user ID (ID=%v) from DB | %s", c.UserId, err.Error())
	// 	c.ResCh <- &GetTimelineRes{
	// 		Username:            "",
	// 		IconType:            "",
	// 		IconBackgroundColor: "",
	// 		Error:               fmt.Errorf("bad request"),
	// 	}
	// 	return
	// }

	// userInfo := selectData[0]

	c.ResCh <- &GetTimelineRes{
		Timeline: timeline,
		Error:    nil,
	}
}
