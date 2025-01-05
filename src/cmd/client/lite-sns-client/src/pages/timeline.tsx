import { useEffect, useState } from "react"
import { callAPI } from "../utils/api_utils"
import { useConfig } from "../providers/configProvider"

export default function Timeline() {
    const config = useConfig()
    const [username, setUsername] = useState("")

    useEffect(() => {
        // APIサーバーからユーザー情報を取得する。
        callAPI(
            `http://${config.appServer.ip}:${config.appServer.port}${config.appServer.apiPrefix}/auth_user/get_user_info`,
            "POST",
            {
                // TODO: ContextからユーザーIDを取得し、ここに設定する。
            },
            0,
            (response: any) => {
                console.log(`username: ${response.username}`)
                setUsername(response.username)
            }
        )
    }, [])

    return (
        <>
            <h2>タイムライン</h2>
            <div>{username}&nbsp;としてサインインしています。</div>
        </>
    )
}
