import { useEffect, useState } from "react"
import { callAPI } from "../utils/api_utils"
import { useConfig } from "../providers/configProvider"
import { useAuth } from "../providers/authProvider"

export default function Timeline() {
    const config = useConfig()
    const auth = useAuth()  // ユーザーIDなどの情報
    const [username, setUsername] = useState("")

    useEffect(() => {
        console.log(`tokenString: ${auth.tokenString}`)
        // APIサーバーからユーザー情報を取得する。
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
    }, [auth])

    return (
        <>
            <h2>タイムライン</h2>
            <div>{username}&nbsp;としてサインインしています。</div>
        </>
    )
}
