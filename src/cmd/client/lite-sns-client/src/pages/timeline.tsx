import { useEffect, useState } from "react";
import { callAPI } from "../utils/api_utils";
import { useConfig } from "../providers/configProvider";
import { useAuth } from "../providers/authProvider";
import Card from "../components/atoms/card";
import { useNavigate } from "react-router";
import UserPage from "../components/templates/user_page";
import Button from "../components/molecules/button";
import MenuButton from "../components/molecules/menu_button";

export default function Timeline() {
  const config = useConfig();
  const auth = useAuth(); // ユーザーIDなどの情報
  const navigate = useNavigate()
  const [username, setUsername] = useState("");
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    console.log(`tokenString: ${auth.tokenString}`);
    // APIサーバーからユーザー情報を取得する。
    // TODO: この動作はあくまでデバッグ用。本来このページではタイムラインに表示する投稿の情報を取得する。
    callAPI(
      `http://${config.appServer.ip}:${config.appServer.port}${config.appServer.apiPrefix}/auth_user/get_user_info`,
      "POST",
      {},
      auth.userId,
      auth.tokenString,
      (response: any) => {
        if (!response || !response.username) {
          console.log(`redirecting to /signin`)
          navigate("/signin")
          return
        }
        setUsername(response.username);
      },
      undefined,
      (response: any) => {
        console.log(`redirecting to /signin`)
        navigate("/signin")
      },
    );

    // TODO: APIサーバーからタイムラインに表示する投稿の情報を取得する。
    callAPI(
      `http://${config.appServer.ip}:${config.appServer.port}${config.appServer.apiPrefix
      }/auth_user/get_timeline?current_oldest_post_id=${50}`,
      "GET",
      {},
      auth.userId,
      auth.tokenString,
      (response: any) => {
        if (!response || !response.timeline) {
          console.log(`redirecting to /signin`)
          navigate("/signin")
          return
        }
        setPosts(() => response.timeline);
      },
      undefined,
      (response: any) => {
        console.log(`redirecting to /signin`)
        navigate("/signin")
      },
    );
  }, [auth]);

  return (
    <UserPage>
      <div>{username}&nbsp;としてサインインしています。</div>
      {posts.map((post: any) => {
        return (
          // <Card key={post.PostId}>
          //   <div>{post.PostId}</div>
          //   <div>{post.UserId}</div>
          //   <div>{post.UserName}</div>
          //   <div>{post.UserIconBg}</div>
          //   <div>{post.CreatedAt}</div>
          //   <div>{post.UpdatedAt}</div>
          //   <div>{post.PostText}</div>
          // </Card>
          <Card key={post.PostId}>
            <div>
              {/* 左寄せにするヘッダ要素 */}
              <div>
                {/* ユーザーアイコン */}
                <div></div>
                {/* ユーザー名 */}
                <div>{post.UserName}</div>
                {/* 投稿日時 */}
                <div>{post.CreatedAt}</div>
              </div>
              {/* 右寄せにするヘッダ要素 */}
              <div>
                <MenuButton>︙</MenuButton>
              </div>
            </div>
            {/* 投稿本文 */}
            <div>{post.PostText}</div>
            <div>
              {/* いいねボタン */}
              <div>
                <MenuButton>いいね</MenuButton>
              </div>
              {/* お気に入りボタン */}
              <div>
                <MenuButton>お気に入り</MenuButton>
              </div>
              {/* コメントボタン */}
              <div>
                <MenuButton>コメント</MenuButton>
              </div>
              {/* シェアボタン */}
              <div>
                <MenuButton>シェア</MenuButton>
              </div>
            </div>
          </Card>
        );
      })}
    </UserPage>
  );
}
