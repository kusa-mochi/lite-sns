import { css } from "@emotion/css"
import { useTheme } from "../../providers/themeProvider"
import MenuButton from "../molecules/menu_button"

type Props = {
    height: string
}

export default function MenuBar (props: Props) {
    const { height } = props
    const theme = useTheme()

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
            <MenuButton>AAA</MenuButton>
            <MenuButton>AAA</MenuButton>
            <MenuButton>AAA</MenuButton>
            <MenuButton>AAA</MenuButton>
        </div>
    )
}
