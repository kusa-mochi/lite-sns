import { css } from "@emotion/css";
import Button from "../components/molecules/button";

export default function Home() {

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
                    <Button onClick={() => {location.replace("/signup")}}>Sign up</Button>
                </div>
                <div className={signinButtonContainerStyle}>
                    <Button onClick={() => {location.replace("/signin")}}>Sign in</Button>
                </div>
            </div>
        </>
    )
}
