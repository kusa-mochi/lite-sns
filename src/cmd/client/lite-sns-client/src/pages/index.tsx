import { css } from "@emotion/css";
import Button from "../components/molecules/button";
import { useNavigate } from "react-router";

export default function Home() {
    const navigate = useNavigate()

    const signupButtonContainerStyle = css`
        margin: 8px;
    `

    const signinButtonContainerStyle = css`
        margin: 8px;
    `

    return (
        <>
            <div>
                TBD: The promotion page will be here :)
            </div>
            <div>
                Go to <a href="/test">Test Page</a>
            </div>
            <div>
                <div className={signupButtonContainerStyle}>
                    <Button onClick={() => {navigate("/signup")}}>Sign up</Button>
                </div>
                <div className={signinButtonContainerStyle}>
                    <Button onClick={() => {navigate("/signin")}}>Sign in</Button>
                </div>
            </div>
        </>
    )
}
