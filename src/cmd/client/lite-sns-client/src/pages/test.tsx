import { css } from "@emotion/css"

export default function Test() {
    const testStyle = css`
        color: blue;
    `
    return (
        <>
            <div className={testStyle}>This is a test component.</div>
            <div>
                <p>Go to <a href="/test2">Test2 Page</a></p>
            </div>
        </>
    )
}
