import { MouseEvent, useState } from "react"
import Button from "../components/molecules/button"

export default function Signin() {
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
                <Button disabled>Disabled button</Button>
                <Button active>Active button</Button>
                <Button onClick={ClickTest}>Click Test</Button>
                <div>cnt:{cnt}</div>
                <p>Go to <a href="/test2">Test2 Page</a></p>
            </div>
        </>
    )
}
