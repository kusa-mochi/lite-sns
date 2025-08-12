import { css } from "@emotion/css"
import { useTheme } from "../../providers/themeProvider"
import MenuButton from "../molecules/menu_button"
import { useState } from "react"
import UserIcon_Default from "../icons/UserIcon_Default"
import { useNavigate } from "react-router"
import HomeIcon from "../icons/HomeIcon"
import SearchIcon from "../icons/SearchIcon"
import PostIcon from "../icons/PostIcon"
import ActivitiesIcon from "../icons/ActivitiesIcon"

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
                <HomeIcon width="24px" height="24px" />
            </MenuButton>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/search")}>
                <SearchIcon width="24px" height="24px" />
            </MenuButton>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/post")}>
                <PostIcon width="24px" height="24px" />
            </MenuButton>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/activities")}>
                <ActivitiesIcon width="24px" height="24px" />
            </MenuButton>
            <MenuButton width={menuButtonWidth} onClick={() => GoTo("/user")}>
                <UserIcon_Default width="24px" height="24px" />
            </MenuButton>
        </div>
    )
}
