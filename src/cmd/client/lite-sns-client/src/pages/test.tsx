import Button from "../components/molecules/button"

export default function Test() {
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
                <p>Go to <a href="/test2">Test2 Page</a></p>
            </div>
        </>
    )
}
