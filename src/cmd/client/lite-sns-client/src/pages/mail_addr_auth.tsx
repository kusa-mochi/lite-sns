import { useNavigate } from "react-router";
import Button from "../components/molecules/button";
import Card from "../components/atoms/card";

export default function MailAddrAuth() {
    const navigate = useNavigate()
    return (
        <Card topBorder>
            <div>メールアドレスが認証できました。</div>
            <div>次のボタンからサインインできます。</div>
            <Button primary onClick={() => navigate("/signin")}>Signin</Button>
        </Card>
    )
}
