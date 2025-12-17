import { useEffect, useState } from "react";
import UserPage from "../components/templates/user_page";
import { useAuth } from "../providers/authProvider";
import Card from "../components/atoms/card";

export default function Post() {
    const auth = useAuth(); // ユーザーIDなどの情報
    const [username, setUsername] = useState("");

    useEffect(() => {
        console.log(`tokenString: ${auth.tokenString}`);
    }, [auth])
    return (
        <UserPage>
            {/* 投稿作成ページ */}
            <div>{username}&nbsp;としてサインインしています。</div>
            <Card>
                
            </Card>
        </UserPage>
    )
}
