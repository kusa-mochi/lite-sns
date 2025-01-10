import { useNavigate } from "react-router";
import Button from "../components/molecules/button";

export default function MailAddrAuth() {
    const navigate = useNavigate()
    return (
        <>
            <div>メールアドレスが認証できました。</div>
            <div>次のボタンからサインインできます。</div>
            <Button primary onClick={() => navigate("/signin")}>Signin</Button>
        </>
    )
}
