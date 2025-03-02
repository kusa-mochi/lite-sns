import { css } from "@emotion/css"
import { useTheme } from "../../providers/themeProvider"
import MenuButton from "../molecules/menu_button"
import { useState } from "react"
import UserIcon_Default from "../icons/UserIcon_Default"
import { useNavigate } from "react-router"

type Props = {
    height?: string
}

export default function MenuBar(props: Props) {
    const { height } = props
    const theme = useTheme()
    const navigate = useNavigate()
    const [menuButtonWidth, setMenuButtonWidth] = useState("56px")

    const GoTo = (pagePath: string) => {
        navigate(pagePath)
    }

    const menuStyle = css`
        position: fixed;
        bottom: 0;
        left: 0;
        width: 100%;
        height: ${height};
        background-color: ${theme.palette.base.bodyBackgroundColor};

        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-around;
        align-items: center;
        align-content: space-around;
    `

    return (
        <div className={menuStyle}>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/timeline")}>
                {/* TODO: 「タイムライン」アイコンに差し替える。 */}
                <UserIcon_Default width="24px" height="24px" />
            </MenuButton>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/search")}>
                {/* TODO: 「検索」アイコンに差し替える。 */}
                <UserIcon_Default width="24px" height="24px" />
            </MenuButton>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/post")}>
                {/* TODO: 「投稿」アイコン（"＋"アイコンとか）に差し替える。 */}
                <UserIcon_Default width="24px" height="24px" />
            </MenuButton>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/activities")}>
                {/* TODO: 「アクティビティ」アイコンに差し替える。 */}
                <UserIcon_Default width="24px" height="24px" />
            </MenuButton>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/user")}>
                <UserIcon_Default width="24px" height="24px" />
            </MenuButton>
        </div>
    )
}
