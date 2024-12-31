import Button from "../components/molecules/button";

export default function MailAddrAuth() {
    return (
        <>
            <div>メールアドレスが認証できました。</div>
            <div>次のボタンからサインインできます。</div>
            <Button primary onClick={() => location.replace("/signin")}>Signin</Button>
        </>
    )
}
