import { useEffect, useState } from "react"

export default function Timeline() {
    const [username, setUsername] = useState("")

    useEffect(() => {
        // TODO: APIサーバーからユーザー名を取得する。
        setUsername("ダミー")
    }, [])

    return (
        <>
            <h2>タイムライン</h2>
            <div>{username}&nbsp;としてサインインしています。</div>
        </>
    )
}
