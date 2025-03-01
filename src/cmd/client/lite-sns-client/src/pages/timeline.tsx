import { useEffect, useState } from "react";
import { callAPI } from "../utils/api_utils";
import { useConfig } from "../providers/configProvider";
import { useAuth } from "../providers/authProvider";
import Card from "../components/atoms/card";
import { css } from "@emotion/css";
import { useNavigate } from "react-router";
import MenuButton from "../components/molecules/menu_button";
import { useTheme } from "../providers/themeProvider";

export default function Timeline() {
  const theme = useTheme()
  const config = useConfig();
  const auth = useAuth(); // ユーザーIDなどの情報
  const navigate = useNavigate()
  const [username, setUsername] = useState("");
  const [posts, setPosts] = useState([]);
  const [menuHeight, setMenuHeight] = useState(40)

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

  const pageStyle = css`
    width: 100%;
    height: 100%;
    overflow-x: hidden;
    overflow-y: auto;

    scrollbar-width: none;
    -ms-overflow-style: none;

    &::webkit-scrollbar {
      display: none;
    }
  `;

  const timelineEndStyle = css`
    width: 100%;
    margin: ${menuHeight}px 0;
  `;

  const menuStyle = css`
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    height: ${menuHeight}px;
    background-color: ${theme.palette.base.bodyBackgroundColor};

    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: space-around;
    align-items: center;
    align-content: space-around;
  `

  return (
    <div className={pageStyle}>
      <div>{username}&nbsp;としてサインインしています。</div>
      {posts.map((post: any) => {
        return (
          <Card key={post.PostId}>
            <div>{post.PostId}</div>
            <div>{post.UserId}</div>
            <div>{post.UserName}</div>
            <div>{post.UserIconBg}</div>
            <div>{post.CreatedAt}</div>
            <div>{post.UpdatedAt}</div>
            <div>{post.PostText}</div>
          </Card>
        );
      })}
      <div className={timelineEndStyle}></div>
      <div className={menuStyle}>
        <MenuButton>AAA</MenuButton>
        <MenuButton>AAA</MenuButton>
        <MenuButton>AAA</MenuButton>
        <MenuButton>AAA</MenuButton>
      </div>
    </div>
  );
}
