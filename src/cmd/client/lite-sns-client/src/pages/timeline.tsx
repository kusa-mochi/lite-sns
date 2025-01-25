import { useEffect, useState } from "react"
import { callAPI } from "../utils/api_utils"
import { useConfig } from "../providers/configProvider"
import { useAuth } from "../providers/authProvider"
import Card from "../components/atoms/card"

export default function Timeline() {
    const config = useConfig()
    const auth = useAuth()  // ユーザーIDなどの情報
    const [username, setUsername] = useState("")
    const [posts, setPosts] = useState([])

    useEffect(() => {
        console.log(`tokenString: ${auth.tokenString}`)
        // APIサーバーからユーザー情報を取得する。
        // TODO: この動作はあくまでデバッグ用。本来このページではタイムラインに表示する投稿の情報を取得する。
        callAPI(
            `http://${config.appServer.ip}:${config.appServer.port}${config.appServer.apiPrefix}/auth_user/get_user_info`,
            "POST",
            {},
            auth.userId,
            auth.tokenString,
            (response: any) => {
                console.log(`username: ${response.username}`)
                setUsername(response.username)
            },
        )

        // TODO: APIサーバーからタイムラインに表示する投稿の情報を取得する。
        callAPI(
            `http://${config.appServer.ip}:${config.appServer.port}${config.appServer.apiPrefix}/auth_user/get_timeline?current_oldest_post_id=${50}`,
            "GET",
            {},
            auth.userId,
            auth.tokenString,
            (response: any) => {
                console.log(response)
                setPosts(() => response.timeline)
            },
        )
    }, [auth])

    return (
        <>
            <div>{username}&nbsp;としてサインインしています。</div>
            {
                posts.map((post: any) => {
                    return (
                        <Card key={post.PostId}>
                            <div>{post.PostId}</div>
                            <div>{post.UserId}</div>
                            <div>{post.CreatedAt}</div>
                            <div>{post.UpdatedAt}</div>
                            <div>{post.PostText}</div>
                        </Card>
                    )
                })
            }
        </>
    )
}
