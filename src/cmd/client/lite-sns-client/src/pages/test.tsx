import { MouseEvent, useState } from "react"
import Button from "../components/molecules/button"

export default function Test() {
    const [cnt, setCnt] = useState(0)
    function ClickTest(e: MouseEvent) {
        setCnt(cnt + 1)
        console.log(`x:${e.pageX}, y:${e.pageY}`)
    }
    return (
        <>
            <Button primary>
                This is a test component.
            </Button>
            <div>
                <Button>Sign up</Button>
                <Button>Sign in</Button>
                <Button enabled={false}>Disabled button</Button>
                <Button focused>Focused button</Button>
                <Button active>Active button</Button>
                <Button onClick={ClickTest}>Click Test</Button>
                <div>cnt:{cnt}</div>
                <p>Go to <a href="/test2">Test2 Page</a></p>
            </div>
        </>
    )
}
